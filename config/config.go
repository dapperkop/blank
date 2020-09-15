package config

import (
	"github.com/dapperkop/blank/helpers"
	"github.com/dapperkop/blank/helpers/filesystem"
	"github.com/dapperkop/blank/logger"
	"github.com/dapperkop/blank/publish"
	"github.com/dapperkop/blank/types"
	"github.com/BurntSushi/toml"
)

var (
	config    = helpers.InitConfig()
	configDir = publish.LoadDirs().ConfigDir
)

// Get func ...
func Get(mode string) types.Config {
	var (
		err   error
		fpath = filesystem.Getwd() + "/" + configDir + "/" + mode + ".toml"
	)

	if filesystem.IsNotExist(fpath) {
		return config
	}

	_, err = toml.DecodeFile(fpath, &config)

	if err != nil {
		logger.Logger.Fatalln(err)
	}

	return config
}
