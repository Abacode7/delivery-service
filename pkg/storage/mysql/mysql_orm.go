package mysql

import (
	"fmt"
	"github.com/Abacode7/delivery-service/pkg/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type storage struct {
	db *gorm.DB
}

func NewStorage(username, password, dbHost, dbName string)  storage{
	dbUrl := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName)

	// Database initialisation happens here
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil{
		fmt.Println(err)
	}

	// All model migrations happen here
	db.AutoMigrate(&domain.User{})

	return storage{db}
}

func (s *storage) FindByID(id int) *domain.User  {
	return nil
}

func (s * storage) FindAll() []domain.User{
	return nil
}

func (s *storage) Save(user domain.User) (*domain.User, error){
	return nil, nil
}