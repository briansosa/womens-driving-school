package scraping

import (
	"braiton/braiton-home/cmd/internal/domain"
	"braiton/braiton-home/cmd/internal/scraping"
	"braiton/braiton-home/cmd/internal/scraping/argenprop"
	"braiton/braiton-home/cmd/internal/scraping/buscadorprop"

	"fmt"
)

// func MakeScrapingUrl(url string) {
// 	c := colly.NewCollector()
// 	asd := []string{url, "asd"}
// 	_ = buscadorprop.MakeScraping(c, asd)
// 	// fmt.Println(departments)
// }

func ScrapingDepartments(entities []domain.Entity) [][]scraping.Department {
	var allDepartmentsConsulted [][]scraping.Department
	for _, entity := range entities {
		departments, err := routerSwitch(entity)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			allDepartmentsConsulted = append(allDepartmentsConsulted, departments)
		}
	}
	return allDepartmentsConsulted
}

func routerSwitch(entity domain.Entity) ([]scraping.Department, error) {
	switch entity.Name {
	case "Buscadorprop":
		return buscadorprop.MakeScraping(entity.GetAssociatedUrls())
	// case "Zonaprop":
	// 	return zonaprop.MakeScraping(entity.GetAssociatedUrls())
	case "Argenprop":
		return argenprop.MakeScraping(entity.GetAssociatedUrls())
	default:
		return nil, fmt.Errorf("No se encontró scraping para la entidad: %s", entity.Name)
	}
}
