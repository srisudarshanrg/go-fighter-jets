package functions

import (
	"sort"
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
		var id, numberBuilt, year int
		var name, manufacturer, typeJet, features, role, country, imageLink string
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

	var id, numberBuilt, year int
	var name, manufacturer, typeJet, features, role, country, imageLink string
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
	or lower(features) like $1
	or lower(role) like $1
	or lower(country) like $1`

	rows, err := db.Query(getJetsQuery, keyArg)
	if err != nil {
		return nil, err
	}

	var results []models.FighterJet
	for rows.Next() {
		var id, numberBuilt, year int
		var name, manufacturer, typeJet, features, role, country, imageLink string
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

func GetTwoFighterJetsForCompare(fighterJet1 string, fighterJet2 string) ([]models.FighterJet, error) {
	getFirstFighterJet := `select * from jets where jet_name=$1`
	row1 := db.QueryRow(getFirstFighterJet, fighterJet1)

	getSecondFighterJet := `select * from jets where jet_name=$1`
	row2 := db.QueryRow(getSecondFighterJet, fighterJet2)

	var id, numberBuilt, year int
	var name, manufacturer, typeJet, features, role, country, imageLink string
	var generation float64
	var createdAt, updatedAt *time.Time

	err := row1.Scan(&id, &name, &manufacturer, &typeJet, &year, &features, &numberBuilt, &role, &country, &imageLink, &generation, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	jet1 := models.FighterJet{
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

	err = row2.Scan(&id, &name, &manufacturer, &typeJet, &year, &features, &numberBuilt, &role, &country, &imageLink, &generation, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	jet2 := models.FighterJet{
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

	return []models.FighterJet{jet1, jet2}, nil
}

// GetDistinctGeneration gets distinct generations from the database
func GetDistinctGeneration() ([]float64, error) {
	getDistinctGenerationQuery := `select distinct generation from jets`
	rows, err := db.Query(getDistinctGenerationQuery)
	if err != nil {
		return nil, err
	}

	var generationsList []float64
	for rows.Next() {
		var generation float64

		err = rows.Scan(&generation)
		if err != nil {
			return nil, err
		}

		generationsList = append(generationsList, generation)
	}

	sort.Float64s(generationsList)

	return generationsList, nil
}

// GetJetsByGeneration gets fighter jets according to a specific generation
func GetJetsByGeneration(generation float64) ([]models.FighterJet, error) {
	getJetsByGenerationQuery := `select * from jets where generation=$1`
	rows, err := db.Query(getJetsByGenerationQuery, generation)
	if err != nil {
		return nil, err
	}

	var jets []models.FighterJet
	for rows.Next() {
		var id, numberBuilt, year int
		var name, manufacturer, typeJet, features, role, country, imageLink string
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

// GetJetsByCentury gets jets from the 20th and 21st century and returns them as separate lists
func GetJetsByCentury() ([]models.FighterJet, []models.FighterJet, error) {
	get20thQuery := `select * from jets where year < 2000`
	rows20th, err := db.Query(get20thQuery)
	if err != nil {
		return nil, nil, err
	}

	get21stQuery := `select * from jets where year >= 2000`
	rows21st, err := db.Query(get21stQuery)
	if err != nil {
		return nil, nil, err
	}

	var jets20th []models.FighterJet
	var jets21st []models.FighterJet

	for rows20th.Next() {
		var id, numberBuilt, year int
		var name, manufacturer, typeJet, features, role, country, imageLink string
		var generation float64
		var createdAt, updatedAt *time.Time

		err = rows20th.Scan(&id, &name, &manufacturer, &typeJet, &year, &features, &numberBuilt, &role, &country, &imageLink, &generation, &createdAt, &updatedAt)
		if err != nil {
			return nil, nil, err
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
		jets20th = append(jets20th, jet)
	}

	for rows21st.Next() {
		var id, numberBuilt, year int
		var name, manufacturer, typeJet, features, role, country, imageLink string
		var generation float64
		var createdAt, updatedAt *time.Time

		err = rows21st.Scan(&id, &name, &manufacturer, &typeJet, &year, &features, &numberBuilt, &role, &country, &imageLink, &generation, &createdAt, &updatedAt)
		if err != nil {
			return nil, nil, err
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
		jets21st = append(jets21st, jet)
	}
	return jets20th, jets21st, nil
}

// GetDistinctCountry gets distinct countries from the database
func GetDistinctCountries() ([]string, error) {
	getDistinctCountriesQuery := `select distinct country from jets`
	rows, err := db.Query(getDistinctCountriesQuery)
	if err != nil {
		return nil, err
	}

	var countriesList []string
	for rows.Next() {
		var country string

		err = rows.Scan(&country)
		if err != nil {
			return nil, err
		}

		countriesList = append(countriesList, country)
	}

	return countriesList, nil
}

// GetJetsByCountry gets fighter jets according to a specific country
func GetJetsByCountry(country string) ([]models.FighterJet, error) {
	getJetsByCountryQuery := `select * from jets where country=$1`
	rows, err := db.Query(getJetsByCountryQuery, country)
	if err != nil {
		return nil, err
	}

	var jets []models.FighterJet

	for rows.Next() {
		var id, numberBuilt, year int
		var name, manufacturer, typeJet, features, role, country, imageLink string
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
