// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.5.1

package person

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
	"strings"
)

// AllGenders lists all 4 values in order.
var AllGenders = []Gender{
	Unspecified, Male, Female, Other,
}

// AllGenderEnums lists all 4 values in order.
var AllGenderEnums = enum.IntEnums{
	Unspecified, Male, Female, Other,
}

const (
	genderEnumStrings = "UnspecifiedMaleFemaleOther"
	genderEnumInputs  = "unspecifiedmalefemaleother"
)

var (
	genderEnumIndex = [...]uint16{0, 11, 15, 21, 26}
)

// Ordinal returns the ordinal number of a Gender. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Gender) Ordinal() int {
	switch v {
	case Unspecified:
		return 0
	case Male:
		return 1
	case Female:
		return 2
	case Other:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Gender, which is
// the same as the const identifier but without prefix or suffix.
func (v Gender) String() string {
	o := v.Ordinal()
	return v.toString(o, genderEnumStrings, genderEnumIndex[:])
}

func (v Gender) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllGenders) {
		return fmt.Sprintf("Gender(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Gender is one of the defined constants.
func (v Gender) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Gender) Int() int {
	return int(v)
}

var invalidGenderValue = func() Gender {
	var v Gender
	for {
		if !slices.Contains(AllGenders, v) {
			return v
		}
		v++
	} // AllGenders is a finite set so loop will terminate eventually
}()

// GenderOf returns a Gender based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Gender is returned.
func GenderOf(v int) Gender {
	if 0 <= v && v < len(AllGenders) {
		return AllGenders[v]
	}
	return invalidGenderValue
}

// Parse parses a string to find the corresponding Gender, accepting one of the string values or
// a number. It is used by AsGender.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Gender)
//	err := v.Parse(s)
//	...  etc
func (v *Gender) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := genderTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Gender) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Gender(num)
		return v.IsValid()
	}
	return false
}

func (v *Gender) parseFallback(in, s string) error {
	if v.parseString(s, genderEnumInputs, genderEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised gender")
}

func (v *Gender) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllGenders[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// genderTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var genderTransformInput = func(in string) string {
	return strings.ToLower(in)
}

// AsGender parses a string to find the corresponding Gender, accepting either one of the string values or
// a number. It wraps Parse.
// The input case does not matter.
func AsGender(s string) (Gender, error) {
	var v = new(Gender)
	err := v.Parse(s)
	return *v, err
}

// MustParseGender is similar to AsGender except that it panics on error.
// The input case does not matter.
func MustParseGender(s string) Gender {
	v, err := AsGender(s)
	if err != nil {
		panic(err)
	}
	return v
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Gender) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, genderEnumStrings, genderEnumIndex[:])
}

func (v Gender) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Gender) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Gender) invalidError() error {
	return fmt.Errorf("%d is not a valid gender", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The identifier representation is chosen according to -marshaljson.
func (v Gender) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, genderEnumStrings, genderEnumIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Gender) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Gender) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := genderTransformInput(in)

	if v.parseString(s, genderEnumInputs, genderEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised gender")
}

// genderMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var genderMarshalNumber = func(v Gender) string {
	return strconv.FormatInt(int64(v), 10)
}
