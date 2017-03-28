package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	//Arrange
	testCases := []struct {
		expectedUserCounter   int
		expectedUserName      string
		expectedProfileLength int
	}{
		{1, "u1", 1},
		{3, "u3", 3},
		{26, "u26", 26},
		{27, "u27", 1},
		{28, "u28", 2},
		{44, "u44", 18},
		{52, "u52", 26},
		{53, "u53", 1},
	}

	for _, tc := range testCases {

		//Act
		actualResult, err := NewUser(int64(tc.expectedUserCounter))

		//Assert
		assert := assert.New(t)
		assert.NoError(err)
		if assert.NotNil(actualResult) {
			assert.Equal(actualResult.User, tc.expectedUserName)
			assert.Len(actualResult.Profile, tc.expectedProfileLength)
		}
	}
}
