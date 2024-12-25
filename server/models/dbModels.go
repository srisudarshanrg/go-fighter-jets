package models

import "time"

// FighterJet is the struct for a jet in the database
type FighterJet struct {
	ID           int
	JetName      string
	Manufacturer string
	Type         string
	Year         string
	Features     string
	NumberBuilt  int
	Role         string
	Country      string
	ImageLink    string
	Generation   string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
