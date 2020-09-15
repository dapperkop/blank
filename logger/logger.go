package logger

import (
	"os"
	"time"

	"github.com/dapperkop/blank/consts"
	"github.com/dapperkop/blank/helpers"
	"github.com/dapperkop/blank/types"
	"github.com/BurntSushi/toml"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	debug         = consts.DefaultLoggerDebug
	fileFormatter = &logrus.TextFormatter{DisableColors: consts.DefaultLoggerFileFormatterDisableColors}
	level         = consts.DefaultLoggerLevel
	levels        = map[bool]logrus.Level{consts.DefaultLoggerDebug: consts.DefaultLoggerLevel, !consts.DefaultLoggerDebug: consts.NotDefaultLoggerLevel}
	linkName      = consts.DefaultLoggerLinkName
	logsDir       = loadDirs().LogsDir
	maxAge        = consts.DefaultLoggerMaxAge
	pattern       = consts.DefaultLoggerPattern
	publishDir    = consts.DefaultPublishDir
	publishFile   = consts.DefaultPublishFile
	reportCaller  = consts.DefaultLoggerReportCaller
	reportCallers = map[bool]bool{consts.DefaultLoggerDebug: consts.DefaultLoggerReportCaller, !consts.DefaultLoggerDebug: !consts.DefaultLoggerReportCaller}
	timezone      = consts.DefaultLoggerTimezone
	ttyFormatter  = &logrus.TextFormatter{ForceColors: consts.DefaultLoggerTtyFormatterForceColors, FullTimestamp: consts.DefaultLoggerTtyFormatterFullTimestamp}

	// Logger var ...
	Logger = new()
)

func getLevelHooks() logrus.LevelHooks {
	var (
		hook       logrus.Hook
		levelHooks = make(logrus.LevelHooks)
	)

	hook = lfshook.NewHook(getRotateLogs(), fileFormatter)

	for _, v := range hook.Levels() {
		levelHooks[v] = append(levelHooks[v], hook)
	}

	return levelHooks
}

func getLocation() *time.Location {
	var (
		err      error
		location *time.Location
	)

	location, err = time.LoadLocation(timezone)

	if err != nil {
		Logger.Fatalln(err)
	}

	return location
}

func getRotateLogs() *rotatelogs.RotateLogs {
	var (
		dir            string
		err            error
		linkNameOption string
		locationOption *time.Location
		maxAgeOption   time.Duration
		p              string
		rotateLogs     *rotatelogs.RotateLogs
	)

	dir, err = os.Getwd()

	if err != nil {
		Logger.Fatalln(err)
	}

	linkNameOption = dir + "/" + logsDir + "/" + linkName
	locationOption = getLocation()
	maxAgeOption = maxAge
	p = linkNameOption + ".%Y-%m-%d_%H:%M"

	rotateLogs, err = rotatelogs.New(
		p,
		rotatelogs.WithLinkName(linkNameOption),
		rotatelogs.WithLocation(locationOption),
		rotatelogs.WithMaxAge(maxAgeOption),
	)

	if err != nil {
		Logger.Fatalln(err)
	}

	return rotateLogs
}

func loadDirs() types.Dirs {
	var (
		dir string
		err error
	)

	dir, err = os.Getwd()

	if err != nil {
		Logger.Fatalln(err)
	}

	var (
		dirs  = helpers.InitDirs()
		fpath = dir + "/" + publishDir + "/" + publishFile
	)

	_, err = os.Stat(fpath)

	if os.IsNotExist(err) {
		return dirs
	}

	_, err = toml.DecodeFile(fpath, &dirs)

	if err != nil {
		Logger.Fatalln(err)
	}

	return dirs
}

func new() *logrus.Logger {
	var logger = logrus.New()

	logger.SetFormatter(ttyFormatter)
	logger.SetLevel(level)
	logger.SetReportCaller(reportCaller)

	return logger
}

func setDebug(value bool) {
	debug = value
	level = levels[debug]
	reportCaller = reportCallers[debug]

	Logger.SetLevel(level)
	Logger.SetReportCaller(reportCaller)
}

func setHook() {
	if len(Logger.Hooks) > 0 {
		// Update timezone for rotateLogs
		Logger.ReplaceHooks(getLevelHooks())
	} else {
		// Set Logger hook and File formatter
		Logger.AddHook(lfshook.NewHook(getRotateLogs(), fileFormatter))
	}
}

func setTimezone(value string) {
	timezone = value

	setHook()
}

// Setup func ...
func Setup(config types.Logger) {
	setDebug(config.Debug)
	setTimezone(config.Timezone)
}
