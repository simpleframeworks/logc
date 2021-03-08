# LogC

A common and universal logging interface that every package in the `simple frameworks` org uses. 

Adapters to common logging libraries are provided so whatever logging library that you decide to use can be supported.

```go

type Logger interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	WithError(err error) Logger
}

```

## Using with logrus

```go

log = logc.NewLogrus(logrus.New())

log.Trace("some Trace log")
log.Debug("some Debug log")
log.Info("some Info log")
log.Warn("some Warn log")
log.Error("some Error log")


log.WithField("RequestID",1234).Trace("some Trace log")


log.WithFields(map[string]interface{}{
  "RequestID":  1234,
  "Name":       "SomeName",
}).Info("some Info log")


someError := errors.New("an error occurred")
log.WithError(someError).Error("some Error log")

```

---

Currently only logrus is supported but more packages are planned.