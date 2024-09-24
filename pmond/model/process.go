package model

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-set/v2"
	"github.com/joe-at-startupmedia/depgraph"
	"github.com/sirupsen/logrus"
	"os"
	"os/user"
	"pmon3/pmond/protos"
	"pmon3/pmond/utils/conv"
	"pmon3/pmond/utils/cpu"
	"strings"
	"time"
)

type ProcessStatus int64

const dateTimeFormat = "2006-01-02 15:04:05"

var restartCount = make(map[uint32]uint32)

const (
	StatusQueued ProcessStatus = iota
	StatusInit
	StatusRunning
	StatusStopped
	StatusFailed
	StatusClosed
	StatusBackoff
)

func (s ProcessStatus) String() string {
	switch s {
	case StatusQueued:
		return "queued"
	case StatusInit:
		return "init"
	case StatusRunning:
		return "running"
	case StatusStopped:
		return "stopped"
	case StatusFailed:
		return "failed"
	case StatusClosed:
		return "closed"
	case StatusBackoff:
		return "backoff"
	}
	return "unknown"
}

func StringToProcessStatus(s string) ProcessStatus {
	switch s {
	case "queued":
		return StatusQueued
	case "init":
		return StatusInit
	case "running":
		return StatusRunning
	case "stopped":
		return StatusStopped
	case "failed":
		return StatusFailed
	case "closed":
		return StatusClosed
	case "backoff":
		return StatusBackoff
	}
	return StatusFailed
}

type Process struct {
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Pointer      *os.Process   `gorm:"-" json:"-"`
	Log          string        `gorm:"column:log" json:"log"`
	Name         string        `gorm:"unique" json:"name"`
	ProcessFile  string        `json:"process_file"`
	Args         string        `json:"args"`
	EnvVars      string        `json:"env_vars"`
	Username     string        `json:"username"`
	Dependencies string        `json:"dependencies"`
	Groups       []*Group      `gorm:"many2many:process_groups;"`
	Status       ProcessStatus `json:"status"`
	ID           uint32        `gorm:"primary_key" json:"id"`
	Pid          uint32        `gorm:"column:pid" json:"pid"`
	Uid          uint32        `gorm:"column:uid" json:"uid"`
	Gid          uint32        `gorm:"column:gid" json:"gid"`
	RestartCount uint32        `gorm:"-" json:"-"`
	AutoRestart  bool          `json:"auto_restart"`
}

func (p *Process) RenderTable() []string {
	cpuVal, memVal := "0%", "0.0 MB"
	if p.Status == StatusRunning {
		cpuVal, memVal = cpu.GetExtraInfo(int(p.Pid))
	}

	return []string{
		p.GetIdStr(),
		p.Name,
		p.GetPidStr(),
		p.GetRestartCountStr(),
		p.Status.String(),
		p.Username,
		cpuVal,
		memVal,
		p.UpdatedAt.Format(dateTimeFormat),
	}
}

func (p *Process) Stringify() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.ID)
}

func (p *Process) Json() (string, error) {
	output, err := json.Marshal(p)
	return string(output), err
}

func (p *Process) GetIdStr() string {
	return conv.Uint32ToStr(p.ID)
}

func (p *Process) GetPidStr() string {
	return conv.Uint32ToStr(p.Pid)
}

func (p *Process) GetRestartCount() uint32 {
	return restartCount[p.ID]
}

func (p *Process) GetRestartCountStr() string {
	return conv.Uint32ToStr(p.RestartCount)
}

func (p *Process) ResetRestartCount() {
	restartCount[p.ID] = 0
}

func (p *Process) IncrRestartCount() {
	restartCount[p.ID] += 1
}

func (p *Process) GetGroupHashSet() *set.HashSet[*Group, string] {
	return set.HashSetFrom[*Group, string](p.Groups)
}

func (p *Process) GetGroupNames() []string {
	groupNames := make([]string, len(p.Groups))
	for i := range p.Groups {
		groupNames[i] = p.Groups[i].Name
	}
	return groupNames
}

func (p *Process) ToAppsConfigApp() *AppsConfigApp {

	flags := ExecFlags{
		User:          p.Username,
		Log:           p.Log,
		Args:          p.Args,
		EnvVars:       p.EnvVars,
		Name:          p.Name,
		NoAutoRestart: !p.AutoRestart,
	}

	if len(p.Dependencies) > 0 {
		flags.Dependencies = strings.Split(p.Dependencies, " ")
	}

	if len(p.Groups) > 0 {
		flags.Groups = p.GetGroupNames()
	}

	return &AppsConfigApp{
		File:  p.ProcessFile,
		Flags: flags,
	}
}

//non-receiver methods begin

