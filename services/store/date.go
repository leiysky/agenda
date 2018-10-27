package store

import (
	"errors"
	"fmt"
	"regexp"
)

type DateType struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

// NewDate can parse a string into DateType
// Date format is similar to UNIX time
// Format: YYYY:MM:DDThh:mm
// Example: 2018-10-24T13:59
func NewDate(dateStr string) (*DateType, error) {
	if istrue, err := regexp.Match("[0-9]{4}-((0[1-9]|1[0-2])-(0[1-9]|1[0-9]|2[0-8])|(0[13-9]|1[0-2])-(29|30)|(0[13578]|1[02])-31)T([01][0-9]|2[0-3]):[0-5][0-9]", []byte(dateStr)); istrue == false {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("invalid date string")
	}
	date := new(DateType)
	if _, err := fmt.Sscanf(dateStr, "%d-%d-%dT%d:%d", &date.Year, &date.Month, &date.Day, &date.Hour, &date.Minute); err != nil {
		return nil, err
	}
	return date, nil
}

func DateToString(date DateType) string {
	str := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d", date.Year, date.Month, date.Day, date.Hour, date.Minute)
	return str
}

func (date DateType) Lt(date1 DateType) bool {
	if date.Year == date1.Year {
		if date.Month == date1.Month {
			if date.Day == date1.Day {
				if date.Hour == date1.Hour {
					if date.Minute == date1.Minute {
						return false
					} else if date.Minute > date1.Minute {
						return false
					} else {
						return true
					}
				} else if date.Hour > date1.Hour {
					return false
				} else {
					return true
				}
			} else if date.Day > date1.Day {
				return false
			} else {
				return true
			}
		} else if date.Month > date1.Month {
			return false
		} else {
			return true
		}
	} else if date.Year > date1.Year {
		return false
	} else {
		return true
	}
}

func (date DateType) Gt(date1 DateType) bool {
	if date.Year == date1.Year {
		if date.Month == date1.Month {
			if date.Day == date1.Day {
				if date.Hour == date1.Hour {
					if date.Minute == date1.Minute {
						return false
					} else if date.Minute < date1.Minute {
						return false
					} else {
						return true
					}
				} else if date.Hour < date1.Hour {
					return false
				} else {
					return true
				}
			} else if date.Day < date1.Day {
				return false
			} else {
				return true
			}
		} else if date.Month < date1.Month {
			return false
		} else {
			return true
		}
	} else if date.Year < date1.Year {
		return false
	} else {
		return true
	}
}

func (date DateType) Eq(date1 DateType) bool {
	if date.Year == date1.Year {
		if date.Month == date1.Month {
			if date.Day == date1.Day {
				if date.Hour == date1.Hour {
					if date.Minute == date1.Minute {
						return true
					}
				}
			}
		}
	}
	return false
}

func (date DateType) Between(start, end DateType) bool {
	if (date.Lt(end) && date.Gt(start)) || date.Eq(start) || date.Eq(end) {
		return true
	}
	return false
}
