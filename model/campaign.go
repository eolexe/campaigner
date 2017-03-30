package model

import (
	"errors"
	"strconv"
)

var (
	ErrCampaingIndexOutOfRange = errors.New("campaign index must be positive value")
)

type Campaign struct {
	Name    string  `json:"campaign_name"`
	Price   float64 `json:"price"`
	Targets Targets `json:"target_list"`
}

type Campaigns []*Campaign

func NewCampaigns(numberOfCampaigns int64, numberOfTargets int64, numberOfTargetAttributes int64) (Campaigns, error) {
	result := Campaigns{}

	for index := int64(1); index <= numberOfCampaigns; index++ {
		campaign, err := NewCampaign(index, numberOfTargets, numberOfTargetAttributes)

		if err != nil {
			return nil, err
		}

		result = append(result, campaign)
	}

	return result, nil
}

func NewCampaign(index int64, numberOfTargets int64, numberOfTargetAttributes int64) (*Campaign, error) {
	if index < 0 {
		return nil, ErrCampaingIndexOutOfRange
	}

	targets, err := newTargets(numberOfTargets, numberOfTargetAttributes)

	if err != nil {
		return nil, err
	}

	RndMutex.Lock()
	price := float64(Rnd.Int63n(10000)) * 0.01
	RndMutex.Unlock()

	return &Campaign{
		Name:    "campaign" + strconv.FormatInt(index, 10),
		Price:   price,
		Targets: targets,
	}, nil
}
