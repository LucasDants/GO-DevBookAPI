package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

// Cria um repositorio de usuarios
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

// Insere usuario no DB
func (repository users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

// Busca todos os usuarios filtrados
func (repository users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	rows, err := repository.db.Query("select id, name, nick, email, createdAt from users where name LIKE ? or nick Like ?", nameOrNick, nameOrNick)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository users) SearchByID(ID uint64) (models.User, error) {
	row, err := repository.db.Query("select id, name, nick, email, createdAt from users where id = ?", ID)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (repository users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository users) SearchByEmail(email string) (models.User, error) {
	row, err := repository.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (repository users) Follow(followedID, followerID uint64) error {
	statement, err := repository.db.Prepare("insert ignore into followers(user_id, follower_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(followedID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository users) Unfollow(followedID, followerID uint64) error {
	statement, err := repository.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(followedID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository users) GetFollowers(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdAt
		from users u inner join followers s on u.id = s.follower_id where s.user_id = ?
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository users) GetFollowing(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdAt
		from users u inner join followers s on u.id = s.user_id where s.follower_id = ?
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository users) GetPassword(userID uint64) (string, error) {
	row, err := repository.db.Query("select password from users where id = ?", userID)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", err
		}
	}
	return user.Password, nil
}

func (repository users) ChangePassword(userID uint64, password string) error {
	statement, err := repository.db.Prepare("update users set password = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}
