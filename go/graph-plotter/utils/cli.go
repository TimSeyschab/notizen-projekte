package utils

import (
	"flag"
	"os"
	"strconv"
)

type CLIParams struct {
	FunctionType string
	Params       []float64
}

func ParseCLI() CLIParams {
	funcType := flag.String("type", "linear", "Function type: linear, quadratic, cubic")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	params := parseArgsToFloats(args)

	return CLIParams{
		FunctionType: *funcType,
		Params:       params,
	}
}

func parseArgsToFloats(args []string) []float64 {
	var floats []float64
	for _, arg := range args {
		f, _ := strconv.ParseFloat(arg, 64) // implement conversion here
		floats = append(floats, f)
	}
	return floats
}
