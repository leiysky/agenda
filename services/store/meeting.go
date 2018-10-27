package store

import (
	"errors"
)

type MeetingType struct {
	Title         string   `json:"title"`
	Participators []string `json:"participators"`
	StartTime     DateType `json:"start_time"`
	EndTime       DateType `json:"end_time"`
}

type MeetingsType []MeetingType

func CreateMeeting(meeting MeetingType) error {
	client, err := GetClient()
	if err != nil {
		return err
	}
	meetings := client.getMeetings()
	meetings = append(meetings, meeting)
	client.setMeetings(meetings)
	if err := client.Commit(); err != nil {
		return err
	}
	return client.Dump()
}

func (cl *ClientType) getMeetings() MeetingsType {
	return cl.DB.Collection.Meetings
}

func (cl *ClientType) setMeetings(meetings MeetingsType) {
	temp := make(MeetingsType, len(meetings))
	copy(temp, meetings)
	cl.DB.Collection.Meetings = temp
}

func (meeting *MeetingType) AddParticipator(participator string) error {
	user := UserType{participator, ""}
	if !user.IsExist() {
		return errors.New(participator + " is a invalid user")
	}
	for _, one := range meeting.Participators {
		if one == user.Username {
			return errors.New(participator + " has been in the meeting")
		}
	}
	if !user.CanTakePartIn(meeting) {
		return errors.New("there is a confilict of time when adding user " + participator)
	}
	meeting.Participators = append(meeting.Participators, participator)
	return nil
}
