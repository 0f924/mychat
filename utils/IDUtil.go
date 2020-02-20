package utils

import (
	"github.com/sony/sonyflake"
	"log"
	"strconv"
)

func GetRandomId() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	return strconv.Itoa(int(id))
}
