package user

import (
	. "2019_2_Shtoby_shto/src/custom_type"
	"2019_2_Shtoby_shto/src/database"
)

type UserHandler interface {
	PutUser(user User) error
	GetUserById(id StringUUID) (User, error)
	GetUserByLogin(login string) (User, error)
}

type service struct {
	db *database.DataManager
}

func CreateInstance(db *database.DataManager) UserHandler {
	return &service{
		db: db,
	}
}

func (s *service) PutUser(user User) error {
	//err := s.db.ExecuteQuery("insert into users(id, login, password) values($1, $2, $3)", user.ID.String(), user.Login, user.Password)
	return s.db.CreateRecord(&user)
}

func (s *service) GetUserById(id StringUUID) (User, error) {
	user := User{}
	user.ID = id
	err := s.db.FindDictById(&user)
	return user, err
}

func (s *service) GetUserByLogin(login string) (User, error) {
	user := User{}
	err := s.db.FindDictByLogin(&user, "login", login)
	return user, err
}