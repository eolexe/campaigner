package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCharByIndex(t *testing.T) {
	//Arrange
	testCases := []struct {
		paramIndex    int64
		expectedChar  string
		expectedError error
	}{
		{-1, "", ErrCharIndexOutOfRange},
		{1, "A", nil},
		{4, "D", nil},
		{26, "Z", nil},
		{27, "", ErrCharIndexOutOfRange},
	}

	for _, tc := range testCases {
		//Act
		actualResult, err := getCharByIndex(tc.paramIndex)

		//Assert
		assert := assert.New(t)

		if tc.expectedError != nil {
			assert.Error(err, tc.expectedError.Error())
			assert.Equal(tc.expectedChar, actualResult)
		} else {
			assert.NoError(err)
			assert.Equal(tc.expectedChar, actualResult)
		}

	}
}
