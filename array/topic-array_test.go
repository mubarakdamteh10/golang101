package array

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArrayTopic(t *testing.T) {

	arrayTopic := NewArrayTopic()

	reflectValue := reflect.Indirect(reflect.ValueOf(arrayTopic))

	for index := 0; index < reflectValue.NumField(); index++ {
		assert.False(t, reflectValue.Field(index).IsZero(), "Field %s is zero value", reflectValue.Type().Field(index).Name)
	}
}

func TestSumArray(t *testing.T) {
	t.Run("success_sum_list", func(t *testing.T) {
		// Arrange
		request := []int{1, 2, 3, 4, 5}
		// Act
		service := arrayTopic{}
		result := service.SumArray(request)

		// Assert

		assert.Equal(t, 15, result)
	})

	t.Run("success_sum_list_more_number", func(t *testing.T) {
		// Arrange
		request := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		// Act
		service := arrayTopic{}
		result := service.SumArray(request)

		// Assert

		assert.Equal(t, 55, result)
	})
}
