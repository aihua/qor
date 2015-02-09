package media_library

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"os"

	"mime/multipart"
)

var ErrNotImplemented = errors.New("not implemented")

type Base struct {
	Url        string
	Valid      bool
	FileName   string
	FileHeader *multipart.FileHeader
	CropOption *CropOption
	Reader     io.Reader
}

func (b *Base) Scan(value interface{}) error {
	switch v := value.(type) {
	case []*multipart.FileHeader:
		if len(v) > 0 {
			file := v[0]
			b.FileHeader, b.FileName, b.Valid = file, file.Filename, true
		}
	case []uint8:
		b.Url, b.Valid = string(v), true
	default:
		fmt.Errorf("unsupported driver -> Scan pair for MediaLibrary")
	}
	return nil
}

func (b Base) Value() (driver.Value, error) {
	if b.Valid {
		return b.FileName, nil
	}
	return nil, nil
}

func (b Base) URL(...string) string {
	return b.Url
}

func (b Base) String() string {
	return b.URL()
}

func (b Base) GetFileName() string {
	return b.FileName
}

func (b Base) GetFileHeader() *multipart.FileHeader {
	return b.FileHeader
}

func (b Base) GetURLTemplate(tag string) (path string) {
	if path = parseTagOption(tag).Get("url"); path == "" {
		path = "/system/{{class}}/{{primary_key}}/{{column}}/{{basename}}.{{nanotime}}.{{extension}}"
	}
	return
}

func (b *Base) SetCropOption(option *CropOption) {
	b.CropOption = option
}

func (b Base) Retrieve(url string) (*os.File, error) {
	return nil, ErrNotImplemented
}

func (b Base) Crop() error {
	return ErrNotImplemented
}
