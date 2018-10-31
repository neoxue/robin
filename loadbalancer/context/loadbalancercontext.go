package context

import "github.com/neoxue/robin/loadbalancer"

/*
loadbalancercontext is the context of a loadbalancer
*/
type loadbalancercontext struct {
	logger loadbalancer.ILogger `inject:"foo"`
	lb     loadbalancer.ILoadBalancer
}
