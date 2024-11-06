package person

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/rickb777/date/v2"
	"github.com/rickb777/go-tutorial1/person/country"
	"strings"
)

//go:generate enumeration -i person.go -type Gender -ic -marshaljson identifier
type Gender int

const (
	Unspecified Gender = iota
	Male
	Female
	Other
)

// Address holds
type Address struct {
	UUID     uuid.UUID    `json:"id"`
	Address  string       `json:"address,omitempty"`
	City     string       `json:"city,omitempty"`
	Postcode string       `json:"postcode,omitempty"`
	Country  country.Code `json:"country,omitempty"` // ISO-3166 2-letter code
}

func (a Address) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", a.Address, a.City, a.Postcode, strings.ToUpper(string(a.Country)))
}

type Person struct {
	UUID      uuid.UUID   `json:"id"`
	Title     string      `json:"title,omitempty"`
	Name      string      `json:"name"`
	Birthday  date.Date   `json:"birthday,omitempty"`
	Gender    Gender      `json:"gender,omitempty"`
	Addresses []Address   `json:"addresses,omitempty"`
	MovedIn   []date.Date `json:"movedIn,omitempty"`
}

const daysPerYearGregorianE4 = 3652425 // 365.2425 days per year in the Gregorian calendar

func (p Person) Age() int {
	ageInDays := (today() - p.Birthday) * 10000
	return int(ageInDays / daysPerYearGregorianE4)
}

func (p Person) IsAdult() bool {
	return p.Age() >= 18
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s (%s age %d)", p.Title, p.Name, p.Birthday, p.Age())
}

// today provides a source of today's date, pluggable for testing.
var today = func() date.Date { return date.Today() }
