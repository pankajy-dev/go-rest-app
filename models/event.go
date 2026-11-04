package models

import (
	"time"

	"github.com/pankajy-dev/go-rest-app/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:required`
	Description string    `binding:required`
	Location    string    `binding:required`
	DateTime    time.Time `binding:required`
	UserId      int
}

func (e Event) Save() error {
	//later store it in DB

	query := `
	INSERT INTO event (name, description, location, dateTime, user_id) 
	values (?, ?, ?, ?, ?)
	`
	// Prepares the query to be executed
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("Cannot insert record in DB.")
		return err
	}
	defer stmt.Close()
	// Pass the arguments for each '?'
	// We use exec when we change the data
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.Id = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM event"

	// Queryb is used when we do select
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var events []Event

	for rows.Next() {
		var event Event
		// This works like scan of fmt, all the columns need to be specified in same order in which they exist in DB
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventId(id int64) (*Event, error) {
	query := `
		SELECT * FROM event
		WHERE id = ?
	`

	var event Event

	row := db.DB.QueryRow(query, id)

	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `
		UPDATE event 
		SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
		WHERE id = ?
	`

	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.UserId, event.Id)
	return err

}

func (event Event) DeleteEvent() error {
	query := `
		DELETE FROM event
		WHERE id = ?
	`

	_, err := db.DB.Exec(query, event.Id)
	return err
}
