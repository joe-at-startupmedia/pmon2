package kill

import (
	"pmon3/cli/cmd/list"
	"pmon3/cli/pmq"
	"pmon3/pmond"
	"pmon3/pmond/model"
	"time"

	"github.com/spf13/cobra"
)

var (
	forceKill bool
)

var Cmd = &cobra.Command{
	Use:   "kill",
	Short: "Terminate all processes",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		Kill(model.StatusStopped)
	},
}

func init() {
	Cmd.Flags().BoolVarP(&forceKill, "force", "f", false, "force kill all processes")
}

func Kill(processStatus model.ProcessStatus) {
	pmq.New()
	if forceKill {
		pmq.SendCmd("kill", "force")
	} else {
		pmq.SendCmd("kill", "")
	}
	newCmdResp := pmq.GetResponse()
	if len(newCmdResp.GetError()) > 0 {
		pmond.Log.Fatalf(newCmdResp.GetError())
	}
	time.Sleep(pmond.Config.GetCmdExecResponseWait())
	//list command will call pmq.Close
	list.Show()
}