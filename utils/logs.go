package utils
import (
	"github.com/sirupsen/logrus"
	"fmt"
)
var Logger *logrus.Logger

type CustomFormatter struct{}


func Log(logger *logrus.Logger, level logrus.Level, message string) {
    logger.WithFields(logrus.Fields{}).Log(level, message)
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
    var color int
    switch entry.Level {
    case logrus.InfoLevel:
        color = 32 // Green
    case logrus.WarnLevel:
        color = 34 // Yellow
    case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
        color = 31 // Red
    default:
        color = 0 // Default color
    }

    message := fmt.Sprintf("[%s] %s\n", entry.Time.Format("15:04:05"), entry.Message)
    return []byte("\x1b[" + fmt.Sprintf("%d", color) + "m" + message + "\x1b[0m"), nil
}