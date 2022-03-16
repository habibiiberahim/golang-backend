package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceSuccess(t *testing.T) {
	simpleService, err := InitializedService(false)

	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
}

func TestServiceError(t *testing.T) {
	simpleService, err := InitializedService(true)

	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}

func TestConnection(t *testing.T) {
	connection, cleanup := InitializedConnection("Database	")
	assert.NotNil(t, connection)
	cleanup()
}
