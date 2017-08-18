package log

import (
	"github.com/getsentry/raven-go"
	"fmt"
	"time"
	"os"
	"github.com/spf13/viper"
)

type logLevel struct {
	ID int
	Name string
}

var (
	LevelError = logLevel{1, "ERROR"}
	LevelWarning = logLevel{2, "WARNING"}
	LevelInfo = logLevel{3, "INFO"}
	LevelDebug = logLevel{4, "DEBUG"}
)

func printLog(message interface{}, level logLevel) {
	fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05.999Z"), "[" + level.Name + "]", message)
}

func Panic(err error) {
	if err != nil && viper.GetInt("loglevel") >= LevelError.ID {
		raven.CaptureErrorAndWait(err, nil)
		printLog(err, LevelError)
		panic(err)
	}
}

func Fatal(err error) {
	if err != nil && viper.GetInt("loglevel") >= LevelError.ID {
		raven.CaptureErrorAndWait(err, nil)
		printLog(err, LevelError)
		os.Exit(1)
	}
}

func Warning(args interface{}) {
	if args != nil && viper.GetInt("loglevel") >= LevelWarning.ID {
		printLog(args, LevelWarning)
	}
}

func Info(args interface{}) {
	if args != nil && viper.GetInt("loglevel") >= LevelInfo.ID {
		printLog(args, LevelInfo)
	}
}

func Debug(args interface{}) {
	if args != nil && viper.GetInt("loglevel") >= LevelDebug.ID {
		printLog(args, LevelDebug)
	}
}