package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCharByIndexPanic(t *testing.T) {
	//Act
	_, err := getCharByIndex(-1)

	//Assert
	assert := assert.New(t)
	assert.Error(err, ErrCharIndexOutOfRange.Error())
}

func TestGetCharByIndex(t *testing.T) {
	//Arrange
	testCases := []struct {
		paramIndex   int64
		expectedChar string
	}{
		{1, "A"},
		{4, "D"},
		{26, "Z"},
		{27, "A"},
		{52, "Z"},
		{77, "Y"},
		{78, "Z"},
		{208, "Z"},
		{210, "B"},
	}

	for _, tc := range testCases {
		//Act
		actualResult, err := getCharByIndex(tc.paramIndex)

		//Assert
		assert := assert.New(t)
		assert.NoError(err)
		assert.Equal(tc.expectedChar, actualResult)
	}
}
