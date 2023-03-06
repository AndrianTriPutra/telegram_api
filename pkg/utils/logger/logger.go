package logger

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func Newlogger(level string) {
	log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter) //default
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.Formatter.(*logrus.TextFormatter).DisableColors = false    // remove colors
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false // remove timestamp from test output

	switch level {
	case "trace":
		log.Level = logrus.TraceLevel
	case "debug":
		log.Level = logrus.DebugLevel
	case "info":
		log.Level = logrus.InfoLevel
	case "warning":
		log.Level = logrus.WarnLevel
	case "error":
		log.Level = logrus.ErrorLevel
	case "fatal":
		log.Level = logrus.FatalLevel
	case "panic":
		log.Level = logrus.PanicLevel
	}

	//log.Out = os.Stdout

}

func Level(level, head string, msg interface{}) {
	switch level {
	case "trace":
		log.WithFields(logrus.Fields{}).Trace(head, msg)
	case "debug":
		log.WithFields(logrus.Fields{}).Debug(head, msg)
	case "info":
		log.WithFields(logrus.Fields{}).Info(head, msg)
	case "warning":
		log.WithFields(logrus.Fields{}).Warn(head, msg)
	case "error":
		log.WithFields(logrus.Fields{}).Error(head, msg)
	case "fatal":
		log.WithFields(logrus.Fields{}).Fatal(head, msg)
	case "panic":
		log.WithFields(logrus.Fields{}).Panic(head, msg)
	}
}
