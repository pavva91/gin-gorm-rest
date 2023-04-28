package controllers

import (
	// "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type pingServiceMock struct {
	handlePingFn func() (string, error)
}

func (mock pingServiceMock) HandlePing() (string, error) {
	return mock.handlePingFn()
}

type PingTestSuite struct {
	suite.Suite
	GinContextPointer  *gin.Context
	GinEnginePointer   *gin.Engine
	HttpResponseWriter http.ResponseWriter
	W                  *httptest.ResponseRecorder
	ServiceMock        pingServiceMock
}

// Setup Stub Values
func (suite *PingTestSuite) SetupTest() {
	// not strictly required to unit test (will run also without this line)
	gin.SetMode(gin.TestMode)
	suite.W = httptest.NewRecorder()

	suite.GinContextPointer, _ = gin.CreateTestContext(suite.W)
}

// func TestPingWithError(t *testing.T) {
// 	serviceMock := pingServiceMock{}
// 	serviceMock.handlePingFn = func() (string, error) {
// 		return "", errors.New("error executing ping")
// 	}
// 	services.PingService = serviceMock
//
// 	response := httptest.NewRecorder()
// 	context, _ := gin.CreateTestContext(response)
//
// 	PingController.Ping(context)
//
// 	if response.Code != http.StatusInternalServerError {
// 		t.Error("response code should be 500")
// 	}
//
// 	if response.Body.String() != "error executing ping" {
// 		t.Error("response body should say 'error'")
// 	}
// }

// func TestPingNoError1(t *testing.T) {
//
// 	expectedHttpStatus := http.StatusOK
// 	expectedHttpBody := "{\"message\":\"pong\"}"
//
// 	serviceMock := pingServiceMock{}
// 	serviceMock.handlePingFn = func() (string, error) {
// 		return "pong", nil
// 	}
// 	services.PingService = serviceMock
//
// 	response := httptest.NewRecorder()
// 	context, _ := gin.CreateTestContext(response)
//
// 	PingController.Ping(context)
//
// 	if response.Code != expectedHttpStatus {
// 		t.Error("response code should be 200")
// 	}
//
// 	if response.Body.String() != expectedHttpBody {
// 		t.Error("response body should say 'pong'")
// 	}
// }

func (suite *PingTestSuite) Test_Ping_Return200() {

	expectedHttpStatus := http.StatusOK
	expectedHttpBody := "{\"message\":\"pong\"}"
	// expectedHttpBody := "pong"

	suite.ServiceMock = pingServiceMock{}
	suite.ServiceMock.handlePingFn = func() (string, error) {
		return "pong", nil
	}
	services.PingService = suite.ServiceMock

	PingController.Ping(suite.GinContextPointer)

	actualHttpStatus := suite.GinContextPointer.Writer.Status()
	actualHttpBody := suite.W.Body.String()

	assert.Equal(suite.T(), actualHttpStatus, expectedHttpStatus)
	assert.Equal(suite.T(), actualHttpBody, expectedHttpBody)

}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSuitePingController(t *testing.T) {
	suite.Run(t, new(HealthTestSuite))
}
