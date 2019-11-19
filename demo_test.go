package learning_badges

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	expected := "hello world"
	result := Print()
	assert.EqualValues(t, expected, result)
}
