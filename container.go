package run

type Container struct {
	Config
	Logger
}

var DefaultContainer = Container{
	Config: DefaultConfig,
	Logger: DefaultLogger,
}
