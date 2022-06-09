package database

import "log"

type Rate struct {
	Src  string `binding:"required"`
	Like bool
}

const (
	AddRatingToFilm         = "INSERT INTO View(UserID, FilmID, FilmScore) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE FilmScore = ?"
	AddRatingToPlaylist     = "INSERT INTO PlaylistScore(UserID, PlaylistID, Score) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE Score = ?"
	GetUserRatingOfFilm     = "SELECT FilmScore FROM View WHERE UserID = ? AND FilmID = ?"
	GetUserRatingOfPlaylist = "SELECT Score FROM PlaylistScore WHERE UserID = ? AND PlaylistID = ?"
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
