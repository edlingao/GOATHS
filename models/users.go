package models

import (
	"errors"
	"exercise-app/libs"
	"log"
)

type UserModel string

const USER UserModel = "USER_MODEL"

func (u UserModel) GetByUsername(username string) (User, error) {
  var user User
  libs.ConnectDB()
  defer libs.DisconnectDB()

  err := libs.DB.QueryRowx(`
    SELECT id, email, username, password
    FROM users
    WHERE username = ?
  `, username).StructScan(&user)

  if err != nil {
    log.Print(err)

    return user, err
  }

  return user, nil
}

func (u UserModel) GetByEmail(email string) ( User, error ) {
  var user User
  libs.ConnectDB()
  defer libs.DisconnectDB()
  
  err := libs.DB.QueryRowx(`
    SELECT id, email, username, password
    FROM users
    WHERE email = ?
  `, email).StructScan(&user)

  if err != nil {
    log.Print(err)
    
    return user, err
  }
  
  return user, nil
}

func (u UserModel) Get(property string, value string) ( User, error ) {
  var user User
  libs.ConnectDB()
  defer libs.DisconnectDB()
  
  err := libs.DB.QueryRowx(`
    SELECT id, email, username, password
    FROM users
    WHERE ? = ?
  `, property, value).StructScan(&user)

  if err != nil {
    log.Print(err)
    
    return user, err
  }
  
  return user, nil
}

func (u UserModel) Register(email string, username string, password string) (User, error) {
	var user User
	libs.ConnectDB()
	defer libs.DisconnectDB()

	encypted_password := EncryptPassword(password)

	row, err := libs.DB.Exec(`
    INSERT INTO users (email, username, password)
    VALUES(?, ?, ?)
  `, email, username, encypted_password)

	if err != nil {
		log.Print(err)
		return User{}, err
	}

	lastInsertedId, err := row.LastInsertId()

	if err != nil {
		log.Print(err)
		return User{}, err
	}

  error := libs.DB.QueryRowx(`
    SELECT id, email, username, password
    FROM users
    WHERE id = ?
  `, lastInsertedId).StructScan(&user)

	if error != nil {
		log.Print(error)
		return User{}, error
	}

	return user, nil
}

func (u UserModel) Login(email string, password string) ( User, error ) {
  user, err := u.GetByEmail(email)
  valid := ValidatePassword(user.Password, password)

  if !valid || err != nil {
    return User{}, errors.New("Invalid Username or password")
  }
  
  return user, nil
}

