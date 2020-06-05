package loadbalance

/*
IServer is the interface of a server
*/
type IServer interface {
	SetMetaInfo(meta)
	GetMetaInfo() meta
	SetAliveFlag(isAliveFlag bool)
	IsAlive() bool
}
