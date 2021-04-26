package models

import (
	"fmt"
	"net/mail"
	"strings"
	"errors"
	"time"
	"gorm.io/gorm"
	"gorm.io/datatypes"
)

// User definition
type Event struct {
	gorm.Model
	Email string `gorm:"type:varchar(100);unique" validate:"required unique" json:"email"`
	Environment string `gorm:"type:varchar(100)" json:"environment"`
	Component string `gorm:"type:varchar(100)" json:"component"`
	Message string `gorm:"type:varchar(100)" json:"message"`
	Data datatypes.JSON `json:"data"`
}


// Strip removes all whitespaces from the request body
func (e *Event) Strip() {
	e.Email = strings.TrimSpace(e.Email)
	e.Environment = strings.TrimSpace(e.Environment)
	e.Component = strings.TrimSpace(e.Component)
}

//Create adds a new historical event to the database
func (e *Event) Create(db *gorm.DB) (*Event, error) {
	var err error
	if !isEmailValid(e.Email) {
		return nil, errors.New("Email is invalid. Kindly check it")
	}
	err = db.Debug().Create(e).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func isEmailValid(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

// Retrieve retrieves all events that matches a filter map
func (e *Event) Retrieve(db *gorm.DB, filters map[string]interface{}) ([]*Event, error) {
	var (
		events []*Event
		createdAt time.Time
	)

	query := db.Model(&events)
	if val, ok := filters["date"]; ok {
		layout := "2-1-2006"
		createdAt, _ = time.Parse(layout, val.(string))
		delete(filters, "date")
		query.Where("created_at > ?", createdAt)
	}
	if val, ok := filters["text"]; ok {
		query.Where("message iLIKE ?", fmt.Sprintf("%v%v%v",  "%",val,"%"))
		delete(filters, "text")
	}

	query.Where(filters).Find(&events)
	return events, nil
}