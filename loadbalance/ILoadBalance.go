package loadbalance

/*
ILoadBalance defines the rule of choosing a appropriate server to initialize a call
*/
type ILoadBalance interface {
	AddServers(newServers []Server)
	ChooseServer(key string) Server
	MarkServerDown(server Server)
	GetAllServers() []Server
	GetReachableServers() []Server
}
