package mongodb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	err := initClient()
	time.Sleep(1 * time.Minute)
	assert.NoError(t, err)
}
