package models

import (
	"time"

	"eventBooking.com/m/db"
)

type Event struct {
	Id          int64
	name        string `binding:"required"`
	description string `binding:"required"`
	location    string `binding:"required"`
	dateTime    time.Time
	User_id     int64 `binding:"required"`
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(Name,Description,Location,DateTime,User_id)
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.name, e.description, e.location, e.dateTime, e.User_id)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.Id = id
	return err
}

func (event Event) Update() error {
	query := `
		UPDATE events SET Name= ?, Description = ?, Location=?, DateTime=?, User_id=?
		WHERE ID = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.name, event.description, event.location, event.dateTime, event.User_id, event.Id)
	return err
}

func (event Event) Delete() error {
	query := `
		DELETE FROM events WHERE ID = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Id)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.name, &event.description, &event.location, &event.dateTime, &event.User_id)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
func GetEventById(eventId int64) (*Event, error) {
	query := "SELECT * FROM events WHERE ID = ?"
	row := db.DB.QueryRow(query, eventId)
	var event Event
	err := row.Scan(&event.Id, &event.name, &event.description, &event.location, &event.dateTime, &event.User_id)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e Event) UserRegistration(userId int64) error {
	query := `
	INSERT INTO registrations(user_id,event_id)
	VALUES (?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId, e.Id)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := `
	DELETE FROM registrations WHERE user_id = ? AND event_id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId, e.Id)
	return err
}
