package helpers

import (
	"regexp"
	"strings"
)

/*
   Converts all items of a string slice to lowercase equivalent.
*/
func ToLowerSlice(slc []string) []string {
	loweredSlc := []string{}
	for i := 0; i < len(slc); i++ {
		loweredSlc = append(loweredSlc, strings.ToLower(slc[i]))
	}
	return loweredSlc
}

/*
	Converts given string to acronym (Eg. Lord of the Rings: LOTR)
*/
func ToAcronym(source string) string {
	reg, _ := regexp.Compile(`\B.|\P{L}`)
	abbr := reg.ReplaceAllString(source, "")
	return strings.ToUpper(abbr)
}
