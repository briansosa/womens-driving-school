package application

import (
	"braiton/braiton-home/cmd/internal/domain"
	"braiton/braiton-home/cmd/internal/platform/postgres"
	"braiton/braiton-home/cmd/internal/utils"
	"fmt"
)

func GetEntities() ([]domain.Entity, error) {
	var entities []domain.Entity
	err := postgres.GetAll(&entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func GetEntitiesWithRelationship() ([]domain.Entity, error) {
	var entities []domain.Entity
	err := postgres.GetAllWithAssociatedTables(&entities, domain.EntityUrl{})
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func GetEntityWithRelationship(id uint) (domain.Entity, error) {
	entity := domain.Entity{
		ID: id,
	}
	err := postgres.GetFirstWithAssociatedTables(&entity, domain.EntityUrl{})
	if err != nil {
		return domain.Entity{}, err
	}
	return entity, nil
}

func GetEntity(id uint) (domain.Entity, error) {
	entity := domain.Entity{
		ID: id,
	}
	err := postgres.GetFirst(&entity)
	if err != nil {
		return domain.Entity{}, err
	}
	return entity, nil
}

func CreateEntity(entityDto domain.EntityDto) error {
	var entityModel domain.Entity
	err := utils.MapperStructs(entityDto, &entityModel)
	if err != nil {
		return fmt.Errorf("No se pudo crear la entidad. Error: %s", err)
	}
	err = postgres.Create(&entityModel)
	if err != nil {
		return fmt.Errorf("No se pudo crear la entidad. Error: %s", err)
	}
	return nil
}

func UpdateEntity(entityDto domain.EntityDto, id uint) error {
	var entityModel domain.Entity
	err := utils.MapperStructs(entityDto, &entityModel)
	if err != nil {
		return fmt.Errorf("No se pudo actualizar la entidad. Error: %s", err)
	}
	entityToFind := domain.Entity{
		ID: id,
	}

	err = postgres.Update(&entityToFind, entityModel)
	if err != nil {
		return fmt.Errorf("No se pudo actualizar la entidad. Error: %s", err)
	}
	return nil
}

func DeleteEntity(id uint) error {
	entity := domain.Entity{
		ID: id,
	}
	err := postgres.Delete(&entity)
	if err != nil {
		return fmt.Errorf("No se pudo borrar la entidad. Error: %s", err)
	}
	return nil
}
