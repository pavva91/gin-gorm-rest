package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var validationUtility = new(ValidationUtility)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ValidationTestSuite struct {
	suite.Suite
	Int64String    string
	NonInt64String string
	EmptyString    string
	NonEmptyString string
	ZeroString     int
	NonZeroString  int
}

// MyMockedObject is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
type Event struct {
	mock.Mock
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ValidationTestSuite) SetupTest() {
	suite.Int64String = "12"
	suite.NonInt64String = "12a"
	suite.EmptyString = ""
	suite.NonEmptyString = "asdf"
	suite.ZeroString = 0
	suite.NonZeroString = 1
}

func (suite *ValidationTestSuite) Test_IsInt64_Int64String_ReturnTrue() {
	expected := true
	actual := validationUtility.IsInt64(suite.Int64String)
	assert.Equal(suite.T(), actual, expected)
}

func (suite *ValidationTestSuite) Test_IsInt64_NonInt64String_ReturnFalse() {
	expected := false
	actual := validationUtility.IsInt64(suite.NonInt64String)
	assert.Equal(suite.T(), actual, expected)
}

func (suite *ValidationTestSuite) Test_IsEmpty_EmptyString_ReturnTrue() {
	expected := true
	actual := validationUtility.IsEmpty(suite.EmptyString)
	assert.Equal(suite.T(), actual, expected)
}

func (suite *ValidationTestSuite) Test_IsEmpty_NonEmptyString_ReturnFalse() {
	expected := false
	actual := validationUtility.IsEmpty(suite.NonEmptyString)
	assert.Equal(suite.T(), actual, expected)
}

func (suite *ValidationTestSuite) Test_IsZero_ZeroString_ReturnTrue() {
	expected := true
	actual := validationUtility.IsZero(suite.ZeroString)
	assert.Equal(suite.T(), actual, expected)
}

func (suite *ValidationTestSuite) Test_IsZero_NonZeroString_ReturnFalse() {
	expected := false
	actual := validationUtility.IsZero(suite.NonZeroString)
	assert.Equal(suite.T(), actual, expected)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ValidationTestSuite))
}
