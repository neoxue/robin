package loadbalancer

type IServerListFilter interface {
	GetFilteredListOfServers() []IServer
}
