package store

import (
	"fmt"
)

type Meeting struct {
	Title         string
	Participators []string
	StartTime     string
	EndTime       string
}

func CreateMeeting(meeting *Meeting) error {
	fmt.Printf("Title: %s\n", meeting.Title)
	return nil
}
