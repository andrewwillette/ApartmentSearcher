package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	handleCliConfigs()

	aptmts := []apartment{}
	for _, aptmt := range uliHttpQueries {
		resp, err := http.Get(aptmt.url)
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		html := string(body)
		aptmts = append(aptmts, getApartments(html)...)
	}
	display(aptmts)
}

// display frontend display of apartments
func display(apartments []apartment) {
	for _, apt := range apartments {
		fmt.Printf("%+v\n", apt)
	}
}
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func (a apartmentQuery) printExpected() {
	if verbose {
		log.Println("Expected output for apartment query : -------------- ")
		log.Println(a.expectedOutput)
	}
}

var uliHttpQueries = []apartmentQuery{
	{
		url: "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=2133&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=August%2C+2022",
	}, {
		url: "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=4&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=August%2C+2022",
	}, {
		url: "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=1980&field_bedrooms_value%5B%5D=1_bed&field_bedrooms_value%5B%5D=1_bed_den&field_bedrooms_value%5B%5D=2_bed&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=July%2C+2022",
	}, {
		url: "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=1980&field_property_target_id%5B%5D=2133&field_bedrooms_value%5B%5D=1_bed&field_bedrooms_value%5B%5D=1_bed_den&field_bedrooms_value%5B%5D=2_bed&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=August%2C+2022",
	},
}

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "get apartment details",
	Short: "apartment buying quickly",
	Long:  `Get apartments quickly but longer text`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func handleCliConfigs() {
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
