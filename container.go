package run

type Container struct {
	Config
	Logger
}

func (c Container) Run(f func(Container) error) {
	c.Logger.Error.Err(f(c))
}
