package models

import (
	"strconv"
	"time"
)

type Exercise struct {
	ID       int       `json:"id" db:"id"`
	User_ID  string    `json:"user_id" db:"user_id"`
	Date     time.Time `json:"date" db:"date"`
	Feeling  string    `json:"feeling" db:"feeling"`
  Hours    int    `json:"hours" db:"hours"`
  Minutes  int    `json:"minutes" db:"minutes"`
	Calories int       `json:"calories" db:"calories"`
	User     User      `json:"user"`
}

func (e Exercise) GetID() string {
	return strconv.FormatInt(int64(e.ID), 10)
}

func (e Exercise) GetCalories() string {
	return strconv.FormatInt(int64(e.Calories), 10)
}

func (e Exercise) GetDuration() string {
  hours := strconv.FormatInt(int64(e.Hours), 10)
  minutes := strconv.FormatInt(int64(e.Minutes), 10)

  if e.Minutes < 10 {
    minutes = "0" + minutes
  }

  return hours + "h " + minutes + "m"
}

func (e Exercise) GetFormattedDate() string {
	return e.Date.Format("02/01/06")
}

func (e Exercise) GetFeeling() string {
	feeling := "😌 Relaxed"

	switch e.Feeling {
	case "normal":
		feeling = "🙂 Normal"
	case "focused":
		feeling = "😎 Focused"
	case "inspired":
		feeling = "🚀 Inspired"
	case "stressed":
		feeling = "😵 Stressed"
	}

	return feeling
}
