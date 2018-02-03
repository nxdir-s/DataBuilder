package main

import (
	obj "dataBuilder/objects"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func ConsumeMatchData(body []byte) error {
	var matchData obj.MatchData
	err := json.Unmarshal(body, &matchData)
	if err != nil {
		return errors.Wrap(err, "Error unmarshalling to type MatchDto in ConsumeMatchData")
	}

	summonerIdentity, err := FindSummonerIdentity(matchData)
	if err != nil {
		return errors.Wrap(err, "Error in ConsumeMatchData")
	}

	summonerStats, err := FindSummonerStats(matchData, summonerIdentity)
	if err != nil {
		return errors.Wrap(err, "Error in ConsumeMatchData")
	}

	//Need to do something with stats now

	log := fmt.Sprintf("MatchData: %v", summonerStats)

	fmt.Println(log)

	return nil
}

func FindSummonerStats(data obj.MatchData, summonerIdentity *obj.ParticipantIdentityDto) (*obj.ParticipantStatsDto, error) {
	match := data.Match

	stats := new(obj.ParticipantStatsDto)

	//Finding the participant stats using the participant id
	for _, participant := range match.Participants {
		if participant.ParticipantId == summonerIdentity.ParticipantId {
			stats = &participant.Stats
		}
	}

	if stats == nil {
		return nil, errors.New("Could not find summonerStats in FindSummonerStats")
	}

	return stats, nil
}

func FindSummonerIdentity(data obj.MatchData) (*obj.ParticipantIdentityDto, error) {
	participants := data.Match.ParticipantIdentities
	identity := new(obj.ParticipantIdentityDto)

	//Loops through Participant identities to to get a match on Summoner Name
	//Since Summoner Name is in recieved data we need to get the participant id based on it
	for _, participant := range participants {
		if participant.Player.SummonerName == data.SummonerName {
			identity.Player = participant.Player
			identity.ParticipantId = participant.ParticipantId
		}
	}

	if identity == nil {
		return nil, errors.New("Could not find summoner in participant identities")
	}

	return identity, nil
}
