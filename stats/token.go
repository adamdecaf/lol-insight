package stats

import (
	"github.com/TrevorSStone/goriot"
	"io/ioutil"
	"os"
)

var (
	// defaults
	DefaultRiotApiTokenPath = "./.riot-api"
	DefaultApiToken = NewToken()
)

type ApiToken struct {
	Token string
}

func (a ApiToken) setToken() {
	goriot.SetAPIKey(a.Token)
}

func NewToken() ApiToken {
	token := ""

	// read from env variable
	if t := os.Getenv("RIOT_API_TOKEN"); t != "" {
		token = t
	}

	// read from file path
	r, err := os.Open(DefaultRiotApiTokenPath)
	if err == nil {
		bytes, err := ioutil.ReadAll(r)
		if err == nil {
			token = string(bytes)
		}
	}

	tt := ApiToken{
		Token: token,
	}

	tt.setToken()

	return tt
}
