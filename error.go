package gear

type ErrorLogger interface {
	Error(msg string, kvs ...interface{})
}

type ErrorLoggerFunc func(msg string, kvs ...interface{})

func (e ErrorLoggerFunc) Error(msg string, kvs ...interface{}) {
	e(msg, kvs...)
}
