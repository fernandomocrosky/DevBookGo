package repositories

import (
	"database/sql"
	"fmt"

	"github.com/fernandomocrosky/DevBookGo/src/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(`
		INSERT INTO users 
			(name, nick, email, password)
		VALUES
			(?,?,?,?)
	`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

func (repo *UserRepository) GetAllUsers(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := repo.db.Query(`
		SELECT id, name, nick, email, created_at FROM users
		WHERE name LIKE ? or nick LIKE ?
	`, nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepository) GetUserById(id int64) (models.User, error) {

	rows, err := repo.db.Query(`
		SELECT id, name, nick, email, created_at
		FROM users WHERE id = ?
	`, id)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo *UserRepository) UpdateUser(id uint64, user models.User) error {
	statement, err := repo.db.Prepare(`
		UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?;
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nick, user.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) DeleteUser(id uint64) error {
	statement, err := repo.db.Prepare(`
		DELETE FROM users WHERE id = ?;
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) GetUserByEmail(email string) (models.User, error) {
	row, err := repo.db.Query(`
		select id, password from users where email = ?
	`, email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
