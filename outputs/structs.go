package outputs

import (
	"database/sql"
)

// StructsOutput represents a TextQL output that transforms sql.Rows into Structs formatted
// string data using encoding/Structs
type StructsOutput struct {
	options         *StructsOutputOptions
	header          []string
	name            string
}

// StructsOutputOptions define options that are passed to encoding/Structs for formatting
// the output in specific ways.
type StructsOutputOptions struct {
}

// NewStructsOutput returns a new StructsOutput configured per the options provided.
func NewStructsOutput(opts *StructsOutputOptions) *StructsOutput {
	StructsOutput := &StructsOutput{
		options: opts,
	}
	return StructsOutput
}

func (StructsOutput *StructsOutput) Index(rows *sql.Rows) ([]int64, error) {
	defer rows.Close()
	index := make([]int64, 0)
	var id int64
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {

		}
		index = append(index, id)
	}
	return index, nil
}
