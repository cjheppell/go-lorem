package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

type unit int

const (
	bytes unit = iota
	kilobytes
	megabytes
	gigabytes
)

type options struct {
	Size string
	File string
}

// NewGoLoremCommand creates a new go-lorem Cobra command with subcommands configured
func NewGoLoremCommand() *cobra.Command {
	opts := &options{}

	goLoremCmd := &cobra.Command{
		Use:   "go-lorem",
		Short: "go-lorem is used to generate a lorem ipsum string of a specified size",
		RunE: func(cmd *cobra.Command, args []string) error {
			size, err := parseSize(opts.Size)
			if err != nil {
				return err
			}
			log.Printf("Size was: %d", size)
			return nil
		},
	}

	goLoremCmd.PersistentFlags().StringVarP(&opts.Size, "size", "s", "", "Requested size of the lorem ipsum string "+
		"(in the form <int>[B|K|M|G] where B, K, M, G represent bytes, kilobytes, megabytes and gigabytes respectively")
	goLoremCmd.MarkPersistentFlagRequired("size")

	goLoremCmd.PersistentFlags().StringVarP(&opts.File, "file", "f", "", "Optional parameter specifying the file to write to, rather than the default of stdout")

	return goLoremCmd
}

func parseSize(size string) (int, error) {
	numFunc := func(c rune) bool {
		return unicode.IsLetter(c)
	}
	unitFunc := func(c rune) bool {
		return unicode.IsNumber(c)
	}

	fields := strings.FieldsFunc(size, numFunc)
	if len(fields) != 1 {
		return 0, fmt.Errorf("Invalid size specified '%s'. Please ensure the size follows the format specified in the usage", size)
	}

	unitSize, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, err
	}

	fields = strings.FieldsFunc(size, unitFunc)
	if len(fields) != 1 || len(fields[0]) > 1 || len(fields[0]) != 1 || !isValidUnit(rune(fields[0][0])) {
		return 0, fmt.Errorf("Invalid unit specified '%s'. Please ensure the size follows the format specified in the usage", size)
	}

	unit, err := parseUnit(rune(fields[0][0]))
	if err != nil {
		return 0, err
	}

	switch unit {
	case bytes:
		return unitSize, nil
	case kilobytes:
		return 1024 * unitSize, nil
	case megabytes:
		return 1024 * 1024 * unitSize, nil
	case gigabytes:
		return 1024 * 1024 * 1024 * unitSize, nil
	default:
		return 0, fmt.Errorf("Unrecognised unit '%s'", unit)
	}
}

func isValidUnit(unit rune) bool {
	validUnits := []rune{'B', 'K', 'M', 'G'}
	for _, elem := range validUnits {
		if unit == elem {
			return true
		}
	}
	return false
}

func parseUnit(unit rune) (unit, error) {
	switch unit {
	case 'B':
		return bytes, nil
	case 'K':
		return kilobytes, nil
	case 'M':
		return megabytes, nil
	case 'G':
		return gigabytes, nil
	default:
		return 0, fmt.Errorf("Unrecognised unit '%s'", unit)
	}
}
