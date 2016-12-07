package main

import (
	"bytes"
	"fmt"
)

type Table struct {
	Schema string
	Name string
}
type TableSet []*Table

func (ts TableSet) String() string {
	var buffer bytes.Buffer

	for _, t := range ts {
		buffer.WriteString(fmt.Sprintf("Schema:%q;\nName: %q.\n\n",t.Schema, t.Name))
	}

	return buffer.String()
}

func loadTables() (TableSet, error) {
	const query = baseQuery + `SELECT object_schema, object_name from INFO where object_type = 'TABLE'`
	rows, err := pg.Query(query)
	if err != nil {
		return nil, err
	}

	tbs := make([]*Table, 0, 100)

	defer rows.Close()
	for rows.Next() {
		var tb = new(Table)
		err = rows.Scan(&tb.Schema, &tb.Name)
		if err != nil {
			return nil, err
		}
		tbs = append(tbs, tb)
	}
	return TableSet(tbs), nil
}

type Columns struct {

}



