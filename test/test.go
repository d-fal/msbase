package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	tx := (1 + 2) / 2
	assert.NotNil(t, tx, "Error")
	t.Error("OOoops")
}