func FromFileAndExecFlags(processFile string, flags *ExecFlags, logPath string, user *user.User, groups []*Group) *Process {

	var processParams = []string{flags.Name}
	if len(flags.Args) > 0 {
		processParams = append(processParams, strings.Split(flags.Args, " ")...)
	}

	p := Process{
		Pid:          0,
		Log:          logPath,
		Name:         flags.Name,
		ProcessFile:  processFile,
		Args:         strings.Join(processParams[1:], " "),
		EnvVars:      flags.EnvVars,
		Pointer:      nil,
		Status:       StatusQueued,
		AutoRestart:  !flags.NoAutoRestart,
		Dependencies: strings.Join(flags.Dependencies, " "),
		Groups:       groups,
	}

	if user != nil {
		p.Uid = conv.StrToUint32(user.Uid)
		p.Gid = conv.StrToUint32(user.Gid)
		p.Username = user.Username
	}

	return &p
}

func ComputeDepGraph(appsPtr *[]Process) (*[]Process, *[]Process, error) {

	apps := *appsPtr

	if len(apps) > 1 {
		g := depgraph.New()
		depAppNames := make(map[string]Process)
		nonDepAppNames := make(map[string]Process)
		for _, app := range apps {
			if len(app.Dependencies) > 0 {
				appDependencies := strings.Split(app.Dependencies, " ")
				depAppNames[app.Name] = app
				for _, dep := range appDependencies {
					err := g.DependOn(app.Name, dep)
					if err != nil {
						logrus.Errorf("encountered error building app dependency tree: %s", err)
						return nil, nil, err
					}
				}
			} else {
				nonDepAppNames[app.Name] = app
			}
		}

		if len(g.Leaves()) > 0 {

			dependentApps := make([]Process, 0)

			topoSorted := g.TopoSorted()
			for _, appName := range topoSorted {
				if depAppNames[appName].ProcessFile != "" {
					dependentApps = append(dependentApps, depAppNames[appName])
				} else if nonDepAppNames[appName].ProcessFile != "" {
					dependentApps = append(dependentApps, nonDepAppNames[appName])
					delete(nonDepAppNames, appName)
				} else if nonDepAppNames[appName].ProcessFile == "" {
					logrus.Warnf("dependencies: %s is not a valid app name", appName)
				}
			}

			nonDependentApps := make([]Process, len(nonDepAppNames))
			i := 0
			for appName := range nonDepAppNames {
				nonDependentApps[i] = nonDepAppNames[appName]
				i++
			}

			return &nonDependentApps, &dependentApps, nil
		} else {

			return appsPtr, nil, nil
		}

	}

	return appsPtr, nil, nil
}

func ProcessNames(processesPtr *[]Process) []string {

	if processesPtr == nil {
		return []string{}
	}

	processes := *processesPtr

	if len(processes) == 0 {
		return []string{}
	}

	names := make([]string, len(processes))

	i := 0
	for _, p := range processes {
		names[i] = p.Name
		i++
	}

	return names
}

//protobuf methods begin

func (p *Process) ToProtobuf() *protos.Process {
	newProcess := protos.Process{
		Id:           p.ID,
		CreatedAt:    p.CreatedAt.Format(dateTimeFormat),
		UpdatedAt:    p.UpdatedAt.Format(dateTimeFormat),
		Pid:          p.Pid,
		Log:          p.Log,
		Name:         p.Name,
		ProcessFile:  p.ProcessFile,
		Args:         p.Args,
		EnvVars:      p.EnvVars,
		Status:       p.Status.String(),
		AutoRestart:  p.AutoRestart,
		Uid:          p.Uid,
		Username:     p.Username,
		Gid:          p.Gid,
		RestartCount: p.GetRestartCount(),
		Dependencies: p.Dependencies,
		Groups:       GroupsArrayToProtobuf(p.Groups),
	}
	return &newProcess
}

func ProcessFromProtobuf(p *protos.Process) *Process {
	createdAt, err := time.Parse(dateTimeFormat, p.GetCreatedAt())
	if err != nil {
		fmt.Println(err)
	}
	updatedAt, err := time.Parse(dateTimeFormat, p.GetUpdatedAt())
	if err != nil {
		fmt.Println(err)
	}
	newProcess := Process{
		ID:           p.GetId(),
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		Pid:          p.GetPid(),
		Log:          p.GetLog(),
		Name:         p.GetName(),
		ProcessFile:  p.GetProcessFile(),
		Args:         p.GetArgs(),
		EnvVars:      p.GetEnvVars(),
		Status:       StringToProcessStatus(p.GetStatus()),
		AutoRestart:  p.GetAutoRestart(),
		Uid:          p.GetUid(),
		Username:     p.GetUsername(),
		Gid:          p.GetGid(),
		RestartCount: p.GetRestartCount(),
		Dependencies: p.GetDependencies(),
		Groups:       GroupsArrayFromProtobuf(p.GetGroups()),
	}
	return &newProcess
}

func GetGroupString(p *protos.Process) string {
	var processNamesStr string
	groupLength := len(p.Groups)
	if groupLength > 0 {
		processNameArray := make([]string, groupLength)
		for i := range p.Groups {
			processNameArray[i] = p.Groups[i].Name
		}
		processNamesStr = strings.Join(processNameArray, ", ")
	}
	return processNamesStr
}
