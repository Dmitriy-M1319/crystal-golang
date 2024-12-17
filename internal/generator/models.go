package generator

import "time"

type XlsxFile struct {
	ID        uint64
	CreatedAt time.Time
	Filename  string
}
