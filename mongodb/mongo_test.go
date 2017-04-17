package mongodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	err := initClient()
	assert.NoError(t, err)
}
