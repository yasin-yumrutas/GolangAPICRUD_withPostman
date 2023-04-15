package models

import (
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreateOn    time.Time `json:"createon"`
	ChangeOn    time.Time `json:"changeon"`
}
