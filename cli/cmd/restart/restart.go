package restart

import (
	"pmon3/cli"
	"pmon3/cli/cmd/base"
	"pmon3/cli/cmd/list"
	"pmon3/pmond/model"
	"time"

	"github.com/spf13/cobra"
)

var flag model.ExecFlags

var Cmd = &cobra.Command{
	Use:     "restart [id or name]",
	Short:   "(Re)start a process by id or name",
	Aliases: []string{"start"},
	Args:    cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(flag.User) > 0 && flag.User == "root" && !base.IsRoot() {
			cli.Log.Fatalf("cannot set process user to root without sudo")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmdRun(cmd.CalledAs(), args, flag.Json())
	},
}

func init() {
	Cmd.Flags().BoolVarP(&flag.NoAutoRestart, "no-autorestart", "n", false, "do not restart upon process failure")
	Cmd.Flags().StringVarP(&flag.User, "user", "u", "", "the processes run user")
	Cmd.Flags().StringVarP(&flag.Log, "log", "l", "", "the processes stdout log")
	Cmd.Flags().StringVarP(&flag.Args, "args", "a", "", "the processes extra arguments")
	Cmd.Flags().StringVarP(&flag.LogDir, "log_dir", "d", "", "the processes stdout log dir")
}

func cmdRun(calledAs string, args []string, flags string) {
	base.OpenSender()
	defer base.CloseSender()
	sent := base.SendCmdArg2(calledAs, args[0], flags)
	newCmdResp := base.GetResponse(sent)
	if len(newCmdResp.GetError()) > 0 {
		cli.Log.Fatalf(newCmdResp.GetError())
	}
	time.Sleep(cli.Config.GetCmdExecResponseWait())
	list.Show()
}
