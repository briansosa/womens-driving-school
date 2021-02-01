package application

import (
	"braiton/braiton-home/cmd/internal/domain"
	"braiton/braiton-home/cmd/internal/platform/postgres"
	"braiton/braiton-home/cmd/internal/utils"
	"fmt"
)

func GetDepartments() ([]domain.Department, error) {
	var departments []domain.Department
	err := postgres.GetAll(&departments)
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func GetDepartment(id uint) (domain.Department, error) {
	department := domain.Department{
		ID: id,
	}
	err := postgres.GetFirst(&department)
	if err != nil {
		return domain.Department{}, err
	}
	return department, nil
}

func CreateDepartment(departmentDto domain.DepartmentDto) error {
	var departmentModel domain.Department
	err := utils.MapperStructs(departmentDto, &departmentModel)
	if err != nil {
		return fmt.Errorf("No se pudo crear el departamento. Error: %s", err)
	}
	err = postgres.Create(&departmentModel)
	if err != nil {
		return fmt.Errorf("No se pudo crear el departamento. Error: %s", err)
	}
	return nil
}

func UpdateDepartment(departmentDto domain.DepartmentDto, id uint) error {
	var departmentModel domain.Department
	err := utils.MapperStructs(departmentDto, &departmentModel)
	if err != nil {
		return fmt.Errorf("No se pudo actualizar el departamento. Error: %s", err)
	}
	departmentToFind := domain.Department{
		ID: id,
	}

	err = postgres.Update(&departmentToFind, departmentModel)
	if err != nil {
		return fmt.Errorf("No se pudo actualizar el departamento. Error: %s", err)
	}
	return nil
}

func DeleteDepartment(id uint) error {
	department := domain.Department{
		ID: id,
	}
	err := postgres.Delete(&department)
	if err != nil {
		return fmt.Errorf("No se pudo borrar el departamento. Error: %s", err)
	}
	return nil
}
