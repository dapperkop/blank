package publish

import (
	"os"
	"strings"

	"github.com/dapperkop/blank/consts"
	"github.com/dapperkop/blank/helpers"
	"github.com/dapperkop/blank/helpers/filesystem"
	"github.com/dapperkop/blank/logger"
	"github.com/dapperkop/blank/publish/config"
	"github.com/dapperkop/blank/publish/logs"
	"github.com/dapperkop/blank/publish/migrations"
	"github.com/dapperkop/blank/types"
	"github.com/BurntSushi/toml"
)

var (
	configDir     string
	logsDir       string
	migrationsDir string
	publishDir    = consts.DefaultPublishDir
	publishFile   = consts.DefaultPublishFile
)

// LoadDirs func ...
func LoadDirs() types.Dirs {
	var (
		dirs  = helpers.InitDirs()
		err   error
		fpath = filesystem.Getwd() + "/" + publishDir + "/" + publishFile
	)

	if filesystem.IsNotExist(fpath) {
		return dirs
	}

	_, err = toml.DecodeFile(fpath, &dirs)

	if err != nil {
		logger.Logger.Fatalln(err)
	}

	return dirs
}

// Need func ...
func Need() bool {
	return config.Need() || logs.Need() || migrations.Need()
}

// Run func ...
func Run(mode string) {
	if !Need() {
		return
	}

	config.Run(mode)
	logs.Run()
	migrations.Run()

	saveDirs()
}

func saveDirs() {
	var path = filesystem.Getwd() + "/" + publishDir

	if filesystem.IsNotExist(path) {
		filesystem.CreateDir(path)
	}

	var (
		dirs = types.Dirs{ConfigDir: configDir, LogsDir: logsDir, MigrationsDir: migrationsDir}
		err  error
		file *os.File
		name = path + "/" + publishFile
	)

	file, err = os.Create(name)

	if err != nil {
		logger.Logger.Fatalln(err)
	}

	err = toml.NewEncoder(file).Encode(dirs)

	if err != nil {
		logger.Logger.Fatalln(err)
	}

	err = file.Close()

	if err != nil {
		logger.Logger.Fatalln(err)
	}
}

func setDirs(dirs types.Dirs) {
	if config.Need() {
		configDir = config.GetFlagDir()
	} else {
		configDir = dirs.ConfigDir
	}

	if logs.Need() {
		logsDir = logs.GetFlagDir()
	} else {
		logsDir = dirs.LogsDir
	}

	if migrations.Need() {
		migrationsDir = migrations.GetFlagDir()
	} else {
		migrationsDir = dirs.MigrationsDir
	}
}

// Setup func ...
func Setup() {
	var dirs = LoadDirs()

	// Set loaded dirs
	setDirs(dirs)

	config.Setup(dirs.ConfigDir, configDir)
	logs.Setup(dirs.LogsDir, logsDir)
	migrations.Setup(dirs.MigrationsDir, migrationsDir)
}

// Status func ...
func Status() string {
	var status []string

	if config.Need() {
		status = append(status, "config")
	}

	if logs.Need() {
		status = append(status, "logs")
	}

	if migrations.Need() {
		status = append(status, "migrations")
	}

	return strings.Join(status, " & ")
}
