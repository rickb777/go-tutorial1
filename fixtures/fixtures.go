package fixtures

import (
	"github.com/gofrs/uuid"
	"github.com/rickb777/date/v2"
	"github.com/rickb777/go-tutorial1/person"
	"github.com/rickb777/go-tutorial1/person/country"
	"math/rand/v2"
)

var addresses = []person.Address{
	{Address: "1 Old Down Rd", City: "Andover", Postcode: "SP10 3JP", Country: country.GB},
	{Address: "24 Ashleigh Rd", City: "Barnstaple", Postcode: "EX32 8JY", Country: country.GB},
	{Address: "46 Princess Anne Rd", City: "Boston", Postcode: "PE21 9AP", Country: country.GB},
	{Address: "Flat 1a, Tŷ Gwyn, 13 Stryd y Dŵr", City: "Bridgend", Postcode: "CF314AD", Country: country.GB},
	{Address: "1 Crosswell Cl, North Petherton", City: "Bridgewater", Postcode: "TA6 6SF", Country: country.GB},
	{Address: "18 Alderson Ct", City: "Bridlington", Postcode: "Bridlington", Country: country.GB},
	{Address: "62 Oxford St", City: "Burnham-on-Sea", Postcode: "TA81EW", Country: country.GB},
	{Address: "Bryn Ogwen, England Rd N", City: "Caernarfon", Postcode: "LL55 1HS", Country: country.GB},
	{Address: "10 Graham St, Longtown", City: "Carlisle", Postcode: "CA6 5NR", Country: country.GB},
	{Address: "34 Mount Camel", City: "Camelford", Postcode: "PL32 9UW", Country: country.GB},
	{Address: "6 St Mary's Place", City: "Chippenham", Postcode: "SN15 1EN", Country: country.GB},
	{Address: "20 Middlemarch Rd", City: "Coventry", Postcode: "CV6 3GF", Country: country.GB},
	{Address: "34 Kings Mill Park", City: "Driffield", Postcode: "YO25 6UZ", Country: country.GB},
	{Address: "1 McFarlane Ave, Kingholm Quay", City: "Dumfries", Postcode: "DG1 4GB", Country: country.GB},
	{Address: "14 Cottar St", City: "Glasgow", Postcode: "G20 0NL", Country: country.GB},
	{Address: "Silverhill, Whiting Bay", City: "Isle of Arran", Postcode: "KA27 8QL", Country: country.GB},
	{Address: "451 Braunstone Ln", City: "Leicester", Postcode: "LE3 3DD", Country: country.GB},
	{Address: "6 Queen Mary's Walk", City: "Llanelli", Postcode: "SA15 1PG", Country: country.GB},
	{Address: "2 Stockton Gardens", City: "London", Postcode: "N17 7HY", Country: country.GB},
	{Address: "24 Hillside Rd, Rothbury", City: "Morpeth", Postcode: "NE65 7YH", Country: country.GB},
	{Address: "29 St John's Rd", City: "Newbury", Postcode: "RG14 7PY", Country: country.GB},
	{Address: "29, Plymouth Rd", City: "Penarth", Postcode: "CF64 3DA", Country: country.GB},
	{Address: "10 Leys Ave", City: "Rothwell", Postcode: "NN14 6JF", Country: country.GB},
	{Address: "62 Pankhurst Cres", City: "Stevenage", Postcode: "SG2 0QF", Country: country.GB},
	{Address: "22 MacDonald Dr", City: "Stirling", Postcode: "FK7 9ER", Country: country.GB},
	{Address: "3 Norbury Ave, Marple", City: "Stockport", Postcode: "SK6 6NB", Country: country.GB},
	{Address: "15 Thrushel Close", City: "Swindon", Postcode: "SN25 3PP", Country: country.GB},
	{Address: "37 Laburnum St", City: "Taunton", Postcode: "TA1 1LB", Country: country.GB},
	{Address: "82 Woodland Rd", City: "Wolverhampton", Postcode: "WV3 8AW", Country: country.GB},
	{Address: "15 Reckondales Fld, Stamford Bridge", City: "York", Postcode: "YO41 1FL", Country: country.GB},
}

