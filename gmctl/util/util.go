package util

import (
	"io"
	"os"
	"path/filepath"
)

func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}

// 复制文件夹及其内容
func CopyDir(sourceDir, destinationDir string) error {
	// 确保目标文件夹存在
	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		return err
	}

	// 获取源文件夹中的文件和子文件夹
	files, err := os.ReadDir(sourceDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		sourcePath := filepath.Join(sourceDir, file.Name())
		destinationPath := filepath.Join(destinationDir, file.Name())

		if file.IsDir() {
			// 如果是子文件夹，递归调用复制函数
			if err := CopyDir(sourcePath, destinationPath); err != nil {
				return err
			}
		} else {
			// 如果是文件，执行文件复制
			if err := copyFile(sourcePath, destinationPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// 复制单个文件
func copyFile(sourceFile, destinationFile string) error {
	source, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}
