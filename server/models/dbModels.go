package models

import "time"

// FighterJet is the struct for a jet in the database
type FighterJet struct {
	ID           int
	JetName      string
	Manufacturer string
	Type         string
	Year         int
	Features     string
	NumberBuilt  int
	Role         string
	Country      string
	ImageLink    string
	Generation   float64
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
