package database

import (
	"log"
)

const (
	queryGenres    = "SELECT GenreName FROM Genre"
	queryCountries = "SELECT CountryName FROM Country"
)

func queryCategory(query string) ([]*string, error) {
	var items []*string
	results, err := db.Query(query)
	if err != nil {
		log.Println("Error fetching category: " + err.Error())
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		item := new(string)

		err = results.Scan(&item)
		if err != nil {
			log.Println("Error fetching category: " + err.Error())
			return nil, err
		}

		items = append(items, item)
	}
	return items, nil
}

func FetchCategories() ([]*string, []*string, error) {
	countries, err := queryCategory(queryCountries)
	if err != nil {
		return nil, nil, err
	}
	genres, err := queryCategory(queryGenres)
	if err != nil {
		return nil, nil, err
	}

	return genres, countries, nil
}
