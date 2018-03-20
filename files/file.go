package files

import (
	"io/ioutil"
	"os"
	"io"
	"fmt"
	"path/filepath"
)

const (
	FILE_OVERWRITE = 0
	FILE_APPEND = 1
)

// 判断文件是否存在
func Exists(p string) bool {
	_, err := os.Stat(p)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// WriteBytesToFile writes a byte array to file
func WriteBytesToFile(path string, bytes []byte) error {
	err := ioutil.WriteFile(path, bytes, 0644)
	return err
}

// 写文件
func FilePutContents(filename string,content string,mode int) error {
	var f *os.File
	var err error

	if mode == FILE_OVERWRITE {
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	}else{
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	}

	if err != nil {
		return err
	}
	n, err := f.WriteString(content)
	if err == nil && n < len(content) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

//拷贝文件
func CopyFile(source string, dst string) (err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err == nil {
		sourceInfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dst, sourceInfo.Mode())
		}

	}

	return
}

//拷贝目录
func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceInfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourceFilePointer := filepath.Join(source , obj.Name())

		destinationFilePointer := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourceFilePointer, destinationFilePointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourceFilePointer, destinationFilePointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func RemoveDir(dir string) error {
	return os.RemoveAll(dir)
}
