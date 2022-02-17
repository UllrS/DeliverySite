package tools

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

var Logger *logrus.Logger

func InitLogger(lvl logrus.Level) {
	Logger = logrus.New()
	Logger.SetReportCaller(true)
	Logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}
	logfile, err := os.OpenFile("logs/list.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err == nil {
		Logger.Out = os.Stdout
	} else {
		panic(err)
	}
	Logger.SetLevel(lvl)

	Logger.Hooks.Add(&writer.Hook{
		Writer: logfile,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
		},
	})
	Logger.Trace("logger initialized")
}
