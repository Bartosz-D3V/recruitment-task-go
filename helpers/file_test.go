package helpers

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"os"
	_ "strings"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "go.uber.org/zap"
)

func TestReadFileIntoSlice(t *testing.T) {
	// Test cases
	tests := []struct {
		name          string
		path          string
		fileContent   string
		createFile    bool
		expectedLines []int
		expectedError error
	}{
		{
			name:          "Successful read",
			path:          "../testdata/numbers.txt",
			fileContent:   "123\n456\n789\n",
			createFile:    true,
			expectedLines: []int{123, 456, 789},
			expectedError: nil,
		},
		{
			name:          "File not found",
			path:          "nonexistent.txt",
			createFile:    false,
			fileContent:   "",
			expectedLines: nil,
			expectedError: errors.New("open nonexistent.txt: no such file or directory"),
		},
		{
			name:          "Invalid number format",
			path:          "../testdata/invalid.txt",
			fileContent:   "123\nabc\n789\n",
			createFile:    true,
			expectedLines: nil,
			expectedError: errors.New("file contains line with invalid number: abc"),
		},
	}
	log := zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.createFile {
				file, err := os.Create(tt.path)
				assert.NoError(t, err)
				_, err = file.WriteString(tt.fileContent)
				assert.NoError(t, err)
				file.Close()
			}

			// Call function under test
			lines, err := ReadFileIntoSlice(log.Sugar(), tt.path)

			// Assertions
			if tt.expectedError != nil {
				assert.Nil(t, lines)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NotNil(t, lines)
				assert.Equal(t, *lines, tt.expectedLines)
				assert.NoError(t, err)
			}

			// Clean up test file
			if tt.createFile {
				err = os.Remove(tt.path)
				assert.NoError(t, err)
			}
		})
	}
}
