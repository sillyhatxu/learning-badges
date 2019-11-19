package learning_badges

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example", func() {
	Describe("Make an awesome expression", func() {
		Context("smile", func() {
			It("should result :)", func() {
				Expect(Print1()).To(Equal(expected))
			})
		})
	})
})

const expected = "hello world"

func TestPrint1(t *testing.T) {
	result := Print1()
	assert.EqualValues(t, expected, result)
}

func TestPrint2(t *testing.T) {
	result := Print2()
	assert.EqualValues(t, expected, result)
}

func TestPrint3(t *testing.T) {
	result := Print3()
	assert.EqualValues(t, expected, result)
}

func TestPrint4(t *testing.T) {
	result := Print4()
	assert.EqualValues(t, expected, result)
}

func TestPrint5(t *testing.T) {
	result := Print5()
	assert.EqualValues(t, expected, result)
}

func TestPrint6(t *testing.T) {
	result := Print6()
	assert.EqualValues(t, expected, result)
}

func TestPrint7(t *testing.T) {
	result := Print7()
	assert.EqualValues(t, expected, result)
}

func TestPrint8(t *testing.T) {
	result := Print8()
	assert.EqualValues(t, expected, result)
}

func TestPrint9(t *testing.T) {
	result := Print9()
	assert.EqualValues(t, expected, result)
}
