package main

import (
	"regexp"
	"strconv"
)

// getApartments parses html in string format and returns a slice of apartment
func getApartments(html string) []apartment {
	apartments := []apartment{}
	for _, apmtHtml := range parseSingleApmtHtml(html) {
		apartment := apartment{}
		apartment.availDate = getAvailableDate(apmtHtml)
		apartment.rent = getRent(apmtHtml)
		apartment.bedrooms = getBedrooms(apmtHtml)
		apartment.sqFootage = getSqFootage(apmtHtml)
		apartment.unitTitle = getUnitTitle(apmtHtml)
		apartments = append(apartments, apartment)
	}
	return apartments
}

func getUnitTitle(html string) string {
	r, _ := regexp.Compile(`unit-title">.*`)
	sqFootstage1 := r.FindString(html)
	r, _ = regexp.Compile(`>.*<`)
	result := r.FindString(sqFootstage1)
	r, _ = regexp.Compile(`[^>][^<]*`)
	result2 := r.FindString(result)
	return result2
}

// getSqFootage parse html for square footage
// html must represent single apartment
func getSqFootage(html string) int {
	r, _ := regexp.Compile(`\d{0,4}\sSq.\sFeet`)
	sqFootstage1 := r.FindString(html)
	r, _ = regexp.Compile(`\d{0,4}`)
	result := r.FindString(sqFootstage1)
	intVar, _ := strconv.Atoi(result)
	return intVar
}

// getBedrooms parse html string for total bedrooms
// must be single apartment
func getBedrooms(html string) int {
	r, _ := regexp.Compile(`\d\sBedroom`)
	bedroom1 := r.FindString(html)
	r, _ = regexp.Compile(`\d{1,2}`)
	result := r.FindString(bedroom1)
	intVar, _ := strconv.Atoi(result)
	return intVar
}

// getRent parse html string for rent
// must be single apartment
func getRent(html string) int {
	r, _ := regexp.Compile(`\$\d{1,4}`)
	rent1 := r.FindString(html)
	r, _ = regexp.Compile(`\d{1,4}`)
	result := r.FindString(rent1)
	intVar, _ := strconv.Atoi(result)
	return intVar
}

// parseSingleApmtHtml parses whole html doc into individual apartments
// returns array of html strings containing values per apartment
func parseSingleApmtHtml(html string) []string {
	r, _ := regexp.Compile(".*\"avail-date\"(.*\\s){0,10}")
	return r.FindAllString(html, 10)
}

// getAvailableDate parses html and returns available date of apartment
func getAvailableDate(html string) string {
	r, _ := regexp.Compile(`\d{1,2}/\d{1,2}/\d{1,4}`)
	return r.FindString(html)
}
