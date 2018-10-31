package loadbalancer

/*
IPing defines the interface of ping
*/
type IPing interface {
	IsAlive(server IServer) bool
}
