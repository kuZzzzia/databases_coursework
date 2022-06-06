package database

import (
	"database/sql"
	"log"
)

type DiscussionItem struct {
	ID       int
	Date     string
	Review   sql.NullString
	UserID   sql.NullInt16
	UserName sql.NullString
}

func FetchDiscussionsForFilm(id int) ([]*DiscussionItem, error) {
	var discussion []*DiscussionItem

	results, err := db.Query(
		"SELECT DiscussionID, `Date`, Review, UserID, Username FROM Film_Discussion_With_Users WHERE FilmID = ? ORDER BY d.Date",
		id)
	if err != nil {
		log.Println("Error fetching roles")
		return nil, err
	}

	defer results.Close()
	for results.Next() {
		discussionItem := new(DiscussionItem)

		err = results.Scan(&discussionItem.ID, &discussionItem.Date, &discussionItem.Review, &discussionItem.UserID, &discussionItem.UserName)
		if err != nil {
			log.Println("Error fetching roles")
			return nil, err
		}

		discussion = append(discussion, discussionItem)
	}
	return discussion, nil
}
