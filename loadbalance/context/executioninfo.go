package context

type executioninfo struct {
	server                       loadbalance.IServer
	NumberOfPastAttemptsOnServer int
	NumberOfPastServersAttempted int
}
