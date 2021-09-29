package repositories

import (
	"api/src/models"
	"database/sql"
)

type publications struct {
	db *sql.DB
}

func NewPublicationsRepository(db *sql.DB) *publications {
	return &publications{db}
}

func (repository publications) Create(publication models.Publications) (uint64, error) {
	statement, err := repository.db.Prepare("insert into publications (title, content, author_id) values (?, ?, ?)")
	if err != nil {
		return 0, nil
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, nil
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastInsertID), nil
}

func (repository publications) Search(userID uint64) ([]models.Publications, error) {
	rows, err := repository.db.Query(`
		select distinct p.*, u.nick from publications p
	 	inner join users u on u.id = p.author_id
	 	inner join followers f on p.author_id = f.user_id 
	 	where u.id = ? or f.follower_id = ?
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publications

	for rows.Next() {
		var publication models.Publications

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}
		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository publications) Update(ID uint64, publication models.Publications) error {
	statement, err := repository.db.Prepare("update publications set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, ID); err != nil {
		return err
	}

	return nil
}

func (repository publications) SearchByID(publicationID uint64) (models.Publications, error) {
	row, err := repository.db.Query(`
		select p.*, u.nick from 
		publications p inner join users u
		on u.id = p.author_id where p.id = ?
	`, publicationID)
	if err != nil {
		return models.Publications{}, err
	}
	defer row.Close()

	var publication models.Publications

	if row.Next() {
		if err = row.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return models.Publications{}, err
		}
	}
	return publication, nil
}

func (repository publications) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("delete from publications where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository publications) SearchByUser(userID uint64) ([]models.Publications, error) {
	rows, err := repository.db.Query(`
		select p.*, u.nick from publications p 
		join users u on u.id = p.author_id 
		where p.author_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publications

	for rows.Next() {
		var publication models.Publications

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}
		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository publications) Like(publicationID uint64) error {
	statement, err := repository.db.Prepare("update publications set likes = likes + 1 where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

func (repository publications) Unlike(publicationID uint64) error {
	statement, err := repository.db.Prepare(`
		update publications set likes =
		CASE WHEN likes > 0 THEN likes - 1
		ELSE likes END
		where id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}