// RandomAddress picks a random address from a short list.
func RandomAddress() person.Address {
	idx := rand.IntN(len(addresses))
	return addresses[idx]
}

//-------------------------------------------------------------------------------------------------

var (
	firstNames = map[person.Gender][]string{
		person.Unspecified: {
			"Andy", "Chris",
		},

		person.Other: {
			"Andy", "Chris",
		},

		person.Female: {
			"Ann", "Alice", "Belinda", "Charlotte", "Christine", "Chloë",
			"Daisy", "Darcie", "Delilah", "Emma", "Eloise",
			"Félicité", "Georgie", "Hannah", "Ivy", "Jane", "Katherine",
			"Louise", "Miriam", "Niamh", "Rachel", "Rebecca", "Roberta", "Robin",
			"Siân", "Sørina", "Susan", "Zoë",
		},

		person.Male: {
			"Aaron", "Adam", "Albert", "Andrew", "Charles", "David", "Declan", "Derek", "Donald",
			"Edward", "François", "George", "Henry", "Ian", "John", "José", "Levi", "Mark", "Martin",
			"Neil", "Noël", "Owain", "Øyvind", "Rowan", "Ron", "Robert", "Seán", "Thomas", "Timothy",
		},
	}

	lastNames = []string{
		"Anderson", "Birkett", "Chase", "Davidson", "Engels",
		"Frank", "Garth", "Hubert", "Illingworth", "Janson",
		"Kirk", "Lucas", "Monroe", "Norbert", "Oppenheimer",
		"Parish", "Queen", "Rogers", "Smith", "Tucker",
		"Umbridge", "Vickers", "Williams", "Yarrow", "Zeland",
	}

	titles = map[person.Gender][]string{
		person.Female: {"Mrs", "Ms", "Miss", "Dr", "Lady"},
		person.Male:   {"Mr", "Mr", "Mr", "Dr", "Sir"},
	}
)

const daysPerYearGregorianE4 = 3652425 // 365.2425 days per year in the Gregorian calendar

func randomPerson(gender person.Gender, age int) person.Person {
	firstIdx := rand.IntN(len(firstNames[gender]))
	lastIdx := rand.IntN(len(lastNames))

	name := firstNames[gender][firstIdx] + " " + lastNames[lastIdx]

	ageInDays := (age*daysPerYearGregorianE4 + (rand.IntN(365) * 10000)) / 10000

	return person.Person{
		UUID:      uuid.Must(uuid.NewV4()),
		Name:      name,
		Birthday:  date.Today() - date.Date(ageInDays),
		Gender:    gender,
		Addresses: nil,
	}
}

func RandomGender() person.Gender {
	return person.Gender(rand.IntN(len(person.AllGenders)))
}

func RandomMaleOrFemale() person.Gender {
	return person.Gender(rand.IntN(2)) + person.Male
}

// RandomAdult invents a person who is at least 18 years old.
func RandomAdult(gender person.Gender) person.Person {
	titleIdx := rand.IntN(len(titles[gender]))
	adult := randomPerson(gender, rand.IntN(92)+18)
	if titleIdx < len(titles[gender]) {
		adult.Title = titles[gender][titleIdx]
	}
	adult.Addresses = append(adult.Addresses, RandomAddress())
	movedIn := date.Today().AddDate(0, 0, 7) - date.Date(rand.IntN(3652))
	adult.MovedIn = append(adult.MovedIn, movedIn)
	return adult
}

// RandomChild invents a person who is less than 18 years old.
func RandomChild(gender person.Gender) person.Person {
	return randomPerson(gender, rand.IntN(18))
}
