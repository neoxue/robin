package context

/*
loadbalancercontext is the context of a loadbalancer
*/
type loadbalancercontext struct {
	logger loadbalance.ILogger `inject:"foo"`
	lb     loadbalance.ILoadBalancer
}
