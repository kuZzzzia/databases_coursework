package database

import (
	"database/sql"
	"log"
)

type Message struct {
	ID       int
	Review   string `binding:"required,min=5,max=63"`
	UserName sql.NullString
	FilmID   int
}

func FetchDiscussionsForFilm(id int) ([]*Message, error) {
	var discussion []*Message

	results, err := db.Query(
		"SELECT DiscussionID, Review, Username FROM Film_Discussion_With_Users WHERE FilmID = ? ORDER BY DiscussionID DESC",
		id)
	if err != nil {
		log.Println("Error fetching roles")
		return nil, err
	}

	defer results.Close()
	for results.Next() {
		discussionItem := new(Message)

		err = results.Scan(&discussionItem.ID, &discussionItem.Review, &discussionItem.UserName)
		if err != nil {
			log.Println("Error fetching roles")
			return nil, err
		}

		discussion = append(discussion, discussionItem)
	}
	return discussion, nil
}

func AddMessage(user *User, message *Message, filmID int) error {
	message.UserName = sql.NullString{
		String: user.Username,
		Valid:  true,
	}
	_, err := db.Query("INSERT INTO Discussion(`Date`, Review, UserID, FilmID) VALUES (NOW(), ?, ?, ?)",
		message.Review, user.ID, filmID)
	if err != nil {
		return err
	}
	return err
}
