package config

import (
	"os"

	"github.com/dapperkop/blank/helpers"
	"github.com/dapperkop/blank/helpers/filesystem"
	"github.com/dapperkop/blank/logger"
	"github.com/dapperkop/blank/types"
	"github.com/BurntSushi/toml"
)

var (
	configDir    string
	configFlag   types.NullConfig
	oldConfigDir string
)

func createConfig(mode string) {
	var (
		config = helpers.InitConfig()
		err    error
		file   *os.File
		name   = filesystem.Getwd() + "/" + configDir + "/" + mode + ".toml"
	)

	file, err = os.Create(name)

	if err != nil {
		logger.Logger.Fatalln(err)
	}

	err = toml.NewEncoder(file).Encode(config)

	if err != nil {
		logger.Logger.Fatalln(err)
	}

	err = file.Close()

	if err != nil {
		logger.Logger.Fatalln(err)
	}
}

// GetFlag func ...
func GetFlag() *types.NullConfig {
	return &configFlag
}

// GetFlagDir func ...
func GetFlagDir() string {
	return configFlag.Dir
}

// Need func ...
func Need() bool {
	return configFlag.Valid
}

// Run func ...
func Run(mode string) {
	if !Need() {
		return
	}

	var (
		newPath = filesystem.Getwd() + "/" + configDir
		oldPath = filesystem.Getwd() + "/" + oldConfigDir
	)

	if filesystem.IsExist(oldPath) {
		filesystem.Rename(oldPath, newPath)
	}

	if filesystem.IsNotExist(newPath) {
		filesystem.CreateDir(newPath)
	}

	var path = newPath + "/" + mode + ".toml"

	if filesystem.IsNotExist(path) {
		createConfig(mode)
	}
}

// Setup func ...
func Setup(oldDir string, dir string) {
	configDir = dir
	oldConfigDir = oldDir
}
