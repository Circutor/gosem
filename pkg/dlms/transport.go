package dlms

//go:generate mockery --name Transport --structname TransportMock --filename transportMock.go

// Transport specifies the transport layer.
type Transport interface {
	Connect() (err error)
	Disconnect() (err error)
	Send(src []byte) (out []byte, err error)
}
