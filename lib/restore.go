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
	"syscall"
)

func Unzip(srcPath string, desPath string) error {
	//desPath是中转

	var errorMessage error
	err := os.MkdirAll(desPath, 0755)
	if err != nil {
		fmt.Printf("An error occurred while making directory\n")
		return err
	}

	// file read
	//打开并读取压缩文件中的内容/
	fr, err := zip.OpenReader(srcPath)
	var iserr int
	if err != nil {
		fmt.Println("无法压缩,密码不对")
		errorMessage = fmt.Errorf("密码错误！")
		//fmt.Println("从" + filepath.Join("../mnt", filepath.Base(srcPath)) + "中恢复到" + srcPath)
		iserr = 1
	}

	if iserr == 1 {
		err := os.Rename(filepath.Join("./cache1", filepath.Base(srcPath)), srcPath)
		if err != nil {
			fmt.Printf("An error occurred writing to source path\n")
			return err
		}
		fmt.Println("从" + filepath.Join("./cache1", filepath.Base(srcPath)) + "中恢复到" + srcPath)
		os.RemoveAll("./cache")
		os.RemoveAll("./cache1")

		//panic会直接引起程序崩溃，改成Go的标准错误处理逻辑以配合前端
		//panic(err)
		return errorMessage
	}

	//调整了调用位置，密码错误时无需调用fr.Close()，否则导致程序出错
	defer fr.Close()

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

		//在控制台输出文件的文件名及路径
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

	var restoreToPath string //目标地址
	file, err := os.OpenFile(filepath.Join(desPath, "restoreToPath.txt"), os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return err
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	n := 0
	for {
		line, _ := buf.ReadString('\n') //读一行出来
		line = strings.TrimSpace(line)  //去掉首尾空格
		if n == 0 {
			n = 1
			restoreToPath = line
			//移动文件夹
			_, err1 := os.Stat(filepath.Dir(restoreToPath))
			if err1 != nil {
				os.MkdirAll(filepath.Dir(restoreToPath), 0755)
			}
			err := os.Rename(desPath, restoreToPath)
			if err != nil {
				fmt.Printf("An error occurred while moving files\n")
				os.RemoveAll(desPath)
				//return err
			}

			fmt.Println("从" + desPath + "移到了" + restoreToPath)
		} else {
			arr := strings.Fields(line) //按照空格分隔
			if len(arr) < 2 {
				break
			}
			if arr[0] == "soft" {
				//创建软链接
				//fmt.Println(arr[1]+arr[2])
				os.Symlink(arr[2], arr[1])
			} else if arr[0] == "hard" {
				os.Link(arr[2], arr[1])
			} else {
				syscall.Mkfifo(arr[1], 0666)
			}
		}
	}

	//删除路径文件
	err = os.Remove(filepath.Join(restoreToPath, "restoreToPath.txt"))
	if err != nil {
		fmt.Printf("An error occurred while removing files\n")
		return err
	}

	err = os.Rename(filepath.Join("./cache1", filepath.Base(srcPath)), srcPath)
	if err != nil {
		fmt.Printf("An error occurred while moving files to source path\n")
		return err
	}
	fmt.Println("从" + filepath.Join("./cache1", filepath.Base(srcPath)) + "中移到" + srcPath)

	return nil
}

//使用aes库和base64库实现解密
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

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

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func RunRestore(srcPath, password string) error {
	aeskey := []byte(password)
	num := 16 - len(aeskey)
	for i := 0; i < num; i++ {
		aeskey = append(aeskey, 0)
	}
	pass64, err := ioutil.ReadFile(srcPath)
	if err == nil {
		//fmt.Println("file content =", string(pass64))
		//减少console中输出的过量内容
	} else {
		fmt.Println("read file error =", err)
	}

	//提前备份一份
	//修复缓存位置问题：添加不存在则先创建逻辑

	cacheDir := "./cache1"
	exist, err := PathExists(cacheDir)
	if err != nil {
		fmt.Printf("An error occured during while trying to access the cache directory [%v]\n", err)
	}

	if !exist {
		fmt.Printf("The cache directory [%v] doesn't exist, creating...\n", cacheDir)
		// 创建文件夹
		err := os.Mkdir(cacheDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("Mkdir successed!\n")
		}
	}

	_, err = CopyFile(filepath.Join(cacheDir, filepath.Base(srcPath)), srcPath)
	if err != nil {
		fmt.Printf("An error occurred while copying files\n")
		return err
	}

	//解密
	bytesPass, err := base64.StdEncoding.DecodeString(string(pass64))
	if err != nil {
		fmt.Println(err)
		return err
	}
	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//fmt.Printf("解密后:%s\n", tpass)
	//写入文件
	f, err := os.Create(srcPath) //文件已存在，将会清空
	if err != nil {
		fmt.Println(err)
		return err
	}
	l, err := f.WriteString(string(tpass))
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
	fmt.Println("开始解压")
	err = Unzip(srcPath, "./cache") //这个是中转站 先解压到这里
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = os.RemoveAll(cacheDir)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
