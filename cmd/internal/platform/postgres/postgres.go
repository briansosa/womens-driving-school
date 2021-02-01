package postgres

import (
	"braiton/braiton-home/cmd/internal/domain"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbContext *gorm.DB

var structureToTable = map[string]string{
	"Entity":     "Entities",
	"EntityUrl":  "EntityUrls",
	"Department": "Department",
}

func ConfigDatabase() error {
	dsn := "user=postgres password=admin dbname=DepartmentSearch port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}

	db.AutoMigrate(&domain.Entity{}, &domain.EntityUrl{}, &domain.Department{})
	dbContext = db
	return nil
}

// Parameters example
// &models.Entity{
// 		Name: "Zonaprop",
// 		Type: "Buscador",
// 		Url:  "https://www.zonaprop.com.ar/",
// }
func Create(entity interface{}) error {
	if err := dbContext.Create(entity).Error; err != nil {
		return fmt.Errorf("No record could be created. Error: %s", err)
	}
	return nil
}

// Parameters example
// entity := models.Entity{
// 	ID: 3,
// }
// NOTE: pass the variable reference to the method
func Delete(entity interface{}) error {
	if err := dbContext.Delete(entity).Error; err != nil {
		return fmt.Errorf("No record could be deleted. Error: %s", err)
	}
	return nil
}

// Parameters example
// entity := models.Entity{
// 	ID: 3,
// }
// NOTE: pass the variable reference to the method
func GetFirst(entity interface{}) error {
	if err := dbContext.First(entity).Error; err != nil {
		return fmt.Errorf("No record could be found. Error: %s", err)
	}
	return nil
}

// Parameters example
// var entity []models.Entity
// NOTE: pass the variable reference to the method
func GetAll(entity interface{}) error {
	// if err := dbContext.Preload("EntityUrls").Find(entity).Error; err != nil {
	if err := dbContext.Find(entity).Error; err != nil {
		return fmt.Errorf("No records could be found. Error: %s", err)
	}
	return nil
}

// Parameters example
// entity := models.Entity{
// 	ID: 3,
// }
// entityData := models.Entity{
// 	Name: "Buscadorprop",
// }
// NOTE: pass the variable entity reference to the method
func Update(entity, entityData interface{}) error {
	if err := dbContext.Model(entity).Updates(entityData).Error; err != nil {
		return fmt.Errorf("No records could be updated. Error: %s", err)
	}
	return nil
}

func GetAllWithAssociatedTables(entity interface{}, associatedTables ...interface{}) error {
	db := dbContext
	for _, table := range associatedTables {
		key := strings.Split(reflect.TypeOf(table).String(), ".")[1]
		tableName := structureToTable[key]
		db = db.Preload(tableName)
	}
	if err := db.Find(entity).Error; err != nil {
		return fmt.Errorf("No records could be found. Error: %s", err)
	}
	return nil
}

func GetFirstWithAssociatedTables(entity interface{}, associatedTables ...interface{}) error {
	db := dbContext
	for _, table := range associatedTables {
		key := strings.Split(reflect.TypeOf(table).String(), ".")[1]
		tableName := structureToTable[key]
		db = db.Preload(tableName)
	}
	if err := db.First(entity).Error; err != nil {
		return fmt.Errorf("No record could be found. Error: %s", err)
	}
	return nil
}
