package model

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

//Randomizer mock always return max possible number. User for steady tests
type RandomizerMock struct {
}

func (r *RandomizerMock) Int63n(n int64) int64 {
	return n
}

var (
	testImportCampaignsData Campaigns
)

func TestMain(m *testing.M) {
	Rnd = &RandomizerMock{}
	testImportCampaignsData = Campaigns{
		&Campaign{
			Name:  "campaign1",
			Price: 0.25,
			Targets: []*Target{
				{
					Target:     "attr_A",
					Attributes: MustTargetAttributes(10, "A"),
				},
				{
					Target:     "attr_B",
					Attributes: MustTargetAttributes(17, "B"),
				},
			},
		},
		&Campaign{
			Name:  "campaign2",
			Price: 0.35,
			Targets: []*Target{
				{
					Target:     "attr_A",
					Attributes: MustTargetAttributes(20, "A"),
				},
			},
		},
		&Campaign{
			Name:  "campaign3",
			Price: 1.35,
			Targets: []*Target{
				{
					Target:     "attr_A",
					Attributes: MustTargetAttributes(3, "A"),
				},
			},
		},
	}

	code := m.Run()
	os.Exit(code)
}

func TestNewCampaigns(t *testing.T) {
	//Arrange
	expectedNumberOfCampaigns := 2
	expectedNumberOfTargets := 3
	expectedNumberOfTargetsAttributes := 4

	//Act
	actualResult, err := NewCampaigns(2, 3, 4)

	//Assert
	assert := assert.New(t)
	assert.NoError(err)
	assert.Len(actualResult, expectedNumberOfCampaigns)
	assert.Len(actualResult[0].Targets, expectedNumberOfTargets)
	assert.Len(actualResult[0].Targets[0].Attributes, expectedNumberOfTargetsAttributes)
}

func TestNewCampaign(t *testing.T) {
	//Arrange
	expectedCampaignName := "campaign99"
	expectedTargetName := "attr_B"
	expectedNumberOfTargets := 3
	expectedNumberOfTargetsAttributes := 4

	//Act
	actualResult, err := NewCampaign(99, 3, 4)

	//Assert
	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal(actualResult.Name, expectedCampaignName)
	assert.Len(actualResult.Targets, expectedNumberOfTargets)
	assert.Equal(actualResult.Targets[1].Target, expectedTargetName)
	assert.Len(actualResult.Targets[1].Attributes, expectedNumberOfTargetsAttributes)

}

func TestNewCampaignErrors(t *testing.T) {
	//Arrange
	testCases := []struct {
		numberOfCampaigns        int64
		numberOfTargets          int64
		numberOfTargetAttributes int64
		expectedError            error
	}{
		{-1, 1, 1, ErrCampaingIndexOutOfRange},
		{1, 1, -1, ErrTargetsAttributesOutOfRange},
		{1, 1, 101, ErrTargetsAttributesOutOfRange},
		{1, -1, 1, ErrTargetsOutOfRange},
		{1, 27, 1, ErrTargetsOutOfRange},
	}

	for _, tc := range testCases {
		//Act
		_, err := NewCampaign(tc.numberOfCampaigns, tc.numberOfTargets, tc.numberOfTargetAttributes)

		//Assert
		assert := assert.New(t)
		if assert.Error(err) {
			assert.EqualError(err, tc.expectedError.Error())
		}
	}
}
