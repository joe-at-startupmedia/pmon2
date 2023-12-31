package exec

import (
	"encoding/json"
	"fmt"
	"pmon2/pmond"
	"pmon2/pmond/model"
	"pmon2/cli/proxy"
	"os"
	"path"
	"path/filepath"
)

func loadFirst(execPath string, flags string) ([]string, error) {
	data, err := proxy.RunProcess([]string{"start", execPath, flags})
	if err != nil {
		return nil, err
	}

	var tbData []string
	_ = json.Unmarshal(data, &tbData)

	return tbData, nil
}

// check the process already have
func processExist(execPath string) (*model.Process, bool) {
	var process model.Process
	err := pmond.Db().First(&process, "process_file = ?", execPath).Error
	if err != nil {
		return nil, false
	}

	return &process, true
}

func getExecFile(args []string) (string, error) {
	execFile := args[0]
	_, err := os.Stat(execFile)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("%s not exist", execFile)
	}

	if path.IsAbs(execFile) {
		return execFile, nil
	}

	absPath, err := filepath.Abs(execFile)
	if err != nil {
		return "", fmt.Errorf("get file path error: %v", err.Error())
	}

	return absPath, nil
}