package db

import "database/sql"

type Transactional interface {
	Begin() (*sql.Tx, error)
}
