package repository

// Repository is an interface used for interacting with the backend database
type Repository interface {
	AllSites() ([]Site, error)
	FindSite(id string) (Site, error)
	CreateSite(url string) (Site, error)
	UpdateSite(id, url string) (Site, error)
	DeleteSite(id string) error
}
