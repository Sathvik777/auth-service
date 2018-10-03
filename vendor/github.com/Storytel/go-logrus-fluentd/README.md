# Logrus Fluentd Formatter for Storytel

Formats log messages in a way that works well with fluentd inside GKE.
Also adheres to addition formatting decided by Storytel

## Usage

Set the formatter and supply service name. The formatter takes one argument which is the service name.

```go
log.SetFormatter(logrusfluentd.NewFormatter("service-name"))
```

Ideally we want to keep the text formatter when running in dev-mode

```go
log.SetFormatter(logrusfluentd.NewFormatter("service-name"))
if os.Getenv("ENV") == "dev" {
	log.SetFormatter(&log.TextFormatter{})
}
```
