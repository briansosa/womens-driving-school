package comparator

import (
	"braiton/braiton-home/cmd/internal/scraping"
)

func FilterDepartments(departments [][]scraping.Department) []scraping.Department {
	var departmentsFiltered []scraping.Department
	for _, departmentsInUrl := range departments {
		for _, department := range departmentsInUrl {
			if !ExistDeparment(department, departmentsFiltered) {
				departmentsFiltered = append(departmentsFiltered, department)
			}
		}
	}
	return departmentsFiltered
}

func ExistDeparment(department scraping.Department, departmentsFiltered []scraping.Department) bool {
	for _, value := range departmentsFiltered {
		if compareTitle(department.Title, value.Title) &&
			compareDescription(department.Details, value.Details) {
			return true
		}
	}
	return false
}

func compareTitle(depTitle string, depFilteredTitle string) bool {
	return depTitle == depFilteredTitle
}

func compareDescription(depDescription string, depFilteredDescription string) bool {
	return depDescription == depFilteredDescription
}
