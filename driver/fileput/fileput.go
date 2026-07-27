package fileput

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"goftp.io/server/v2"
	file_driver "goftp.io/server/v2/driver/file"
)

type Driver struct {
	RootPath string
	fd       server.Driver
}

func NewDriver(rootPath string) (server.Driver, error) {
	var err error
	rootPath, err = filepath.Abs(rootPath)
	if err != nil {
		return nil, err
	}

	fd, err := file_driver.NewDriver(rootPath)
	if err != nil {
		return nil, err
	}

	return &Driver{rootPath, fd}, nil
}

func (driver *Driver) Stat(ctx *server.Context, path string) (os.FileInfo, error) {
	return driver.fd.Stat(ctx, path)
}

func (driver *Driver) ListDir(ctx *server.Context, path string, callback func(os.FileInfo) error) error {
	return errors.New("Nope")
}
func (driver *Driver) DeleteDir(ctx *server.Context, path string) error {
	return errors.New("Nope")
}
func (driver *Driver) DeleteFile(ctx *server.Context, path string) error {
	return errors.New("Nope")
}
func (driver *Driver) Rename(ctx *server.Context, fromPath string, toPath string) error {
	return errors.New("Nope")
}
func (driver *Driver) MakeDir(ctx *server.Context, path string) error {
	return errors.New("Nope")
}
func (driver *Driver) GetFile(ctx *server.Context, path string, offset int64) (int64, io.ReadCloser, error) {
	return 0, nil, errors.New("Nope")
}

func (driver *Driver) PutFile(ctx *server.Context, destPath string, data io.Reader, offset int64) (int64, error) {
	bytes, err := driver.fd.PutFile(ctx, destPath, data, offset)
	if err != nil {
		return 0, err
	}

	return bytes, nil
}
