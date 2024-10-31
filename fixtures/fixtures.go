package fixtures

import (
	"github.com/gofrs/uuid"
	"github.com/rickb777/date/v2"
	"github.com/rickb777/go-tutorial1/person"
	"math/rand/v2"
)

var (
	firstNames = map[person.Gender][]string{
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

	ageInDays := age*daysPerYearGregorianE4 + (rand.IntN(365) * 10000)

	return person.Person{
		UUID:      uuid.Must(uuid.NewV4()),
		Name:      name,
		Birthday:  date.Today() - date.Date(ageInDays),
		Gender:    gender,
		Addresses: nil,
	}
}

// RandomAdult invents a person who is at least 18 years old.
func RandomAdult(gender person.Gender) person.Person {
	titleIdx := rand.IntN(len(titles[gender]))
	adult := randomPerson(gender, rand.IntN(92)+18)
	adult.Title = titles[gender][titleIdx]
	return adult
}

// RandomChild invents a person who is less than 18 years old.
func RandomChild(gender person.Gender) person.Person {
	return randomPerson(gender, rand.IntN(18))
}

//-------------------------------------------------------------------------------------------------

//var addresses = []person.Address{
//	{
//		Address:  "Flat 1a, Tŷ Gwyn, 13 Stryd y Dŵr",
//		City:     "Bridgend",
//		Postcode: "CF314AD ",
//		Country:  country.GB,
//	},
//	{
//		Address:  "Bryn Ogwen, England Rd N",
//		City:     "Caernarfon",
//		Postcode: "LL55 1HS",
//		Country:  country.GB,
//	},
//	{
//		Address:  "22 MacDonald Dr",
//		City:     "Stirling",
//		Postcode: "FK7 9ER",
//		Country:  country.GB,
//	},
//	{
//		Address:  "2 Stockton Gardens",
//		City:     "London",
//		Postcode: "N17 7HY",
//		Country:  country.GB,
//	},
//	{
//		Address:  "10 Leys Ave",
//		City:     "Rothwell",
//		Postcode: "NN14 6JF",
//		Country:  country.GB,
//	},
//	{
//		Address:  "62 Pankhurst Cres",
//		City:     "Stevenage",
//		Postcode: "SG2 0QF",
//		Country:  country.GB,
//	},
//	{
//		Address:  "6 St Mary's Place",
//		City:     "Chippenham",
//		Postcode: "SN15 1EN",
//		Country:  country.GB,
//	},
//	{
//		Address:  "62 Oxford St",
//		City:     "Burnham-on-Sea",
//		Postcode: "TA81EW",
//		Country:  country.GB,
//	},
//	{
//		Address:  "29, Plymouth Rd",
//		City:     "Penarth",
//		Postcode: "CF64 3DA",
//		Country:  country.GB,
//	},
//	{
//		Address:  "15 Thrushel Close",
//		City:     "Swindon",
//		Postcode: "SN25 3PP",
//		Country:  country.GB,
//	},
//	{
//		Address:  "82 Woodland Rd",
//		City:     "Wolverhampton",
//		Postcode: "WV3 8AW",
//		Country:  country.GB,
//	},
//}
