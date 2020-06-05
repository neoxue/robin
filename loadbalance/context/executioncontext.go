package context

/*
executioncontext is the context of the execution
*/

type executioncontext struct {
	Context      map[string]interface{}
	RetryHandler loadbalance.IRetry
}

func (ec *executioncontext) GetRequest() interface{} {

	return ""
}
