package zonaprop

import (
	"braiton/braiton-home/cmd/internal/scraping"
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func MakeScraping(url []string) ([]scraping.Department, error) {
	var departments []scraping.Department
	if len(url) < 0 {
		return departments, fmt.Errorf("No hay urls para hacer scraping")
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	for _, pageUrl := range url {
		var source string
		task := chromedp.Tasks{
			chromedp.Navigate(pageUrl),
			chromedp.WaitVisible(`#react-posting-cards`),
			chromedp.OuterHTML("html", &source),
			// chromedp.OuterHTML("#react-posting-cards", &source, chromedp.NodeVisible, chromedp.ByID),
		}

		err := chromedp.Run(ctx, task)
		if err != nil {
			return departments, fmt.Errorf("error inicializar chromedp")
		}

		pageReader := strings.NewReader(source)
		page, err := goquery.NewDocumentFromReader(pageReader)
		if err != nil {
			return departments, fmt.Errorf("error al crear goquery")
		}

		page.Find(".posting-card").Each(func(i int, s *goquery.Selection) {
			department := scraping.Department{
				Address:      GetAddress(s),
				Locality:     GetLocality(s),
				Title:        GetTitle(s),
				Details:      GetDetails(s),
				Price:        GetPrice(s),
				LinkToDetail: pageUrl + GetLinkToDetail(s),
				Image:        GetImageName(s),
			}
			departments = append(departments, department)
		})
	}
	return departments, nil
}

func GetAddress(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find(".posting-location").Text())
}

func GetLocality(s *goquery.Selection) string {
	return ""
}

func GetTitle(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find(".posting-title").Text())
}

func GetDetails(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find(".posting-description").Text())
}

func GetPrice(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find(".first-price").Text())
}

func GetLinkToDetail(s *goquery.Selection) string {
	link, _ := s.Find(".first-price > a").Attr("href")
	return strings.TrimSpace(link)
}

func GetImageLink(s *goquery.Selection) string {
	link, _ := s.Find(".slide-content.is-selected > img").Attr("src")
	if link == "" {
		return ""
	}
	return link
}

func GetImageName(s *goquery.Selection) string {
	linkToImage := GetImageLink(s)
	if linkToImage == "" {
		return ""
	}
	splitedLink := strings.Split(linkToImage, "/")
	imageName := splitedLink[len(splitedLink)-1]
	return imageName
}

// SCRAPING FOR STATIC HTML

// func MakeScraping(c *colly.Collector, url []string) ([]scraping.Department, error) {
// 	var departments []scraping.Department
// 	if len(url) < 0 {
// 		return departments, fmt.Errorf("No hay urls para hacer scraping")
// 	}

// 	c.OnHTML(".publicacion-item ", func(e *colly.HTMLElement) {
// 		department := scraping.Department{
// 			Address:      GetAddress(e),
// 			Locality:     GetLocality(e),
// 			Title:        GetTitle(e),
// 			Details:      GetDetails(e),
// 			Price:        GetPrice(e),
// 			LinkToDetail: GetLinkToDetail(e),
// 			Image:        GetImageName(e),
// 		}
// 		departments = append(departments, department)

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})

// 	for _, urlValue := range url {
// 		c.Visit(urlValue)
// 	}
// 	return departments, nil
// }

// func GetAddress(e *colly.HTMLElement) string {
// 	return e.ChildText(".card__location")
// }

// func GetAddress(e *colly.HTMLElement) string {
// 	return e.ChildText(".calle")
// }

// func GetLocality(e *colly.HTMLElement) string {
// 	return e.ChildText(".localidad")
// }

// func GetTitle(e *colly.HTMLElement) string {
// 	return e.ChildText(".content h2")
// }

// func GetDetails(e *colly.HTMLElement) string {
// 	return e.ChildText(".descripcion p")
// }

// func GetPrice(e *colly.HTMLElement) string {
// 	return e.ChildText(".precio")
// }

// func GetLinkToDetail(e *colly.HTMLElement) string {
// 	link := e.ChildAttr(".col-content > a", "href")
// 	absoluteLink := e.Request.AbsoluteURL(link)
// 	return absoluteLink
// }

// func GetImageLink(e *colly.HTMLElement) string {
// 	return e.ChildAttr(".col-image img", "src")
// }

// func GetImageName(e *colly.HTMLElement) string {
// 	linkToImage := GetImageLink(e)
// 	splitedLink := strings.Split(linkToImage, "/")
// 	imageName := splitedLink[len(splitedLink)-1]
// 	return imageName
// }
