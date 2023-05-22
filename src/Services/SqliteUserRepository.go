package Services

import (
	"4_zad/Models"
	"gorm.io/gorm"
	"time"
)

type SqliteUserRepository struct {
	db *gorm.DB
}

func NewSqliteUserRepository(db *gorm.DB) IRepository[Models.User] {
	migrationError := db.AutoMigrate(&Models.User{})
	if migrationError != nil {
		panic(migrationError)
	}

	return &SqliteUserRepository{db}
}

func (repo *SqliteUserRepository) GetAll() []Models.User {
	var users []Models.User
	repo.db.Find(&users)
	return users
}

func (repo *SqliteUserRepository) GetById(id uint) Models.User {
	var user Models.User
	repo.db.First(&user, id)
	return user
}

func (repo *SqliteUserRepository) GetWithFilter() []Models.User {
	panic("Not implemented!")
}

func (repo *SqliteUserRepository) Add(user Models.User) Models.User {
	var userToCreate Models.User
	userToCreate = user
	repo.db.Create(&userToCreate)
	return user
}

func (repo *SqliteUserRepository) Update(user Models.User) Models.User {
	repo.db.Save(&user)
	return user
}

func (repo *SqliteUserRepository) Delete(id uint) Models.User {
	repo.db.Delete(&Models.User{}, id)
	return Models.User{}
}

func (repo *SqliteUserRepository) Patch(id uint, userPatch any) Models.User {
	var userToUpdate Models.User
	repo.db.First(&userToUpdate, id)
	if userPatch.(Models.PatchProductBinding).Name != nil {
		userToUpdate.Email = *userPatch.(Models.PatchUserBinding).Email
	}
	if userPatch.(Models.PatchProductBinding).Price != nil {
		userToUpdate.Password = *userPatch.(Models.PatchUserBinding).Password
	}

	userToUpdate.UpdatedAt = time.Now()
	repo.db.Save(&userToUpdate)
	return userToUpdate
}
