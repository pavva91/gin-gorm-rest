package stubs

type PingServiceStub struct {
	HandlePingFn func() (string, error)
}

func (stub PingServiceStub) HandlePing() (string, error) {
	return stub.HandlePingFn()
}

