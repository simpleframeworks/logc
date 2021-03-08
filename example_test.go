package logc

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

func Example_logger() {

	logrusLog := logrus.New()
	logrusLog.SetOutput(os.Stdout)
	logrusLog.SetFormatter(&logrus.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})

	log := NewLogrus(logrusLog)

	log.Trace("some Trace log")
	log.Debug("some Debug log")
	log.Info("some Info log")
	log.Warn("some Warn log")
	log.Error("some Error log")

	log.WithField("RequestID", 1234).Trace("some Trace log")

	log.WithFields(map[string]interface{}{
		"RequestID": 1234,
		"Name":      "SomeName",
	}).Info("some Info log")

	someError := errors.New("an error occurred")
	log.WithError(someError).Error("some Error log")

	// Output:
	// level=info msg="some Info log"
	// level=warning msg="some Warn log"
	// level=error msg="some Error log"
	// level=info msg="some Info log" Name=SomeName RequestID=1234
	// level=error msg="some Error log" error="an error occurred"
}
