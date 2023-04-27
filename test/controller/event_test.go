package controllers

import (
	"testing"

	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
	Event1                        models.Event
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
	suite.Event1.ID = 1
	suite.Event1.Title = "Ethereum Hackathon"
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample1() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample2() {
	// expectedLen := 1
	expectedID := 1
	// var eventModel = new(models.Event)
	// events, _ := new(models.Event).ListAllEvents()
	// assert.Equal(suite.T(), len(events), expectedLen)
	assert.Equal(suite.T(), suite.Event1.ID, uint(expectedID))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

func ListEventsTest(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

}
