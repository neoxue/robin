package loadbalance

/*
IRetry defines the interface of the retry rule
*/
type IRetry interface {
	IsRetriable() bool
}
