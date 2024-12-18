package generator

import "time"

type XlsxFile struct {
	ID        uint64
	CreatedAt time.Time `db:"created_at"`
	Filename  string
}
