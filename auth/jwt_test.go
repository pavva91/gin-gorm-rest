package auth

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// var authenticationUtility = new(ValidationUtility)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AuthJWTTestSuite struct {
	suite.Suite
	Email              string
	Username           string
	CorrectSignedToken string
	ExpiredToken       string
	InvalidString      string
	EmptyString        string
}

func (suite *AuthJWTTestSuite) SetupTest() {
	suite.Email = "alice@wonder.land"
	suite.Username = "alice1234"
	suite.CorrectSignedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlMTIzNCIsImVtYWlsIjoiYWxpY2VAd29uZGVyLmxhbmQiLCJleHAiOjE2ODI4NTI4OTgsIm5iZiI6LTYyMTM1NTk2ODAwLCJpYXQiOi02MjEzNTU5NjgwMH0.uUtWG_jUK5pekMwgWfKqORlJvcOTobxYRfQAfEFmDUk"
	suite.ExpiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJvYjk3IiwiZW1haWwiOiJib2JAZ21haWwuY29tIiwiZXhwIjoxNjgyNTIyNjIxLCJuYmYiOi02MjEzNTU5NjgwMCwiaWF0IjotNjIxMzU1OTY4MDB9.hEl-RlgpbDMuNa7zYSL-csCqY4oZwOMG3CsrtURbiyU"
	suite.InvalidString = "invalid string"
	suite.EmptyString = ""
}

func (suite *AuthJWTTestSuite) Test_GenerateJWT_Good_ReturnTokenString() {
	expectedJwtHeader := "{\"alg\":\"HS256\",\"typ\":\"JWT\"}"
	expectedJwtPayloadStart := "{\"username\":\"" + suite.Username + "\",\"email\":\"" + suite.Email + "\",\"exp\":"
	// expectedJwtPayloadEnd := ",\"nbf\":-62135596800,\"iat\":-6213559680"

	actual, err := AuthenticationUtility.GenerateJWT(suite.Email, suite.Username)

	log.Info().Msg("Token Value: " + actual)
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	actualJwtTokenSplit := strings.Split(actual, ".")
	header, err := base64.StdEncoding.DecodeString(actualJwtTokenSplit[0])
	payload, err := base64.StdEncoding.DecodeString(actualJwtTokenSplit[1])

	log.Info().Msg("Signature Value: " + actualJwtTokenSplit[2])
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	assert.IsType(suite.T(), "", actual)
	assert.Equal(suite.T(), 3, len(actualJwtTokenSplit))
	assert.Equal(suite.T(), expectedJwtHeader, string(header))
	assert.Contains(suite.T(), string(payload), expectedJwtPayloadStart)
	assert.Contains(suite.T(), string(payload), "\"nbf\"")
	assert.Contains(suite.T(), string(payload), "\"iat\"")
}

func (suite *AuthJWTTestSuite) Test_DecodeJWT_Good_ReturnClaims() {
	freshToken, err := AuthenticationUtility.GenerateJWT(suite.Email, suite.Username)
	log.Info().Msg("Fresh Token: " + freshToken)

	expectedIssuedAt := "0001-01-01 00:17:30 +0017 LMT"

	claims, err := DecodeJWT(freshToken)
	// log.Info().Msg("Claims Value: " + string(claims.))
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	time.Sleep(1 * time.Second)

	now := time.Now()

	assert.Equal(suite.T(), suite.Username, claims.Username)
	assert.Equal(suite.T(), suite.Email, claims.Email)
	assert.Greater(suite.T(), claims.ExpiresAt.Time, now)
	assert.Less(suite.T(), claims.IssuedAt.Time, now)
	assert.Greater(suite.T(), claims.IssuedAt.Time.Add(3*time.Second), now)
	assert.NotEqual(suite.T(), expectedIssuedAt, claims.IssuedAt.String())
	assert.Equal(suite.T(), nil, err)
}

func (suite *AuthJWTTestSuite) Test_ValidateToken_InvalidString_Error() {
	expectedError := errors.New("token is malformed: token contains an invalid number of segments")
	err := ValidateToken(suite.InvalidString)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	assert.Equal(suite.T(), expectedError.Error(), err.Error())
}

func (suite *AuthJWTTestSuite) Test_ValidateToken_ExpiredToken_Error() {
	expectedError := errors.New("token has invalid claims: token is expired")

	// claims, _ := DecodeJWT(suite.ExpiredToken)
	// log.Info().Msg("Expired Token ExpiresAt: " + claims.ExpiresAt.String())

	err := ValidateToken(suite.ExpiredToken)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	assert.Equal(suite.T(), expectedError.Error(), err.Error())
}

func (suite *AuthJWTTestSuite) Test_ValidateToken_Good_NoError() {
	freshToken, _ := AuthenticationUtility.GenerateJWT(suite.Email, suite.Username)
	log.Info().Msg("Fresh Token: " + freshToken)

	log.Info().Msg("Fresh Token: " + freshToken)
	err := ValidateToken(freshToken)
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	claims, err := DecodeJWT(freshToken)
	log.Info().Msg("ExpiresAt: " + claims.ExpiresAt.String())

	assert.Equal(suite.T(), nil, err)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(AuthJWTTestSuite))
}
