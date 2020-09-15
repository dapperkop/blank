package migrations

import (
	"github.com/dapperkop/blank/helpers/filesystem"
	"github.com/dapperkop/blank/types"
)

var (
	migrationsDir    string
	migrationsFlag   types.NullMigrations
	oldMigrationsDir string
)

// GetFlag func ...
func GetFlag() *types.NullMigrations {
	return &migrationsFlag
}

// GetFlagDir func ...
func GetFlagDir() string {
	return migrationsFlag.Dir
}

// Need func ...
func Need() bool {
	return migrationsFlag.Valid
}

// Run func ...
func Run() {
	if !Need() {
		return
	}

	var (
		newPath = filesystem.Getwd() + "/" + migrationsDir
		oldPath = filesystem.Getwd() + "/" + oldMigrationsDir
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
	migrationsDir = dir
	oldMigrationsDir = oldDir
}
