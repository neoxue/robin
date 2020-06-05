package loadbalance

/*
ILoadBalancer defines the rule of choosing a appropriate server to initialize a call
*/
type ILoadBalancer interface {
	AddServers(newServers []IServer)
	ChooseServer(key interface{}) IServer
	MarkServerDown(server IServer)
	GetAllServers() []IServer
	GetReachableServers() []IServer
}
