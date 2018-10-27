package store

import "errors"

type UserType struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersType []UserType

func CreateUser(user UserType) error {
	client, err := GetClient()
	if err != nil {
		return err
	}
	users := client.getUsers()
	users = append(users, user)
	client.setUsers(users)
	if err := client.Commit(); err != nil {
		return err
	}
	return client.Dump()
}

func (cl *ClientType) getUsers() UsersType {
	return cl.DB.Collection.Users
}

func (cl *ClientType) setUsers(users UsersType) {
	temp := make(UsersType, len(users))
	copy(temp, users)
	cl.DB.Collection.Users = temp
}

func (user UserType) CanTakePartIn(meeting *MeetingType) bool {
	curMeetings := user.GetCurrentMeetings()
	for _, one := range curMeetings {
		if meeting.StartTime.Between(one.StartTime, one.EndTime) || meeting.EndTime.Between(one.StartTime, one.EndTime) {
			return false
		} else if one.StartTime.Between(meeting.StartTime, meeting.EndTime) || one.EndTime.Between(meeting.StartTime, meeting.EndTime) {
			return false
		}
	}
	return true
}

func (user UserType) GetCurrentMeetings() MeetingsType {
	client, _ := GetClient()
	meetings := client.getMeetings()
	curMeetings := MeetingsType{}
	for _, one := range meetings {
		for _, participator := range one.Participators {
			if participator == user.Username {
				curMeetings = append(curMeetings, one)
			}
		}
	}
	return curMeetings
}

func (user UserType) IsExist() bool {
	client, _ := GetClient()
	users := client.getUsers()
	for _, one := range users {
		if one.Username == user.Username {
			return true
		}
	}
	return false
}

func GetAllUsers() (UsersType, error) {
	client, err := GetClient()
	if err != nil {
		return nil, err
	}
	users := client.getUsers()
	return users, nil
}

func DeleteUserByName(name string) error {
	client, err := GetClient()
	if err != nil {
		return err
	}
	users := client.DB.Collection.Users
	isExist := false
	for idx, one := range users {
		if one.Username == name {
			users = append(users[:idx], users[idx+1:]...)
			isExist = true
			break
		}
	}
	if !isExist {
		return errors.New("the user doesn't exist")
	}

	meetings := client.getMeetings()
	isExist = false
	for idx, one := range meetings {
		if one.Participators[0] == name {
			meetings = append(meetings[:idx], meetings[idx+1:]...)
			isExist = true
			break
		}
	}

	client.DB.Collection.Users = users
	if err := client.Commit(); err != nil {
		return err
	}
	return client.Dump()
}
