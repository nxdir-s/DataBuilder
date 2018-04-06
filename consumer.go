package main

import (
	obj "dataBuilder/objects"
	util "dataBuilder/util"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

const (
	accountId int64 = 47440170
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

	summonerStats, participantDto, err := FindSummonerStats(matchData, summonerIdentity)
	if err != nil {
		return errors.Wrap(err, "Error in ConsumeMatchData")
	}

	accId := summonerIdentity.Player.AccountId

	stats := *summonerStats

	sumStats := obj.SummonerMatchStats{
		AccountId:  accId,
		SeasonId:   matchData.Match.SeasonId,
		QueueId:    matchData.Match.QueueId,
		MatchId:    matchData.Match.GameId,
		ChampionId: participantDto.ChampionId,
		Stats:      stats,
		Timeline:   participantDto.Timeline,
	}

	//Need to do something with stats now
	jsonLog, _ := json.Marshal(sumStats)

	log := fmt.Sprintf("MatchData: %v", string(jsonLog))
	fmt.Println(log)

	err = util.Publish(sumStats, "dataApi/summonerStats")
	if err != nil {
		return errors.Wrap(err, "Error in ConsumeMatchData")
	}

	return nil
}

func FindSummonerStats(data obj.MatchData, summonerIdentity *obj.ParticipantIdentityDto) (*obj.ParticipantStatsDto, *obj.ParticipantDto, error) {
	match := data.Match

	stats := new(obj.ParticipantStatsDto)
	participantInfo := new(obj.ParticipantDto)

	//Finding the participant stats using the participant id
	for _, participant := range match.Participants {
		if participant.ParticipantId == summonerIdentity.ParticipantId {
			stats = &participant.Stats
			participantInfo = &participant
			break
		}
	}

	if stats == nil {
		return nil, nil, errors.New("Could not find summonerStats in FindSummonerStats")
	}

	return stats, participantInfo, nil
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
