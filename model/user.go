package model

import (
	"math"
	"strconv"
)

type User struct {
	User    string            `json:"user"`
	Profile map[string]string `json:"profile"`
}

func NewUser(counter int64) (*User, error) {

	profile := map[string]string{}

	profileLength := int64(math.Mod(float64(counter), float64(26)))
	if profileLength == 0 {
		profileLength = 26
	}

	//pretty.Println("=============    (╯°□°）╯︵ ┻━┻)   =============")
	//pretty.Println(profileLength)
	//pretty.Println("=============    ┬─┬ノ( º _ ºノ)   =============")

	for i := int64(1); i <= profileLength; i++ {
		char, err := getCharByIndex(i)
		if err != nil {
			return nil, err
		}
		RndMutex.Lock()
		profile["attr_"+char] = char + strconv.FormatInt(Rnd.Int63n(200), 10)
		RndMutex.Unlock()
	}

	user := &User{
		User:    "u" + strconv.FormatInt(counter, 10),
		Profile: profile,
	}

	return user, nil
}
