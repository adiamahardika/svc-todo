package database

import (
	"svc-todo/entity"

	"gorm.io/gorm"
)

func MigrationPostgres(db *gorm.DB) {

	db.AutoMigrate(&entity.List{}, &entity.SubList{})

	db.Migrator().CreateIndex(&entity.List{}, "id")
	db.Migrator().CreateIndex(&entity.SubList{}, "id")
}
