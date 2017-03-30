package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportCampaigns(t *testing.T) {
	//Arrange
	testCases := []struct {
		Attribute            string
		ExpectedCampaignName string
	}{
		{"attr_A_A1", "campaign3"},
		{"attr_A_A19", "campaign2"},
		{"attr_B_B1", "campaign1"},
		{"attr_B_B16", "campaign1"},
		{"attr_A_A20", ""},
		{"attr_B_B17", ""},
	}

	expectedLengthOfTheSortedCampaignList := 37

	//Act
	actualResult := ImportCampaigns(testImportCampaignsData)

	//Assert
	assert := assert.New(t)
	assert.Len(actualResult, expectedLengthOfTheSortedCampaignList)
	for _, tc := range testCases {

		if tc.ExpectedCampaignName != "" {
			if assert.Contains(actualResult, tc.Attribute) {
				actualCampaign := actualResult[tc.Attribute]
				assert.Equal(tc.ExpectedCampaignName, actualCampaign.Name)
			}
		} else {
			assert.NotContains(actualResult, tc.Attribute)
		}
	}
}

func BenchmarkImportCampaigns(b *testing.B) {
	data, err := NewCampaigns(902, 26, 50)

	if err != nil {
		panic(err)
	}

	ImportCampaigns(data)
}

func TestCampaignsSortedList_SearchByUser(t *testing.T) {
	//Arrange
	assert := assert.New(t)

	testCases := []struct {
		User                       *User
		ExpectedWinnerCampaignName string
	}{
		{
			User: &User{
				User: "u1000",
				Profile: map[string]string{
					"attr_A": "A5",
				},
			},
			ExpectedWinnerCampaignName: "campaign2",
		},
		{
			User: &User{
				User: "u2000",
				Profile: map[string]string{
					"attr_A": "A5",
					"attr_B": "A15",
				},
			},
			ExpectedWinnerCampaignName: "campaign2",
		},
		{
			User: &User{
				User: "u3000",
				Profile: map[string]string{
					"attr_A": "A15",
				},
			},
			ExpectedWinnerCampaignName: "campaign2",
		},
		{
			User: &User{
				User: "u4000",
				Profile: map[string]string{
					"attr_A": "A52",
				},
			},
			ExpectedWinnerCampaignName: "none",
		},
	}

	//Act
	expectedCampaignsSortedList := ImportCampaigns(testImportCampaignsData)
	for _, tc := range testCases {
		actualResult := expectedCampaignsSortedList.SearchByUser(tc.User)

		//Assert
		if tc.ExpectedWinnerCampaignName != "none" {
			assert.NotNil(actualResult)
			assert.Equal(tc.ExpectedWinnerCampaignName, actualResult.Name)
		} else {
			assert.Nil(actualResult)
		}

	}
}
