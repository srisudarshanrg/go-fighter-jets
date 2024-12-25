package functions

import (
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
		var jetName, manufacturer, typeJet, year, features, role, country, imageLink, generation string
		var createdAt, updatedAt *time.Time

		err = rows.Scan(&id, &jetName, &manufacturer, &typeJet, &year, &features, &numberBuilt, &role, &country, &imageLink, &generation, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		jet := models.FighterJet{
			ID:           id,
			JetName:      jetName,
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
