package database

import "log"

type Rate struct {
	Src  string `binding:"required"`
	Like bool
}

const (
	AddRatingToFilm         = "INSERT INTO FilmRating(UserID, FilmID, Rating) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE Rating = ?"
	AddRatingToPlaylist     = "INSERT INTO PlaylistRating(UserID, PlaylistID, Rating) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE Rating = ?"
	GetUserRatingOfFilm     = "SELECT Rating FROM FilmRating WHERE UserID = ? AND FilmID = ?"
	GetUserRatingOfPlaylist = "SELECT Rating FROM PlaylistRating WHERE UserID = ? AND PlaylistID = ?"
)

func AddRating(query string, userID int, likeStatus bool, destID int) error {
	_, err := db.Exec(query,
		userID, destID, likeStatus, likeStatus)
	if err != nil {
		log.Println("Error adding rating: " + err.Error())
		return err
	}
	return err
}

func GetRatingByUser(query string, userID, srcID int) (int, error) {
	res := 0
	err := db.QueryRow(
		query,
		userID, srcID).
		Scan(&res)
	if err != nil {
		log.Println("Error fetching rating: " + err.Error())
		return 0, err
	}
	if res == 0 {
		return -1, nil
	}
	return res, nil
}

func getRating(query string, id int, amount *int) error {
	err := db.QueryRow(query, id).Scan(amount)
	if err != nil {
		log.Println("Error getting rating: " + err.Error())
		return err
	}
	return nil
}
