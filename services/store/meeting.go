package store

type Meeting struct {
	Title         string
	Participators []string
	StartTime     string
	EndTime       string
}

type Meetings []Meeting

func CreateMeeting(meeting *Meeting) error {
	client, err := GetClient()
	if err != nil {
		return err
	}
	if err := client.Commit(); err != nil {
		return err
	}
	client.Dump()
	return nil
}
