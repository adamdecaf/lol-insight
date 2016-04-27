package stats

import (
	"fmt"
	"github.com/TrevorSStone/goriot"
	"log"
)

func FindSummonerViaName(name string) (*goriot.Summoner, error) {
	log.Printf("finding summoner by name '%s'\n", name)

	summoners, err := goriot.SummonerByName(goriot.NA, name)
	if err != nil {
		log.Fatalf("%s\n", err)
		return nil, err
	}

	s := summoners[name]
	fmt.Println(s)
	return &s, nil
}
