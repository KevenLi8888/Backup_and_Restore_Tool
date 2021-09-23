package lib

import (
	"archive/zip"
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(srcPath string, desPath string) { //desPath是中转

	os.MkdirAll(".", 0755) //中转 TODO: path "."是啥意思
	// file read
	//打开并读取压缩文件中的内容/
	fr, err := zip.OpenReader(srcPath)
	var iserr int
	if err != nil {
		fmt.Println("无法压缩,密码不对")
		//os.Rename(filepath.Join("../mnt", filepath.Base(srcPath)), srcPath)
		//fmt.Println("从" + filepath.Join("../mnt", filepath.Base(srcPath)) + "中恢复到" + srcPath)
		iserr = 1
		//panic(err)
	}
	defer fr.Close()
	if iserr == 1 {
		os.Rename(filepath.Join("../mnt", filepath.Base(srcPath)), srcPath)
		fmt.Println("从" + filepath.Join("../mnt", filepath.Base(srcPath)) + "中恢复到" + srcPath)
		panic(err)
	}
	//r.reader.file 是一个集合，里面包括了压缩包里面的所有文件
	for _, file := range fr.Reader.File {
		//判断文件该目录文件是否为文件夹
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(filepath.Join(desPath, file.Name), 0755)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		//为文件时，打开文件
		r, err := file.Open()

		//文件为空的时候，打印错误
		if err != nil {
			fmt.Println(err)
			continue
		}
		//这里在控制台输出文件的文件名及路径
		fmt.Println("unzip: ", file.Name)

		//在对应的目录中创建相同的文件
		NewFile, err := os.Create(filepath.Join(desPath, file.Name))
		if err != nil {
			fmt.Println(err)
			continue
		}
		//将内容复制
		io.Copy(NewFile, r)
		//关闭文件
		NewFile.Close()
		r.Close()
	}

	var pathpathpath string //目标地址
	file, err := os.OpenFile(filepath.Join(desPath, "pathpathpath.txt"), os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	n := 0
	for {
		line, _ := buf.ReadString('\n') //读一行出来
		line = strings.TrimSpace(line)  //去掉首尾空格
		if n == 0 {
			n = 1
			pathpathpath = line
			//移动文件夹
			os.Rename(desPath, pathpathpath)
			//listDir(desPath, pathpathpath, 0)
			fmt.Println("从" + desPath + "移到了" + pathpathpath)
		} else {
			arr := strings.Fields(line) //按照空格分隔
			if len(arr) < 3 {
				break
			}
			if arr[0] == "soft" {
				//创建软链接
				//fmt.Println(arr[1]+arr[2])
				os.Symlink(arr[2], arr[1])
			} else if arr[0] == "hard" {
				os.Link(arr[2], arr[1])
			}
		}
	}

	//删除路径文件
	os.Remove(filepath.Join(pathpathpath, "pathpathpath.txt"))

	os.Rename(filepath.Join("../mnt", filepath.Base(srcPath)), srcPath)
	fmt.Println("从" + filepath.Join("../mnt", filepath.Base(srcPath)) + "中移到" + srcPath)
}

//使用aes库和base64库实现解密

//func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
//	padding := blockSize - len(ciphertext)%blockSize
//	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(ciphertext, padtext...)
//}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//func AesEncrypt(origData, key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	blockSize := block.BlockSize()
//	origData = PKCS5Padding(origData, blockSize)
//	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
//	crypted := make([]byte, len(origData))
//	blockMode.CryptBlocks(crypted, origData)
//	return crypted, nil
//}

//解密

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

//拷贝文件

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func RunRestore(srcPath, password string) {
	//list := os.Args
	//if len(list) != 2 {
	//	fmt.Println("参数错误")
	//	return
	//}

	//fmt.Println("请输入密码：")
	//var aeskey []byte
	//fmt.Scanln(&aeskey)

	aeskey := []byte(password)
	num := 16 - len(aeskey)
	for i := 0; i < num; i++ {
		aeskey = append(aeskey, 0)
	}
	pass64, err := ioutil.ReadFile(srcPath)
	if err == nil {
		fmt.Println("file content =", string(pass64))
	} else {
		fmt.Println("read file error =", err)
	}
	//提前备份一份
	CopyFile(filepath.Join("../mnt", filepath.Base(srcPath)), srcPath)
	//return
	//解密
	bytesPass, err := base64.StdEncoding.DecodeString(string(pass64))
	if err != nil {
		fmt.Println(err)
		return
	}
	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("解密后:%s\n", tpass)
	//写入文件
	f, err := os.Create(srcPath) //文件已存在，将会清空
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(string(tpass))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "写入文件成功")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("开始解压")
	Unzip(srcPath, "./mnt") //这个是中转站 先解压到这里
}

/*
func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("参数错误")
		return
	}
	CopyFile(filepath.Join("../mnt", filepath.Base(list[1])), list[1])
}
*/
