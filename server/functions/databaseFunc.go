package functions

import (
	"strings"
	"time"

	"github.com/srisudarshanrg/go-fighter-jets/server/models"
)

// GetAllFighterJets gets all the fighter jets from the database
func GetAllFighterJets() ([]models.FighterJet, error) {
	getAllJetsQuery := `select * from jets`
	rows, err := db.Query(getAllJetsQuery)
	if err != nil {
		return nil, err
	}

	var jets []models.FighterJet
	for rows.Next() {
		var id, numberBuilt int
		var name, manufacturer, typeJet, year, features, role, country, imageLink string
		var generation float64
		var createdAt, updatedAt *time.Time

		err = rows.Scan(&id, &name, &manufacturer, &typeJet, &year, &features, &numberBuilt, &role, &country, &imageLink, &generation, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		jet := models.FighterJet{
			ID:           id,
			JetName:      name,
			Manufacturer: manufacturer,
			Type:         typeJet,
			Year:         year,
			Features:     features,
			NumberBuilt:  numberBuilt,
			Role:         role,
			Country:      country,
			ImageLink:    imageLink,
			Generation:   generation,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
		}
		jets = append(jets, jet)
	}

	return jets, nil
}

// GetFighterJetByID gets a fighter jet from the database using a given ID
func GetFighterJetByID(idArg int) (models.FighterJet, error) {
	getJetQuery := `select * from jets where id=$1`
	row := db.QueryRow(getJetQuery, idArg)

	var id, numberBuilt int
	var name, manufacturer, typeJet, year, features, role, country, imageLink string
	var generation float64
	var createdAt, updatedAt *time.Time

	err := row.Scan(&id, &name, &manufacturer, &typeJet, &year, &features, &numberBuilt, &role, &country, &imageLink, &generation, &createdAt, &updatedAt)
	if err != nil {
		return models.FighterJet{}, err
	}

	jet := models.FighterJet{
		ID:           id,
		JetName:      name,
		Manufacturer: manufacturer,
		Type:         typeJet,
		Year:         year,
		Features:     features,
		NumberBuilt:  numberBuilt,
		Role:         role,
		Country:      country,
		ImageLink:    imageLink,
		Generation:   generation,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}

	return jet, nil
}

// GetFighterJetBySearchKey gets all fighter jets in the database that match a key
func GetFighterJetsBySearchKey(key string) ([]models.FighterJet, error) {
	keyArg := "%" + strings.ToLower(key) + "%"
	getJetsQuery := `select * from jets where 
	lower(jet_name) like $1 
	or lower(manufacturer) like $1 
	or lower(type) like $1
	or year=$1
	or lower(features) like $1
	or lower(role) like $1
	or lower(country) like $1`

	rows, err := db.Query(getJetsQuery, keyArg)
	if err != nil {
		return nil, err
	}

	var results []models.FighterJet
	for rows.Next() {
		var id, numberBuilt int
		var name, manufacturer, typeJet, year, features, role, country, imageLink string
		var generation float64
		var createdAt, updatedAt *time.Time

		err = rows.Scan(&id, &name, &manufacturer, &typeJet, &year, &features, &numberBuilt, &role, &country, &imageLink, &generation, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		jet := models.FighterJet{
			ID:           id,
			JetName:      name,
			Manufacturer: manufacturer,
			Type:         typeJet,
			Year:         year,
			Features:     features,
			NumberBuilt:  numberBuilt,
			Role:         role,
			Country:      country,
			ImageLink:    imageLink,
			Generation:   generation,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
		}
		results = append(results, jet)
	}

	return results, nil
}
