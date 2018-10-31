package loadbalancer

/*
loadbalancerbuilder build a loadbalancer
use facebookgo/inject to
*/
type loadbalancerbuilder struct {
	Rule              IRule              ``
	Ping              IPing              ``
	Serverlist        IServerList        ``
	ServerListUpdater IServerListUpdater ``
	ServerListFilter  IServerListFilter  ``
}

func (lbb *loadbalancerbuilder) build() ILoadBalancer {

}
