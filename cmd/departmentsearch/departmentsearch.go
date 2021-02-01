package departmentsearch

import (
	"braiton/braiton-home/cmd/departmentsearch/routes"
	"braiton/braiton-home/cmd/departmentsearch/scraping"
	"braiton/braiton-home/cmd/internal/application"
	"braiton/braiton-home/cmd/internal/comparator"
	"braiton/braiton-home/cmd/internal/platform/postgres"
	"fmt"
)

func Run() {
	// config db
	err := postgres.ConfigDatabase()
	if err != nil {
		panic(err)
	}

	// get urls to be consulted
	entities, err := application.GetEntitiesWithRelationship()
	if err != nil {
		panic(err)
	}
	// make scrapping
	deptos := scraping.ScrapingDepartments(entities)
	// fmt.Println(deptos)
	filteredDeptos := comparator.FilterDepartments(deptos)
	fmt.Println(filteredDeptos)
	// compareIfNotExisting
	// manage files
	// put on db
	// config routes
	routes.ConfigRoutes()

}
