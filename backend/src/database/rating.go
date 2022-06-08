package database

import "log"

type Rate struct {
	Src  string `binding:"required"`
	Like bool   `binding:"required"`
}

const (
	AddRatingToFilm         = "INSERT INTO View(UserID, FilmID, FilmScore) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE FilmScore = ?"
	AddRatingToPlaylist     = "INSERT INTO PlaylistScore(UserID, PlaylistID, Score) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE Score = ?"
	GetUserRatingOfFilm     = "SELECT FilmScore FROM View WHERE UserID = ? AND FilmID = ?"
	GetUserRatingOfPlaylist = "SELECT Score FROM PlaylistScore WHERE UserID = ? AND PlaylistID = ?"
)

func AddRate(query string, userID int, likeStatus bool, destID int) error {
	_, err := db.Query(query,
		userID, destID, likeStatus, likeStatus)
	if err != nil {
		return err
	}
	return err
}

func GetUserRate(query string, userID, srcID int) (int, error) {
	res := 0
	err := db.QueryRow(
		query,
		userID, srcID).
		Scan(&res)
	if err != nil {
		log.Println("Error fetching person" + err.Error())
		return 0, err
	}
	if res == 0 {
		return -1, nil
	}
	return res, nil
}
