package context

import "github.com/neoxue/robin/loadbalancer"

type executioninfo struct {
	server                       loadbalancer.IServer
	NumberOfPastAttemptsOnServer int
	NumberOfPastServersAttempted int
}
