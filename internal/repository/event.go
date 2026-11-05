package repository

import (
	"github.com/pankajy-dev/go-rest-app/internal/db"
	"github.com/pankajy-dev/go-rest-app/internal/models"
)

func SaveEvent(e *models.Event) error {
	query := `
	INSERT INTO event (name, description, location, dateTime, user_id)
	values (?, ?, ?, ?, ?)
	`
	// Prepares the query to be executed
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("Cannot insert record in DB.")
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

func GetAllEvents() ([]models.Event, error) {
	query := "SELECT * FROM event"

	// Queryb is used when we do select
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var event models.Event
		// This works like scan of fmt, all the columns need to be specified in same order in which they exist in DB
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventId(id int64) (*models.Event, error) {
	query := `
		SELECT * FROM event
		WHERE id = ?
	`

	var event models.Event

	row := db.DB.QueryRow(query, id)

	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func UpdateEvent(event models.Event) error {
	query := `
		UPDATE event
		SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
		WHERE id = ?
	`

	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.UserId, event.Id)
	return err
}

func DeleteEvent(event models.Event) error {
	query := `
		DELETE FROM event
		WHERE id = ?
	`

	_, err := db.DB.Exec(query, event.Id)
	return err
}

func RegisterForEvent(event models.Event, userId int64) error {
	query := "INSERT INTO registration (event_id, user_id) VALUES (?, ?)"

	_, err := db.DB.Exec(query, event.Id, userId)
	return err
}

func CancelRegistration(event models.Event, userId int64) error {
	query := "DELETE FROM registration WHERE event_id = ? and user_id = ?"

	_, err := db.DB.Exec(query, event.Id, userId)
	return err
}
