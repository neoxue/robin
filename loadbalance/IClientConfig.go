package loadbalance

/*
IClientConfig defines the client configuration used by various APIs to initialize clients or load balancers
*/
type IClientConfig interface {
	GetClientName() string
	GetNameSpace() string
	GetProperties() map[string]interface{}
	LoadProperties(clientName string)
	LoadDefaultValues()
	SetProperty(key IClientConfigKey)
}
