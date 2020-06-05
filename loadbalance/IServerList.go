package loadbalance

/*
IServerList defines the interface of ld's serverlist
*/

type IServerList interface {
	GetInitialListOfServers() []IServer
	GetServersUpdated() []IServer
}
