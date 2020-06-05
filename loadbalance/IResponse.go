package loadbalance

/*
IResponse fetch the response
*/

type IResponse interface {
	HasPayload() bool
	IsSuccess() bool
	Payload() interface{}
}
