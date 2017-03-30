package model

type CampaignsSortedList map[string]*Campaign

func (c CampaignsSortedList) Add(key string, campaign *Campaign) {

	if foundCampaign, ok := c[key]; !ok {
		c[key] = campaign
	} else if foundCampaign.Price < campaign.Price {
		c[key] = campaign
	}
}

func (c CampaignsSortedList) SearchByUser(user *User) *Campaign {

	var foundCampaign *Campaign
	for k, v := range user.Profile {
		key := k + "_" + v
		if campaign, ok := c[key]; ok {
			if foundCampaign == nil {
				foundCampaign = campaign
			}

			if foundCampaign != nil && foundCampaign.Price < campaign.Price {
				foundCampaign = campaign
			}
		}
	}

	return foundCampaign
}

func ImportCampaigns(campaigns Campaigns) CampaignsSortedList {
	result := make(chan CampaignsSortedList)

	proccessFn := func(campaigns Campaigns, result chan CampaignsSortedList) {
		campaignSortedList := CampaignsSortedList{}

		for _, campaign := range campaigns {
			for _, target := range campaign.Targets {
				for _, attribure := range target.Attributes {
					key := target.Target + "_" + attribure
					campaignSortedList.Add(key, campaign)
				}
			}
		}
		result <- campaignSortedList
	}

	batchSize := 100
	numberOfBatches := len(campaigns) / batchSize
	if numberOfBatches == 0 {
		numberOfBatches = 1
	}

	for i := 1; i <= numberOfBatches; i++ {

		start := (i - 1) * batchSize
		end := i * batchSize
		if i == numberOfBatches {
			end = len(campaigns)
		}

		//log.Printf(">>>>>batch: %d - %d \n", start, end)
		go proccessFn(campaigns[start:end], result)
	}

	resultingCampaignSortedList := <-result

	for i := 1; i <= numberOfBatches-1; i++ {
		for key, campaign := range <-result {
			resultingCampaignSortedList.Add(key, campaign)
		}
	}

	return resultingCampaignSortedList
}
