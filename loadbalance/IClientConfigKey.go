package loadbalance

/*
IClientConfigKey defines the interface of the key
*/

type IClientConfigKey interface {
	Key() string
	Type() interface{}
}
