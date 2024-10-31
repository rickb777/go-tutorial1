package person

import (
	"github.com/rickb777/date/v2"
	"gotest.tools/v3/assert"
	"testing"
)

func TestPersonAge(t *testing.T) {
	today = func() date.Date { return date.New(2024, 11, 1) }

	cases := []struct {
		birthday date.Date
		age      int
		adult    bool
	}{
		{birthday: date.New(2020, 3, 23), age: 4, adult: false},
		{birthday: date.New(2000, 10, 31), age: 24, adult: true},
		{birthday: date.New(2000, 11, 2), age: 23, adult: true},
	}

	for i, c := range cases {

		andy := Person{
			Birthday: c.birthday,
		}

		assert.Equal(t, andy.Age(), c.age, i)
	}
}
