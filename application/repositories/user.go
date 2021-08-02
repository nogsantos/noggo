package repositories

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/nogsantos/noggo/domain"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDB struct {
	Db *gorm.DB
}

func (repo UserRepositoryDB) Insert(user *domain.User) (*domain.User, error) {


	prepare_error := user.Prepare()

	if prepare_error != nil {
		log.Fatalf("Error validate data %v", prepare_error)
		return user, prepare_error
	}

	create_error := repo.Db.Create(user).Error

	if create_error != nil {
		log.Fatalf("Error create %v", create_error)
		return user, create_error
	}

	return user, nil
}