package easy

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//读取zip解压后里面的所有文件
func ReadZip(i io.Reader, savePath string) (f []io.Reader, err error) {
	if savePath == "" {
		savePath = "/runtime/zip/"
	}
	b, err := ioutil.ReadAll(i)
	if err != nil {
		fmt.Println("读取文件失败", err.Error())
		return
	}
	logFilePath := ""
	now := time.Now()
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + savePath + now.Format("200601") + "/"
	}
	BuildCatalog(logFilePath)
	name := now.Format("20060102150405")
	zipFileName := name + ".zip"
	fileName := logFilePath + zipFileName
	err = os.WriteFile(fileName, b, 0)
	if err != nil {
		fmt.Println("写入zip文件失败", err.Error())
		return
	}
	//解压缩 到指定目录
	zipPath := logFilePath + "/" + name
	BuildCatalog(zipPath)
	err = Unzip(fileName, zipPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return ReadFile(zipPath)
}
func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}
			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func ReadFile(fileDir string) (f []io.Reader, err error) {
	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		return
	}
	for _, onefile := range files {
		f1, err1 := os.OpenFile(fileDir+"/"+onefile.Name(), os.O_RDONLY, 0600)
		defer f1.Close()
		if err1 != nil {
			err = err1
			return
		}
		b, err1 := ioutil.ReadAll(f1)
		if err1 != nil {
			err = err1
			return
		}
		f = append(f, bytes.NewReader(b))
	}
	return
}

func BuildCatalog(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		/// fmt.Println("创建目录",path) Mkdir是单级目录
		os.MkdirAll(path, 0777)
		// 再修改权限
		os.Chmod(path, 0777)
	}
}