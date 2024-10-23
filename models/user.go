package models

type User struct {
  ID       string `json:"-" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (u User) GetFirstNameLetter() string {
  return string(u.Username[0])
}

func (u User) Exercises() []Exercise {
  exercises, _ := EXERCISE.GetByUserID(u.ID)
  return exercises
}

