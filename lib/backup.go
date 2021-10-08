package lib

import (
	"archive/zip"
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

//用map实现set

type Set map[int]string

func (s Set) Has(key uint64) bool {
	key2 := int(key)
	_, ok := s[key2]
	return ok
}

func (s Set) Add(key uint64, value string) {
	key2 := int(key)
	s[key2] = value
}

func (s Set) Delete(key uint64) {
	key2 := int(key)
	delete(s, key2)
}

func Zip(srcDir string, zipFilePath string) error {

	//创建索引文件 记录源路径
	outputFile, outputError := os.OpenFile("restoreToPath.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return outputError
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	//outputWriter.WriteString(outputString)

	//写入源路径
	_, outputError = outputWriter.WriteString(srcDir + "\n")
	if outputError != nil {
		fmt.Printf("An error occurred while writing to file\n")
		return outputError
	}

	//移动索引文件到源路径
	outputError = os.Rename("./restoreToPath.txt", filepath.Join(srcDir, "restoreToPath.txt"))
	if outputError != nil {
		fmt.Printf("An error occurred while moving files to source path\n")
		return outputError
	}

	// 预防：旧文件无法覆盖 删除当前目录下的相同名字的tar文件 tar文件生成在当前目录
	outputError = os.RemoveAll(zipFilePath)
	if outputError != nil {
		fmt.Printf("An error occurred while removing files\n")
		return outputError
	}

	// 创建：zip文件
	zipfile, _ := os.Create(zipFilePath)
	defer zipfile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	sett := make(Set)
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {
		//给walk传入匿名函数
		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info) //获得文件信息
		header.Name = strings.TrimPrefix(path, srcDir+`/`)

		//判断是否是索引文件
		//是则跳过
		if header.Name == "restoreToPath.txt" {
			return nil
		}
		stat, _ := info.Sys().(*syscall.Stat_t)

		//判断是否是软链接
		//是则读出内容 把内容写到索引文件
		if stat.Mode >= 41000 && stat.Mode <= 42000 {
			//读出软链接内容
			s := ""
			for len := 128; ; len *= 2 {
				b := make([]byte, len)
				n, _ := syscall.Readlink(path, b)
				if n < len {
					s = "soft " + path + " " + string(b[0:n]) + "\n"
					//写到索引文件
					//soft+软链接路径+软链接内容写入索引文件
					outputWriter.WriteString(s)
					return nil
				}
			}
		}

		//是管道文件
		if stat.Mode >= 4500 && stat.Mode <= 4600 {
			s := ""
			s = "pipe " + path + "\n"
			//写到索引文件
			//pipe+管道文件名
			outputWriter.WriteString(s)
			return nil
		}

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		//把硬链接信息写入索引文件
		if sett.Has(stat.Ino) == true {
			s := "hard " + path + " " + sett[int(stat.Ino)] + "\n"
			outputWriter.WriteString(s)
			return nil
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)

		if !info.IsDir() {
			stat, _ := info.Sys().(*syscall.Stat_t)
			sett.Add(stat.Ino, path)
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})

	//刷新索引文件的缓冲
	outputWriter.Flush()

	//索引文件也打包进去
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {
		//给walk传入匿名函数
		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info) //获得文件信息
		header.Name = strings.TrimPrefix(path, srcDir+`/`)
		if header.Name != "restoreToPath.txt" {
			return nil
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		file, _ := os.Open(path)
		defer file.Close()
		io.Copy(writer, file)
		return nil
	})

	//删除路径文件
	outputError = os.Remove(filepath.Join(srcDir, "restoreToPath.txt"))
	if outputError != nil {
		fmt.Printf("An error occurred with writing to source path\n")
		return outputError
	}

	return nil
}

//使用aes库和base64库实现加密

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func RunBackup(srcPath, desPath, password, filename string) error {

	/*
		1. if desPath == "" then 备份到默认路径（./backup）else 备份到desPath
		2. if filename == "" then 备份文件名=默认文件名（原目录名称）else 备份文件名=filename
	*/

	var fname string
	if filename == "" {
		fname = filepath.Base(srcPath) + ".gz"
	} else {
		fname = filename + ".gz"
	}

	var des string
	if desPath == "" {
		defaultBackupDir := "./backup"
		des = filepath.Join(defaultBackupDir, fname)

		//判断是否在并创建./backup文件夹
		exist, err := PathExists(defaultBackupDir)
		if err != nil {
			fmt.Printf("An error occured while trying to access the default backup directory [%v]\n", err)
		}
		if !exist {
			fmt.Printf("The cache directory [%v] doesn't exist, creating...\n", defaultBackupDir)
			// 创建文件夹
			err := os.Mkdir(defaultBackupDir, os.ModePerm)
			if err != nil {
				fmt.Printf("Mkdir failed![%v]\n", err)
			} else {
				fmt.Printf("Mkdir successed!\n")
			}
		}
	} else {
		des = filepath.Join(desPath, fname)
	}

	err := Zip(srcPath, des)
	if err != nil {
		fmt.Println(err)
		return err
	}

	aeskey := []byte(password)
	num := 16 - len(aeskey)
	for i := 0; i < num; i++ {
		aeskey = append(aeskey, 0)
	}

	//读文件
	pass, err := ioutil.ReadFile(des)
	if err == nil {
		//fmt.Println("file content =", string(pass))
		//减少console中输出的过量内容
	} else {
		fmt.Println("read file error =", err)
	}
	xpass, err := AesEncrypt(pass, aeskey)
	if err != nil {
		fmt.Println(err)
		return err
	}
	pass64 := base64.StdEncoding.EncodeToString(xpass)

	//fmt.Printf("加密后:%v\n", pass64)

	//写入文件
	f, err := os.Create(des) //文件已存在，将会清空
	if err != nil {
		fmt.Println(err)
		return err
	}
	l, err := f.WriteString(string(pass64))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}
	fmt.Println(l, "写入文件成功")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
