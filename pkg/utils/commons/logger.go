package commons

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLogger() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{})
	Log.Level = logrus.TraceLevel
	Log.Out = os.Stdout
}
