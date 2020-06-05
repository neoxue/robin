package loadbalance

/*
meta is the meta info of lb's server
*/
type meta struct {
	AppName     string
	ServerGroup string
	ServiceID   string
	InstanceID  string
}
