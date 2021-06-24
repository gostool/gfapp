package oss

import (
	"path"
	"path/filepath"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetUUID() string {
	u := uuid.NewV4()
	return u.String()
}

func GetUploadDirByFileExt(fileExt string) (dir string) {
	dir = "upload"
	extDir := ""
	switch fileExt {
	case ".mp3":
		extDir = "mp3"
	case ".mp4":
		extDir = "mp4"
	case ".jpg", "jpeg", "png", ".webp", ".gif":
		extDir = "img"
	default:
		extDir = "tmp"
	}
	return filepath.Join(dir, extDir)
}

func NewFilePath(filename string) (newFilePath string) {
	fid := GetUUID()
	fileExt := path.Ext(filename)
	fileUploadDir := GetUploadDirByFileExt(fileExt)
	newFilePath = filepath.Join(fileUploadDir, time.Now().Format("2006-01-02")) + "/" + fid + fileExt
	return
}
