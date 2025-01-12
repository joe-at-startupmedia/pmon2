package controller

import (
	"pmon3/cli/controller/base"
	"pmon3/protos"
)

func Reset(idOrName string) *protos.CmdResp {

	sent := base.SendCmd("reset", idOrName)
	newCmdResp := base.GetResponse(sent)
	if len(newCmdResp.GetError()) == 0 {
		List()
	}
	return newCmdResp
}
