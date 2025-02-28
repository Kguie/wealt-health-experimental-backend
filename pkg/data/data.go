package data

// State représente un état avec un label et une valeur.
type State struct {
	Label string
	Value string
}

// Department représente un département avec une valeur.
type Department struct {
	Value string
}

// Liste des états.
var States = []State{
	{Label: "Alabama", Value: "AL"},
	{Label: "Alaska", Value: "AK"},
	{Label: "American Samoa", Value: "AS"},
	{Label: "Arizona", Value: "AZ"},
	{Label: "Arkansas", Value: "AR"},
	{Label: "California", Value: "CA"},
	{Label: "Colorado", Value: "CO"},
	{Label: "Connecticut", Value: "CT"},
	{Label: "Delaware", Value: "DE"},
	{Label: "District Of Columbia", Value: "DC"},
	{Label: "Federated States Of Micronesia", Value: "FM"},
	{Label: "Florida", Value: "FL"},
	{Label: "Georgia", Value: "GA"},
	{Label: "Guam", Value: "GU"},
	{Label: "Hawaii", Value: "HI"},
	{Label: "Idaho", Value: "ID"},
	{Label: "Illinois", Value: "IL"},
	{Label: "Indiana", Value: "IN"},
	{Label: "Iowa", Value: "IA"},
	{Label: "Kansas", Value: "KS"},
	{Label: "Kentucky", Value: "KY"},
	{Label: "Louisiana", Value: "LA"},
	{Label: "Maine", Value: "ME"},
	{Label: "Marshall Islands", Value: "MH"},
	{Label: "Maryland", Value: "MD"},
	{Label: "Massachusetts", Value: "MA"},
	{Label: "Michigan", Value: "MI"},
	{Label: "Minnesota", Value: "MN"},
	{Label: "Mississippi", Value: "MS"},
	{Label: "Missouri", Value: "MO"},
	{Label: "Montana", Value: "MT"},
	{Label: "Nebraska", Value: "NE"},
	{Label: "Nevada", Value: "NV"},
	{Label: "New Hampshire", Value: "NH"},
	{Label: "New Jersey", Value: "NJ"},
	{Label: "New Mexico", Value: "NM"},
	{Label: "New York", Value: "NY"},
	{Label: "North Carolina", Value: "NC"},
	{Label: "North Dakota", Value: "ND"},
	{Label: "Northern Mariana Islands", Value: "MP"},
	{Label: "Ohio", Value: "OH"},
	{Label: "Oklahoma", Value: "OK"},
	{Label: "Oregon", Value: "OR"},
	{Label: "Palau", Value: "PW"},
	{Label: "Pennsylvania", Value: "PA"},
	{Label: "Puerto Rico", Value: "PR"},
	{Label: "Rhode Island", Value: "RI"},
	{Label: "South Carolina", Value: "SC"},
	{Label: "South Dakota", Value: "SD"},
	{Label: "Tennessee", Value: "TN"},
	{Label: "Texas", Value: "TX"},
	{Label: "Utah", Value: "UT"},
	{Label: "Vermont", Value: "VT"},
	{Label: "Virgin Islands", Value: "VI"},
	{Label: "Virginia", Value: "VA"},
	{Label: "Washington", Value: "WA"},
	{Label: "West Virginia", Value: "WV"},
	{Label: "Wisconsin", Value: "WI"},
	{Label: "Wyoming", Value: "WY"},
}

// Liste des départements.
var Departments = []Department{
	{Value: "Sales"},
	{Value: "Marketing"},
	{Value: "Engineering"},
	{Value: "Human Resources"},
	{Value: "Legal"},
}
