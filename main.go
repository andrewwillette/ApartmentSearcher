package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type uliApartmentQuery struct {
	url, apartmentName, date string
}

var apartments = []uliApartmentQuery{
	{
		apartmentName: "1722 Monroe",
		url:           "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=2133&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=August%2C+2022",
		date:          "August",
	},
	{
		apartmentName: "Tobacco Lofts",
		url:           "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=4&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=August%2C+2022",
		date:          "August",
	},
	{
		apartmentName: "Tobacco Lofts",
		url:           "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=4&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=September%2C+2022",
		date:          "September",
	},
}

func main() {
	apartmentsExist(apartments)
}

func apartmentsExist(apartments []uliApartmentQuery) {
	for _, aptmt := range apartments {
		resp, err := http.Get(aptmt.url)
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))

		if err != nil {
			log.Fatal(err)
		}

		noApartments, err := regexp.MatchString("we currently do not have any available units that meet this spec", string(body))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Apartment: %s\nMonth: %s\nAvailable: %t\n\n", aptmt.apartmentName, aptmt.date, !noApartments)
	}
}
