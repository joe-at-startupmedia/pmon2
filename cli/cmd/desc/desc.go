package desc

import (
	"github.com/spf13/cobra"
	"pmon3/cli/cmd/base"
	"pmon3/cli/output/process/desc"
	"pmon3/pmond/model"
	"pmon3/pmond/utils/conv"
	"strconv"
)

var Cmd = &cobra.Command{
	Use:     "desc [id or name]",
	Aliases: []string{"show"},
	Short:   "Show process information by id or name",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdRun(args)
	},
}

func cmdRun(args []string) {
	base.OpenSender()
	defer base.CloseSender()
	sent := base.SendCmd("desc", args[0])
	newCmdResp := base.GetResponse(sent)
	process := newCmdResp.GetProcess()

	rel := [][]string{
		{"status", process.Status},
		{"id", conv.Uint32ToStr(process.Id)},
		{"name", process.Name},
		{"pid", conv.Uint32ToStr(process.Pid)},
		{"restarted", conv.Uint32ToStr(process.RestartCount)},
		{"process", process.ProcessFile},
		{"args", process.Args},
		{"env-vars", process.EnvVars},
		{"user", process.Username},
		{"log", process.Log},
		{"no-autorestart", strconv.FormatBool(!process.AutoRestart)},
		{"dependencies", process.Dependencies},
		{"groups", model.GetGroupString(process)},
		{"created_at", process.CreatedAt},
		{"updated_at", process.UpdatedAt},
	}
	table_desc.Render(rel)
}
