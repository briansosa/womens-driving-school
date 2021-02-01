package domain

type EntityDto struct {
	Id         uint           `json:"id"`
	Name       string         `json:"name"`
	Url        string         `json:"url"`
	Type       string         `json:"type"`
	EntityUrls []EntityUrlDto `json:"entityUrls"`
}

type EntityUrlDto struct {
	Id       uint   `json:"id"`
	EntityID uint   `json:"entityId"`
	Url      string `json:"url"`
}

type DepartmentDto struct {
	Id           uint   `json:"id"`
	Address      string `json:"address"`
	Locality     string `json:"locality"`
	Title        string `json:"title"`
	Details      string `json:"details"`
	Price        string `json:"price"`
	Image        string `json:"image"`
	LinkToDetail string `json:"link-detail"`
}
