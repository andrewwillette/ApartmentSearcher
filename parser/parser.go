package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/andrewwillette/ApartmentSearcher/cmd"
)


type uliApartmentQuery struct {
	url,
	apartmentName,
	date,
	expectedOutput string
}

func (a uliApartmentQuery) printExpected() {
	log.Println("Expected output for apartment query : -------------- ")
	log.Println(a.expectedOutput)
}

var apartments = []uliApartmentQuery{
	{
		apartmentName:  "1722 Monroe",
		url:            "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=2133&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=August%2C+2022",
		date:           "August",
		expectedOutput: "None available",
	},
	// {
	// 	apartmentName: "Tobacco Lofts",
	// 	url:           "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=4&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=August%2C+2022",
	// 	date:          "August",
	// 	expectedOutput: `2 Bedrooms 1269 Sq. Feet $2310/month
	// 1 Bedroom 1196 Sq. Feet $2115/month`,
	// },
	{
		apartmentName: "Pressman",
		url:           "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=1980&field_bedrooms_value%5B%5D=1_bed&field_bedrooms_value%5B%5D=1_bed_den&field_bedrooms_value%5B%5D=2_bed&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=July%2C+2022",
		date:          "July",
		expectedOutput: `2 Bedrooms
1370 Sq. Feet
$2980/month

2 Bedrooms
1320 Sq. Feet
$3060/month

2 Bedrooms
1250 Sq. Feet
$3345/month

2 Bedrooms
1390 Sq. Feet
$3540/month
`,
	},
}

func main() {
	apartmentsExist(apartments)
}

func printAvailableApartments(rawHtml string) {
}

// printApartmentDetails print details for apartment number apartmentNumber
func printApartmentDetails(apartmentNumber int, rawHtml string) {
	r, _ := regexp.Compile("avail-date\">Available\\s*\\d*/\\d*/\\d*")
	htmlParsed1 := r.FindString(rawHtml)
	r2, _ := regexp.Compile("\\d*/\\d*/\\d*")
	availDate := r2.FindString(htmlParsed1)
	fmt.Printf("Apartment: Available Date: %s\n", availDate)
}

// return true if available apartments
func checkAvailableApartments(rawHtml string) bool {
	noAvailableApartments, err := regexp.MatchString("we currently do not have any available units that meet this spec", rawHtml)
	if err != nil {
		return false
	}
	return !noAvailableApartments
}

func apartmentsExist(apartments []uliApartmentQuery) {
	for _, aptmt := range apartments {
		resp, err := http.Get(aptmt.url)
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		html := string(body)
		apartmentsAvailable := checkAvailableApartments(html)
		if apartmentsAvailable {
			printApartmentDetails(0, html)
			aptmt.printExpected()

			// fmt.Println(decodedValue)
		} else {
			// fmt.Printf("Apartment: %s\nMonth: %s\nAvailable: %t\n\n", aptmt.apartmentName, aptmt.date, !noAvailableApartments)
		}
	}
}
