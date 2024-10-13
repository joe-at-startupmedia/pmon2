package del

import (
	"pmon3/cli"
	"pmon3/cli/cmd/base"
	"pmon3/cli/cmd/group/list"
	"pmon3/pmond/protos"
	"time"
)

func Delete(idOrName string) *protos.CmdResp {
	sent := base.SendCmd("group_del", idOrName)
	newCmdResp := base.GetResponse(sent)
	if len(newCmdResp.GetError()) == 0 {
		time.Sleep(cli.Config.GetCmdExecResponseWait())
		list.Show()
	}
	return newCmdResp
}
