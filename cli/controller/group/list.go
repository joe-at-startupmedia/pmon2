package group

import (
	"pmon3/cli/controller/base"
	"pmon3/cli/output/group/list"
	"pmon3/model"
	"pmon3/protos"
)

func Show() *protos.CmdResp {
	sent := base.SendCmd("group_list", "")
	newCmdResp := base.GetResponse(sent)
	if len(newCmdResp.GetError()) == 0 {
		all := newCmdResp.GetGroupList().GetGroups()
		var allGroups [][]string
		for _, g := range all {
			group := model.GroupFromProtobuf(g)
			allGroups = append(allGroups, group.RenderTable())
		}
		table_list.Render(allGroups)
	}
	return newCmdResp
}
