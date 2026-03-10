package main

const (
	ROW_COUNT       = 4561200
	EXT_CODE_COUNT  = 1400
	B_CODE_COUNT    = 13
	D_TYPE_COUNT    = 10
	OUTPUT_FILENAME = "sample.csv"
)

type output_row struct {
	EXT_CODE string
	B_CODE   string
	D_TYPE   int
	N_TIME   int
	TELEGRAM string
}

var HEADER = []string{
	"EXT_CODE",
	"B_CODE",
	"D_TYPE",
	"N_TIME",
	"TELEGRAM",
}
