package country

// Repository is common function country
type Repository interface {
	GetAll() (countryList []*Country, err error)
}
