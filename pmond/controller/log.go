package controller

import (
	"fmt"
	"pmon3/pmond"
	"pmon3/pmond/model"
	"pmon3/pmond/protos"
)

func Log(cmd *protos.Cmd) *protos.CmdResp {
	if len(cmd.GetArg1()) == 0 {
		pmond.Log.Fatal("please input start process id or name")
	}
	val := cmd.GetArg1()
	err, process := model.FindProcessByIdOrName(pmond.Db(), val)
	if err != nil {
		newCmdResp := protos.CmdResp{
			Id:    cmd.GetId(),
			Name:  cmd.GetName(),
			Error: fmt.Sprintf("Process (%s) does not exist", val),
		}
		return &newCmdResp
	}
	newProcess := protos.Process{
		Log: process.Log,
	}
	newCmdResp := protos.CmdResp{
		Id:      cmd.GetId(),
		Name:    cmd.GetName(),
		Process: &newProcess,
	}
	return &newCmdResp
}
