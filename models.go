package main

type apartment struct {
  availDate string
  unitTitle string
  bedrooms int
  sqFootage int
  rent int
}

type apartmentQuery struct {
	url,
	apartmentName,
	expectedOutput string
}
