package objects

type MatchData struct {
	Match        MatchDto `json:"match"`
	SummonerName string   `json:"summonerName"`
	Timestamp    int64    `json:"timestamp"`
}

type SummonerMatchStats struct {
	AccountId  int64                  `json:"accountId"`
	Timestamp  int64                  `json:"timestamp"`
	SeasonId   int                    `json:"seasonId"`
	QueueId    int                    `json:"queueId"`
	MatchId    int64                  `json:"matchId"`
	ChampionId int                    `json:"championId"`
	Stats      ParticipantStatsDto    `json:"stats"`
	Timeline   ParticipantTimelineDto `json:"timeline"`
}
