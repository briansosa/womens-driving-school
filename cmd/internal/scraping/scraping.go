package scraping

type Department struct {
	Address      string
	Locality     string
	Title        string
	Details      string
	Price        string
	Image        string
	LinkToDetail string
}

func (d Department) GetTitle() string        { return d.Title }
func (d Department) GetDetails() string      { return d.Details }
func (d Department) GetAddress() string      { return d.Address }
func (d Department) GetLocality() string     { return d.Locality }
func (d Department) GetPrice() string        { return d.Price }
func (d Department) GetImage() string        { return d.Image }
func (d Department) GetLinkToDetail() string { return d.LinkToDetail }
