package downto

import (
	"github.com/dapperkop/blank/database"
	"github.com/dapperkop/blank/helpers/filesystem"
	"github.com/dapperkop/blank/logger"
	"github.com/dapperkop/blank/types"
	"github.com/pressly/goose"
)

var downToFlag types.NullDownTo

// GetFlag func ...
func GetFlag() *types.NullDownTo {
	return &downToFlag
}

// Need func ...
func Need() bool {
	return downToFlag.Valid
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

	err = goose.DownTo(database.DB, dir, downToFlag.Version)

	if err != nil {
		logger.Logger.Fatalln(err)
	}
}
