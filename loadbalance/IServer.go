package loadbalance

/*
IServer is the interface of a server
*/
type IServer interface {
	SetMetaInfo(Meta)
	GetMetaInfo() Meta
	SetAliveFlag(isAliveFlag bool)
	IsAlive() bool
}
