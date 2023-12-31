package worker

import (
	"errors"
	"fmt"
	"pmon2/cli/service"
	"pmon2/pmond"
	"pmon2/pmond/executor"
	"pmon2/pmond/model"
	"time"
)

func Restart(pFile string, flags *model.ExecFlags) (string, error) {
	m := FindProcess(pFile, flags.Name)
	if m == nil {
		return "", errors.New("try to get process data error")
	}

	cstLog := flags.Log
	if len(cstLog) > 0 && cstLog != m.Log {
		m.Log = cstLog
	}

	// if reset log dir
	if len(flags.LogDir) > 0 && len(flags.Log) == 0 {
		m.Log = ""
	}

	cstName := flags.Name
	if len(cstName) > 0 && cstName != m.Name {
		m.Name = cstName
	}

	// checkout process name whether exist except itself
	if pmond.Db().First(&model.Process{}, "name = ? AND id != ?", cstName, m.ID).Error == nil {
		return "", fmt.Errorf("process name: %s already used, please set other name by --name", cstName)
	}

	extArgs := flags.Args
	if len(extArgs) > 0 {
		m.Args = extArgs
	}

	// get run process user
	runUser, err := GetProcUser(flags)
	if err != nil {
		return "", err
	}

	process, err := executor.Exec(m.ProcessFile, m.Log, m.Name, m.Args, runUser, !flags.NoAutoRestart, flags.LogDir)
	if err != nil {
		return "", err
	}

	// update process extra data
	process.ID = m.ID
	process.CreatedAt = m.CreatedAt
	process.UpdatedAt = time.Now()

	waitData := service.NewProcStat(process).Wait()

	return service.AddData(waitData)
}

func FindProcess(pFile string, name string) *model.Process {
	var rel model.Process
	err := pmond.Db().First(&rel, "process_file = ? AND name = ?", pFile, name).Error
	if err != nil {
		return nil
	}

	return &rel
}
