package stubs

type AuthenticationUtilityStub struct {
	GenerateJWTFn func() (tokenString string, err error)
}

func (stub AuthenticationUtilityStub) GenerateJWT(email string, username string) (tokenString string, err error) {
	return stub.GenerateJWTFn()
}


