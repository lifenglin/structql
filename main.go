package textql_structs

import (
	"fmt"
	"log"
	"github.com/lifenglin/textql-structs/inputs"
	"github.com/lifenglin/textql-structs/outputs"
	"github.com/dinedal/textql/storage"
	"github.com/google/uuid"
	"strings"
	"sync"
)

var SQLite3Storage *storage.SQLite3Storage

func WhereStructs(structs []interface{}, where string, params []interface{}) ([]int64, error) {
	var once sync.Once
	onceFunc := func() {
		SQLite3Storage = storage.NewSQLite3StorageWithDefaults()
	}
	once.Do(onceFunc)

	inputOpts := &inputs.StructsInputOptions{
		Separator: ',',
		Structs: structs,
	}

	input, inputErr := inputs.NewStructsInput(inputOpts)
	name := uuid.New()
	input.SetName(name.String())

	if inputErr != nil {
		log.Printf("Unable to load %v\n", structs)
		return nil, inputErr
	}
	SQLite3Storage.LoadInput(input)
	statement := fmt.Sprintf(strings.Replace(where, "?", "%s", -1), params...)
	sqlQuery := fmt.Sprintf("SELECT %s_id FROM %s", name.String(), name.String(), statement)

	queryResults, queryErr := SQLite3Storage.ExecuteSQLString(sqlQuery)
	if queryErr != nil {
		return nil, queryErr
	}

	outputOpts := &outputs.StructsOutputOptions{
	}

	output := outputs.NewStructsOutput(outputOpts)

	index, indexErr := output.Index(queryResults)

	if indexErr != nil {
		return nil, indexErr
	}
	return index, nil

}