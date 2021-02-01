package domain

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	ID           uint   `json:"id"`
	Address      string `json:"address"`
	Locality     string `json:"locality"`
	Title        string `json:"title"`
	Details      string `json:"details"`
	Price        string `json:"price"`
	Image        string `json:"image"`
	LinkToDetail string `json:"link-detail"`
}

func (d Department) GetTitle() string        { return d.Title }
func (d Department) GetDetails() string      { return d.Details }
func (d Department) GetAddress() string      { return d.Address }
func (d Department) GetLocality() string     { return d.Locality }
func (d Department) GetPrice() string        { return d.Price }
func (d Department) GetImage() string        { return d.Image }
func (d Department) GetLinkToDetail() string { return d.LinkToDetail }

type Entity struct {
	gorm.Model
	ID         uint        `json:"id"`
	Name       string      `json:"name"`
	Url        string      `json:"url"`
	Type       string      `json:"type"`
	EntityUrls []EntityUrl `json:"entityUrls"`
}

type EntityUrl struct {
	gorm.Model
	ID       uint   `json:"id"`
	EntityID uint   `json:"entityId"`
	Url      string `json:"url"`
}

func (e Entity) GetAssociatedUrls() []string {
	var urls []string
	for _, v := range e.EntityUrls {
		urls = append(urls, v.Url)
	}
	return urls
}
