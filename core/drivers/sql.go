package drivers

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// SQLRow :
type SQLRow map[string]string

// SQLDatabaseDriver :
type SQLDatabaseDriver struct {
	ctx       context.Context
	db        *sql.DB
	connected bool
}

// NewSQLDatabaseDriver :
func NewSQLDatabaseDriver() *SQLDatabaseDriver {
	data := &SQLDatabaseDriver{}
	data.ctx = context.TODO()
	return data
}

// Connect :
func (s *SQLDatabaseDriver) Connect(driverName string, dsn string) error {
	var err error
	s.db, err = sql.Open(driverName, dsn)
	if err != nil {
		return err
	}
	return nil
}

// IsConnected :
func (s *SQLDatabaseDriver) IsConnected() bool {
	return s.connected
}

// ExecuteQuery :
func (s *SQLDatabaseDriver) ExecuteQuery(query string) (int64, error) {
	res, err := s.db.Exec(query)
	if err != nil {
		return 0, err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return int64(0), err
	}
	return num, nil
}

// ExecuteSelectQuery :
func (s *SQLDatabaseDriver) ExecuteSelectQuery(query string) ([]SQLRow, error) {
	res, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	rows := make([]SQLRow, 0)
	colNames, _ := res.Columns()
	for res.Next() {
		row := make([]*string, len(colNames))
		scanRow := make([]interface{}, len(row))
		for i := range row {
			v := ""
			row[i] = &v
			scanRow[i] = row[i]
		}
		err := res.Scan(scanRow...)
		if err != nil {
			return nil, err
		}
		sqlRow := make(SQLRow)
		for i := range row {
			key := colNames[i]
			value := row[i]
			sqlRow[key] = *value
		}
		rows = append(rows, sqlRow)
	}
	return rows, nil
}

// Disconnect :
func (s *SQLDatabaseDriver) Disconnect() {
	s.db.Close()
}
