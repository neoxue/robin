package loadbalancer

/*
IClientConfigKey defines the interface of the key
*/

type IClientConfigKey interface {
	Key() string
	Type() interface{}
}
