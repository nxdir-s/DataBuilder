package main

import (
	util "dataBuilder/util"
)

func main() {
	forever := make(chan bool)
	go util.RegisterConsumer("dataBuilder/matchData", ConsumeMatchData)
	go util.RegisterConsumer("dataBuilder/buildDashboard", ConsumeDashboard)
	<-forever
}
