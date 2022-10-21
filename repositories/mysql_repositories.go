package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/titusdishon/go-docker-mysql/config"
	"github.com/titusdishon/go-docker-mysql/entity"
)

var db *sql.DB

func init() {
	config.Connect()
	db = config.GetDb()
}

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	Update(user *entity.User, id int64) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	FindById(id int64) (*entity.User, error)
	Delete(id int64) (int64, error)
}

type repo struct{}

func NewMysqlRepository() UserRepository {
	return &repo{}
}

func (*repo) Save(user *entity.User) (*entity.User, error) {
	stmt, err := db.Prepare("INSERT INTO users (name, email, summary) VALUES(?,?,?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Name, user.Email, user.Summary)
	if err != nil {
		return nil, err
	}
	fmt.Println("result", res)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (*repo) Update(user *entity.User, id int64) (*entity.User, error) {
	stmt, err := db.Prepare(`UPDATE users SET name = ?, email = ?, summary = ? WHERE id = ?;`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email, user.Summary, id)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func (*repo) FindAll() ([]*entity.User, error) {
	result, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	var users []*entity.User

	for result.Next() {
		var u entity.User
		err = result.Scan(&u.ID, &u.Name, &u.Email, &u.Summary)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}
	return users, nil
}

func (*repo) FindById(id int64) (*entity.User, error) {
	var user entity.User
	err := db.QueryRow(`SELECT id, email, name, summary FROM users WHERE id=?;`, id).Scan(&user.ID,
		&user.Email,
		&user.Name,
		&user.Summary)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (*repo) Delete(id int64) (int64, error) {
	result, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		panic(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}
