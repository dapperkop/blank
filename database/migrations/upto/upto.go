package upto

import (
	"github.com/dapperkop/blank/database"
	"github.com/dapperkop/blank/helpers/filesystem"
	"github.com/dapperkop/blank/logger"
	"github.com/dapperkop/blank/types"
	"github.com/pressly/goose"
)

var upToFlag types.NullUpTo

// GetFlag func ...
func GetFlag() *types.NullUpTo {
	return &upToFlag
}

// Need func ...
func Need() bool {
	return upToFlag.Valid
}

// Run func ...
func Run(migrationsDir string) {
	if !Need() {
		return
	}

	var (
		dir = filesystem.Getwd() + "/" + migrationsDir
		err error
	)

	err = goose.UpTo(database.DB, dir, upToFlag.Version)

	if err != nil {
		logger.Logger.Fatalln(err)
	}
}
