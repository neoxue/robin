package context

type IExecutionListener interface {
	OnExecutionStart() abortExectionErr
}

type abortExectionErr interface {
}
