package loadbalance

import "time"

/*
IServerListUpdater defines a interface of dynamic serverlist loadbalancer's update action
*/
type IServerListUpdater interface {
	Start(action UpdateAction)
	Stop()
	Reverse()
	GetLastUpdate() time.Time
	GetCoreThreads() int
}

type UpdateAction interface {
	DoUpdate()
}
