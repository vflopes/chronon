package postgres

import (
	"log"
	"strconv"
)

func uuidToInt64(uuidStr string) int64 {

	numKey, err := strconv.ParseUint(uuidStr[19:23]+uuidStr[24:36], 16, 64)

	if err != nil {
		log.Fatalf("Invalid UUID to convert to Int64: %s (%v)", uuidStr, err)
	}

	return int64(numKey)

}
