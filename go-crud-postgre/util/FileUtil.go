package util

import (
	"os"
	"path/filepath"
)

func CheckFileSize(file *os.File) int64 {
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	return fileInfo.Size()
}

func CheckFileSizeByPath(p_FilePath string) int64 {
	fileInfo, err := os.Stat(p_FilePath)
	if err != nil {
		panic(err)
	}
	return fileInfo.Size()
}

func ListDirectoryContent(p_PathToScan string) []string {
	var result []string
	err := filepath.Walk(p_PathToScan, func(path string, info os.FileInfo, err error) error {
		result = append(result, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return result
}
