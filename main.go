package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"

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
	sort.SliceStable(aptmts, func(i, j int) bool {
		switch sortedInput {
		case rent:
			return aptmts[i].rent < aptmts[j].rent
		case sqFeet:
			return aptmts[i].sqFootage < aptmts[j].sqFootage
		case availDate:
			return aptmts[i].availDate < aptmts[j].availDate
		default:
			return true
		}
	})
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
		url: "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=4&field_property_target_id%5B%5D=1980&field_property_target_id%5B%5D=2133&field_bedrooms_value%5B%5D=studio&field_bedrooms_value%5B%5D=1_bed&field_bedrooms_value%5B%5D=1_bed_den&field_bedrooms_value%5B%5D=2_bed&field_bedrooms_value%5B%5D=2_bed_den&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=July%2C+2022",
	}, {
		url: "https://www.uli.com/residential/apartment-search?field_property_target_id%5B%5D=4&field_property_target_id%5B%5D=1980&field_property_target_id%5B%5D=2133&field_bedrooms_value%5B%5D=studio&field_bedrooms_value%5B%5D=1_bed&field_bedrooms_value%5B%5D=1_bed_den&field_bedrooms_value%5B%5D=2_bed&field_bedrooms_value%5B%5D=2_bed_den&field_available_date_value_1%5Bvalue%5D%5Bdate%5D=August%2C+2022",
	},
}

type aptmtSortable string

const (
	rent      aptmtSortable = "r"
	availDate aptmtSortable = "d"
	sqFeet    aptmtSortable = "s"
)

// String is used both by fmt.Print and by Cobra in help text
func (e *aptmtSortable) String() string {
	return string(*e)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (e *aptmtSortable) Set(v string) error {
	switch v {
	case "r", "d", "s":
		*e = aptmtSortable(v)
		return nil
	default:
		return errors.New(`must be one of "r", "d", or "s"`)
	}
}

// Type is only used in help text
func (e *aptmtSortable) Type() string {
	return "aptmtSortable"
}

var verbose bool
var sortedInput aptmtSortable

var rootCmd = &cobra.Command{
	Use:   "get apartment details",
	Short: "apartment buying quickly",
	Long:  `Get apartments quickly but longer text`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func handleCliConfigs() {
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.Flags().VarP(&sortedInput, "sort", "s", `sort by partcular column: "r": rent, "d": availDate, "s": square feet`)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
