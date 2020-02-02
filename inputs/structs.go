package inputs

import (
	"fmt"
	"reflect"
	"strings"
)

// StructsInput represents a record producing input from a Structs formatted file or pipe.
type StructsInput struct {
	options         *StructsInputOptions
	header          []string
	structs         []interface{}
	name            string
}

type StructsInputOptions struct {
	// Separator is the rune that fields are delimited by.
	Separator rune
	Structs         []interface{}
}

// NewStructsInput sets up a new StructsInput, the first row is read when this is run.
// If there is a problem with reading the first row, the error is returned.
// Otherwise, the returned StructsInput can be reliably consumed with ReadRecord()
// until ReadRecord() returns nil.
func NewStructsInput(opts *StructsInputOptions) (*StructsInput, error) {
	StructsInput := &StructsInput{
		options: opts,
	}
	StructsInput.structs = opts.Structs
	headerErr := StructsInput.readHeader()

	if headerErr != nil {
		return nil, headerErr
	}

	return StructsInput, nil
}

// Name returns the name of the Structs being read.
// By default, either the base filename or 'pipe' if it is a unix pipe
func (StructsInput *StructsInput) Name() string {
	return StructsInput.name
}

// SetName overrides the name of the Structs
func (StructsInput *StructsInput) SetName(name string) {
	StructsInput.name = name
}

// ReadRecord reads a single record from the Structs. Always returns successfully.
// If the record is empty, an empty []string is returned.
// Record expand to match the current row size, adding blank fields as needed.
// Records never return less then the number of fields in the first row.
// Returns nil on EOF
// In the event of a parse error due to an invalid record, it is logged, and
// an empty []string is returned with the number of fields in the first row,
// as if the record were empty.
//
// In general, this is a very tolerant of problems Structs reader.
func (StructsInput *StructsInput) ReadRecord() []string {
	var row []string

	if 0 == len(StructsInput.structs) {
	}
	row = make([]string, 0)
	for index, item := range StructsInput.structs {
		StructType := reflect.TypeOf(item)
		if StructType.Kind() != reflect.Struct {
		}
		StructValue := reflect.ValueOf(item)
		rowSlice := make([]string, 0)
		rowSlice[0] = string(index)
		for i := 0; i < StructType.NumField(); i++ {
			rowSlice = append(rowSlice, fmt.Sprint(StructValue.Field(i).Interface()))
		}
		row = append(row, strings.Join(rowSlice, string(StructsInput.options.Separator)))
	}
	return row
}

func (StructsInput *StructsInput) readHeader() error {

	if 0 == len(StructsInput.structs) {
	}

	StructType := reflect.TypeOf(StructsInput.structs[0])
	if StructType.Kind() != reflect.Struct {
	}

	StructsInput.header = make([]string, 0)
	StructsInput.header[0] = StructsInput.name + "_id"
	for i := 0; i < StructType.NumField(); i++ {
		t := StructType.Field(i)
		column := t.Tag.Get("column")
		StructsInput.header = append(StructsInput.header, column)
	}

	return nil
}

// Header returns the header of the StructsInput. Either the first row if a header
// set in the options, or c#, where # is the column number, starting with 0.
func (StructsInput *StructsInput) Header() []string {
	return StructsInput.header
}
