package models

import (
	"time"

	"eventBooking.com/m/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	User_id     int `binding:"required"`
}

func (e Event) Save() error {
	query := `
	INSERT INTO events(Name,Description,Location,DateTime,User_id)
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.User_id)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
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
	_, err = stmt.Exec(event.ID, event.Name, event.Description, event.Location, event.DateTime, event.User_id, event.ID)
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
	_, err = stmt.Exec(event.ID)
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
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.User_id)
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
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.User_id)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
