package database

import (
	"database/sql"
	"log"
)

type Comment struct {
	ID       int
	Review   string `binding:"required,min=5,max=63"`
	UserName sql.NullString
	FilmID   int
}

func FetchCommentsForFilm(id int) ([]*Comment, error) {
	var discussion []*Comment

	results, err := db.Query(
		"SELECT CommentID, Review, Username FROM Film_Comments_With_Users WHERE FilmID = ? ORDER BY CommentID DESC",
		id)
	if err != nil {
		log.Println("Error fetching comments: " + err.Error())
		return nil, err
	}

	defer results.Close()
	for results.Next() {
		comment := new(Comment)

		err = results.Scan(&comment.ID, &comment.Review, &comment.UserName)
		if err != nil {
			log.Println("Error fetching comments: " + err.Error())
			return nil, err
		}

		discussion = append(discussion, comment)
	}
	return discussion, nil
}

func AddComment(user *User, comment *Comment, filmID int) error {
	comment.UserName = sql.NullString{
		String: user.Username,
		Valid:  true,
	}
	_, err := db.Exec("INSERT INTO `Comment`(`Date`, Review, UserID, FilmID) VALUES (NOW(), ?, ?, ?)",
		comment.Review, user.ID, filmID)
	if err != nil {
		log.Println("Error adding comment to film: " + err.Error())
		return err
	}
	return err
}
