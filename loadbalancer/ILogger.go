package loadbalancer

/*
ILogger defines the logger interface, and use logrus as default
5 level's logs and
*/

type ILogger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Panic(v ...interface{})
}
