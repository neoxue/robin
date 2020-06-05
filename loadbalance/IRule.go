package loadbalance

/*
IRule defines the Rule of a loadbalancer;
*/
type IRule interface {
	Choose(key string) IServer
	SetLoadBalance(lb ILoadBalancer)
	GetLoadBalance() ILoadBalancer
}
