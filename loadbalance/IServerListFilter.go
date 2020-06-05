package loadbalance

type IServerListFilter interface {
	GetFilteredListOfServers() []IServer
}
