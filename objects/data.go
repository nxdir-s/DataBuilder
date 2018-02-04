package objects

type MatchData struct {
	Match        MatchDto `json:"match"`
	SummonerName string   `json:"summonerName"`
}

type SummonerMatchStats struct {
	AccountId int64               `json:"accountId"`
	MatchId   int64               `json:"matchId"`
	Stats     ParticipantStatsDto `json:"stats"`
}
