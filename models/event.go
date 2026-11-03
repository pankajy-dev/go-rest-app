package models

import (
	"time"
)

type Event Struct {
	Id int
	Name string
	Description string
	Location string
	DateTime time.Time
	userId int
}

var events := []Events

func (e Event) save(){
	//later store it in DB
	events = append(events, e)
}