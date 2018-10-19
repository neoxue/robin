package loadbalance

/*
IRule defines the Rule of a loadbalance;
*/
type IRule interface {
	Choose(key string) IServer
	SetLoadBalance(lb ILoadBalance)
	GetLoadBalance() ILoadBalance
}
