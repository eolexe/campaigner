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

func newTargets(numberOfTargets int64, numberOfTargetAttributes int64) (Targets, error) {
	if 26 < numberOfTargets || numberOfTargets < 0 {
		return nil, ErrTargetsOutOfRange
	}

	targetList := Targets{}

	RndMutex.Lock()
	numberOfTargets = Rnd.Int63n(numberOfTargets)
	RndMutex.Unlock()

	for i := int64(1); i <= numberOfTargets; i++ {
		char, err := getCharByIndex(i)
		if err != nil {
			return nil, err
		}

		targetAttributes, err := newTargetAttributes(numberOfTargetAttributes, char)

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

func MustTargetAttributes(numberOfTargetAttributes int64, prefix string) []string {
	result, err := newTargetAttributes(numberOfTargetAttributes, prefix)
	if err != nil {
		panic(err)
	}

	return result
}

func newTargetAttributes(numberOfTargetAttributes int64, prefix string) ([]string, error) {
	if 100 < numberOfTargetAttributes || numberOfTargetAttributes < 0 {
		return nil, ErrTargetsAttributesOutOfRange
	}

	var targetAttributes []string

	RndMutex.Lock()
	numberOfTargetAttributes = Rnd.Int63n(numberOfTargetAttributes)
	RndMutex.Unlock()

	for i := int64(0); i < numberOfTargetAttributes; i++ {
		targetAttributes = append(targetAttributes, prefix+strconv.FormatInt(int64(i), 10))
	}

	return targetAttributes, nil
}
