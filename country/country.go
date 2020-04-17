package country

import (
	"encoding/json"
	"net/http"
)

//
// This API resouce from rest countries please refer this for all documentation
// https://restcountries.eu/#api-endpoints-all
//

// Country containts full of country description
type Country struct {
	Name           string            `json:"name"`
	TopLevelDomain []string          `json:"topLevelDomain"`
	Alpha2Code     string            `json:"alpha2Code"`
	Alpha3Code     string            `json:"alpha3Code"`
	CallingCodes   []string          `json:"callingCodes"`
	Capital        string            `json:"capital"`
	AltSpellings   []string          `json:"altSpellings"`
	Region         string            `json:"region"`
	Subregion      string            `json:"subregion"`
	Population     int64             `json:"population"`
	Latlng         []float64         `json:"latlng"`
	Demonym        string            `json:"demonym"`
	Area           *float64          `json:"area"`
	Gini           *float64          `json:"gini"`
	Timezones      []string          `json:"timezones"`
	Borders        []string          `json:"borders"`
	NativeName     string            `json:"nativeName"`
	NumericCode    string            `json:"numericCode"`
	Currencies     []*Currency       `json:"currencies"`
	Languages      []*Language       `json:"languages"`
	Translations   map[string]string `json:"translations"`
	Flag           string            `json:"flag"`
	RegionalBlocs  []*RegionalBloc   `json:"regionalBlocs"`
	Cioc           string            `json:"cioc"`
}

// Currency containt currency code name and symbol
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// Language containt standard akronim code, name and native name
type Language struct {
	ISO6391    string `json:"iso639_1"`
	ISO6392    string `json:"iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

// RegionalBloc containt acronym, fullname, or others name
type RegionalBloc struct {
	Acronym       string   `json:"acronym"`
	Name          string   `json:"name"`
	OtherAcronyms []string `json:"otherAcronyms"`
	OtherNames    []string `json:"otherNames"`
}

// HTTPCountry struct for get http country
type HTTPCountry struct {
}

// GetAll return all country list information
func (httpCountry *HTTPCountry) GetAll() (countryList []*Country, err error) {
	resp, err := http.Get(URLRESTCOUNTRIES + URLAllCountry)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&countryList)
	return countryList, err
}
