package domain

type IDepartment interface {
	GetTitle() string
	GetDetails() string
	GetAddress() string
	GetLocality() string
	GetPrice() string
	GetImage() string
	GetLinkToDetail() string
}
