package oss

import "io"

type Oss interface {
	UploadFile(reader io.Reader, filename string, size int64) (interface{}, error)
	DownloadFile()
}
