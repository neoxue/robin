package loadbalance

/*
IPing defines the interface of ping
*/
type IPing interface {
	IsAlive(server Server) bool
}
