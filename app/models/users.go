package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos     []Todo
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.PassWord), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select * from users where id = ?`
	// cmd := `select id, uuid, name, email, password, created_at
	// from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select * from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	createCmd := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(createCmd, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	getSessionCmd := `select * from sessions where user_id = ? and email = ?`

	err = Db.QueryRow(getSessionCmd, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserId,
		&session.CreatedAt,
	)
	if err != nil {
		log.Println(err)
	}
	return session, err
}

func (session *Session) CheckSession() (valid bool, err error) {
	cmd := `select * from sessions where uuid = ?`

	err = Db.QueryRow(cmd, session.UUID).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserId,
		&session.CreatedAt,
	)
	if err != nil {
		valid = false
		return
	}
	if session.ID != 0 {
		valid = true
	}
	return valid, err
}

func (session *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, session.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (session *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select * from users where id = ?`
	err = Db.QueryRow(cmd, session.UserId).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	if err != nil {
		log.Println(err)
	}
	return user, err
}
