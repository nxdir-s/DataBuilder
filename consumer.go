package main

import (
	obj "dataBuilder/objects"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func ConsumeMatchData(body []byte) error {
	var match obj.MatchDto
	err := json.Unmarshal(body, &match)
	if err != nil {
		return errors.Wrap(err, "Error unmarshalling to type MatchDto in ConsumeMatchData")
	}

	log := fmt.Sprintf("MatchData: %v", match)

	fmt.Println(log)

	return nil
}
