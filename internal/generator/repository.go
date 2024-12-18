package generator

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type IFileRepository interface {
	InsertFile(f *XlsxFile) error
	GetFilesForPeriod(from time.Time, to time.Time) ([]XlsxFile, error)
	GetFileById(id uint64) (XlsxFile, error)
	DeleteFile(id uint64) error
}

type XlsxFileRepository struct {
	db *sqlx.DB
}

func NewXlsxFileRepository(db *sqlx.DB) *XlsxFileRepository {
	return &XlsxFileRepository{db: db}
}

func (r XlsxFileRepository) InsertFile(f *XlsxFile) error {
	var newId uint64
	err := r.db.Get(&newId,
		"INSERT INTO files(created_at, file_name) VALUES($1, $2) RETURNING id",
		f.CreatedAt, f.Filename)
	if err != nil {
		return fmt.Errorf("failed to insert new file report")
	}
	f.ID = newId
	return nil
}

func (r XlsxFileRepository) GetFilesForPeriod(from time.Time, to time.Time) ([]XlsxFile, error) {
	var result []XlsxFile
	err := r.db.Select(&result, "SELECT * FROM files WHERE created_at BETWEEN $1 AND $2 ORDER BY id",
		from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get files: %s", err.Error())
	}
	return result, nil
}

func (r XlsxFileRepository) GetFileById(id uint64) (XlsxFile, error) {
	var result XlsxFile = XlsxFile{}
	err := r.db.Get(&result, "SELECT * FROM files WHERE id = $1", id)
	if err != nil {
		return result, fmt.Errorf("failed to get files: %s", err.Error())
	}
	return result, nil
}

func (r XlsxFileRepository) DeleteFile(id uint64) error {
	_, err := r.GetFileById(id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("DELETE FROM files WHERE id=$1",
		id)
	if err != nil {
		return fmt.Errorf("failed to delete file with id = %d", id)
	}
	return nil
}
