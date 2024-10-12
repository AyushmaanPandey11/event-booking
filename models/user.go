package models

import (
	"eventBooking.com/m/db"
)

type User struct {
	Id       int64
	name     string `binding:"required"`
	email    string `binding:"required"`
	password string `binding:"required"`
}

func (u User) Save() error {
	query := `
		INSERT INTO users(name,email,password) 
		VALUES (?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.name, u.email, u.password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.Id = id
	return err
}

func (u User) Update() error {
	query := `
		UPDATE users SET name = ?,email = ?,password = ?
		WHERE Id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.name, u.email, u.password)
	return err
}

func (u User) Delete() error {
	query := `
		DELETE FROM users WHERE id=?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Id)
	return err
}

func GetAllUsers() ([]User, error) {
	query := `SELECT (id,name,email) FROM users`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var allUsers []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.name, &user.email)
		if err != nil {
			return nil, err
		}
		allUsers = append(allUsers, user)
	}
	return allUsers, nil
}

func GetUserById(id int64) (*User, error) {
	query := `SELECT (id,name,email) FROM users WHERE id =?`
	stmt, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var user User
	err = stmt.Scan(&user.Id, &user.name, &user.email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
