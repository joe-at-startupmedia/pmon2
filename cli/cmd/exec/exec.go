package exec

import (
	"os/user"
	"pmon3/cli"
	"pmon3/cli/cmd/base"
	"pmon3/cli/cmd/list"
	"pmon3/pmond/model"
	"pmon3/pmond/protos"
	"time"

	"github.com/spf13/cobra"
)

// process failed auto restart
var flag model.ExecFlags

var Cmd = &cobra.Command{
	Use:     "exec [application_binary]",
	Aliases: []string{"run"},
	Short:   "Spawn a new process",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(flag.User) > 0 && flag.User == "root" && !base.IsRoot() {
			base.OutputError("cannot set process user to root without sudo")
			return
		} else if flag.User == "" {
			user, err := user.Current()
			if err == nil {
				flag.User = user.Username
			}
		}
		base.OpenSender()
		defer base.CloseSender()
		flag.File = args[0]
		Exec(flag.Json())
	},
}

func init() {
	Cmd.Flags().BoolVarP(&flag.NoAutoRestart, "no-autorestart", "n", false, "do not restart upon process failure")
	Cmd.Flags().StringVarP(&flag.User, "user", "u", "", "the processes run user")
	Cmd.Flags().StringVarP(&flag.Log, "log", "l", "", "the processes stdout log")
	Cmd.Flags().StringVarP(&flag.Args, "args", "a", "", "the processes extra arguments")
	Cmd.Flags().StringVarP(&flag.EnvVars, "env-vars", "e", "", "the processes environment variables (space-delimited)")
	Cmd.Flags().StringVar(&flag.Name, "name", "", "the processes name")
	Cmd.Flags().StringVar(&flag.LogDir, "log-dir", "", "the processes stdout log dir")
	Cmd.Flags().StringSliceVarP(&flag.Dependencies, "dependencies", "d", []string{}, "provide a list of process names this process depends on")
	Cmd.Flags().StringSliceVarP(&flag.Groups, "groups", "g", []string{}, "provide a list of group names this process is associated to")
}

func Exec(flags string) *protos.CmdResp {
	base.OpenSender()
	sent := base.SendCmd("exec", flags)
	newCmdResp := base.GetResponse(sent)
	if len(newCmdResp.GetError()) == 0 {
		time.Sleep(cli.Config.GetCmdExecResponseWait())
		list.Show()
	}
	return newCmdResp
}
