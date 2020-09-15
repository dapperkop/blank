package logs

import (
	"github.com/dapperkop/blank/helpers/filesystem"
	"github.com/dapperkop/blank/types"
)

var (
	logsDir    string
	logsFlag   types.NullLogs
	oldLogsDir string
)

// GetFlag func ...
func GetFlag() *types.NullLogs {
	return &logsFlag
}

// GetFlagDir func ...
func GetFlagDir() string {
	return logsFlag.Dir
}

// Need func ...
func Need() bool {
	return logsFlag.Valid
}

// Run func ...
func Run() {
	if !Need() {
		return
	}

	var (
		newPath = filesystem.Getwd() + "/" + logsDir
		oldPath = filesystem.Getwd() + "/" + oldLogsDir
	)

	if filesystem.IsExist(oldPath) {
		filesystem.Rename(oldPath, newPath)
	}

	if filesystem.IsNotExist(newPath) {
		filesystem.CreateDir(newPath)
	}
}

// Setup func ...
func Setup(oldDir string, dir string) {
	logsDir = dir
	oldLogsDir = oldDir
}
