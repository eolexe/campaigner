package model

import (
	"errors"
	"strconv"
)

var (
	ErrTargetsOutOfRange           = errors.New("requested number of targets is out of range, set numebr between 0 – 26")
	ErrTargetsAttributesOutOfRange = errors.New("requested number of targets attributes is out of range, set number between 0 – 100")
)

type Target struct {
	Target     string   `json:"target"`
	Attributes []string `json:"attr_list"`
}

type Targets []*Target

func NewTargets(numberOfTargets int64, numberOfTargetAttributes int64) (Targets, error) {
	if 26 < numberOfTargets || numberOfTargets < 0 {
		return nil, ErrTargetsOutOfRange
	}

	targetList := Targets{}

	numberOfTargets = Randomizer.Int63n(numberOfTargets)

	for i := int64(1); i <= numberOfTargets; i++ {
		char, err := getCharByIndex(i)
		if err != nil {
			return nil, err
		}

		targetAttributes, err := NewTargetAttributes(numberOfTargetAttributes, char)

		if err != nil {
			return nil, err
		}

		targetList = append(targetList, &Target{
			Target:     "attr_" + char,
			Attributes: targetAttributes,
		})

	}

	return targetList, nil
}

func NewTargetAttributes(numberOfTargetAttributes int64, prefix string) ([]string, error) {
	if 100 < numberOfTargetAttributes || numberOfTargetAttributes < 0 {
		return nil, ErrTargetsAttributesOutOfRange
	}

	var targetAttributes []string

	numberOfTargetAttributes = Randomizer.Int63n(numberOfTargetAttributes)

	for i := int64(0); i < numberOfTargetAttributes; i++ {
		targetAttributes = append(targetAttributes, prefix+strconv.FormatInt(int64(i), 10))
	}

	return targetAttributes, nil
}
