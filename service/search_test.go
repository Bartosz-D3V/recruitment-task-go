package service

import (
	"github.com/Bartosz-D3V/recruitment-task-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestSearchService_BinarySearchHappyPath(t *testing.T) {
	t.Parallel()
	log := zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel))

	tests := []struct {
		numbers      []int
		numberToFind int
		expected     int
	}{
		{[]int{1}, 1, 0},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{900, 1000, 1100, 1200, 1300}, 890, 0},
		{[]int{900, 1000, 1100, 1200, 1300}, 1250, 4},
		{[]int{10, 12, 14, 15, 18, 20, 30, 40, 41}, 20, 5},
		{[]int{10, 12, 14, 15, 18, 20, 30, 40, 41}, 42, 8},
	}

	for _, test := range tests {
		appConfig := config.AppConfig{
			Logger:  log.Sugar(),
			Numbers: &test.numbers,
		}
		searchSvc := New(appConfig)
		result, err := searchSvc.BinarySearch(test.numberToFind)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}

func TestSearchService_BinarySearchNotFound(t *testing.T) {
	t.Parallel()
	log := zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel))

	tests := []struct {
		numbers      []int
		numberToFind int
	}{
		{[]int{1}, 3},
		{[]int{1}, 0},
		{[]int{1, 2, 3}, 30},
		{[]int{900, 1000, 1100, 1200, 1300}, 1500},
		{[]int{10, 12, 14, 15, 18, 20, 30, 40, 41}, 50},
		{[]int{10, 12, 14, 15, 18, 20, 30, 40, 41}, 5},
	}

	for _, test := range tests {
		appConfig := config.AppConfig{
			Logger:  log.Sugar(),
			Numbers: &test.numbers,
		}
		searchSvc := New(appConfig)
		result, err := searchSvc.BinarySearch(test.numberToFind)

		assert.Equal(t, -1, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "number not found")
	}
}
