package context

import (
	"github.com/neoxue/robin/loadbalancer"
)

/*
executioncontext is the context of the execution
*/

type executioncontext struct {
	Context      map[string]interface{}
	RetryHandler loadbalancer.IRetry
}

func (ec *executioncontext) GetRequest() interface{} {

	return ""
}
