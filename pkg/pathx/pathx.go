package pathx

import "os"

func RootPath() (string, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return rootPath, nil
}

// FileExist 文件是否存在
func FileExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}
