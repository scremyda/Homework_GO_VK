package utilities

import (
	"errors"
	"flag"
	"fmt"
	"uniq/processing"
)

func SetArgs() (uniq.Options, error) {
	var options uniq.Options

	flag.BoolVar(&options.CFlagStated, "c", false, "Count the number of occurrences of a string and"+
		" print this number before the string separated by a space.")
	flag.BoolVar(&options.DFlagStated, "d", false, "Output only those lines that are repeated "+
		"in the input data.")
	flag.BoolVar(&options.UFlagStated, "u", false, "Output only those lines that are not "+
		"repeated in the input data.")
	flag.IntVar(&options.FFlagStated, "f", 0, "Ignore the first numFields of fields in a row. "+
		"A field in a string is a non-empty set of characters separated by a space.")
	flag.IntVar(&options.SFlagStated, "s", 0, "Ignore the first numFields of fields in a row. "+
		"A field in a string is a non-empty set of characters separated by a space.")
	flag.BoolVar(&options.IFlagStated, "i", false, "Ignore the case of letters.")

	flag.Parse()

	if len(flag.Args()) > 2 {
		err := fmt.Errorf("Too many args are given: %d ", len(flag.Args()))
		flag.PrintDefaults()
		return options, err
	}
	if (options.CFlagStated && options.DFlagStated) ||
		(options.CFlagStated && options.UFlagStated) ||
		(options.DFlagStated && options.UFlagStated) {
		err := errors.New("parameters 'c', 'd' and 'u' cannot be used at the same time")
		flag.PrintDefaults()
		return options, err
	}
	if options.SFlagStated < 0 || options.FFlagStated < 0 {
		err := errors.New("num_fields or num_chars cannot be negative")
		flag.PrintDefaults()
		return options, err
	}

	return options, nil
}
