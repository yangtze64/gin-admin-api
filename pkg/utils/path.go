package utils

import (
	"os"
	"path"
)

func GetRootPath() (string, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return rootPath, nil
}

func GetRuntimePath() (string, error) {
	rootPath, err := GetRootPath()
	if err != nil {
		return "", err
	}
	return path.Join(rootPath, "runtime"), nil
}

func GetConfigFilepath() (string, error) {
	rootPath, err := GetRootPath()
	if err != nil {
		return "", err
	}
	return path.Join(rootPath, "config.yaml"), nil
}

// IsFile 是否是文件
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

// FileIsExist 文件是否存在
func FileIsExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}
