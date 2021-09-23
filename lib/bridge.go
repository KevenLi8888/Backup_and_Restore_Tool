package lib

import (
	"fmt"
	"github.com/wailsapp/wails"
	"io/ioutil"
)

type Backup struct {
	savePath   string
	sourcePath string
	runtime    *wails.Runtime
	logger     *wails.CustomLogger
}

func NewBackup() (*Backup, error) {
	result := &Backup{}
	return result, nil
}

func (b *Backup) SelectSourceDir() (string, error) {
	src := b.runtime.Dialog.SelectDirectory()
	if src != "" {
		b.logger.Info("Source directory selected:" + src)
		// 后端逻辑
		b.sourcePath = src
		// UI逻辑
		return src, nil
	} else {
		b.logger.Info("No directory selected!")
		err := fmt.Errorf("no directory selected")
		return "", err
	}
}

func (b *Backup) SelectDestDir() (string, error) {
	src := b.runtime.Dialog.SelectDirectory()
	if src != "" {
		b.logger.Info("Destination directory selected:" + src)
		// UI逻辑
		// 后端逻辑
		return src, nil
	} else {
		b.logger.Info("No directory selected!")
		err := fmt.Errorf("No directory selected!")
		return "", err
	}
}

func (b *Backup) SelectRestoreFile() (string, error) {
	src := b.runtime.Dialog.SelectFile()
	if src != "" {
		b.logger.Info("Restore file selected:" + src)
		// 后端逻辑
		b.sourcePath = src
		// UI逻辑
		return src, nil
	} else {
		b.logger.Info("No file selected!")
		err := fmt.Errorf("no file selected")
		return "", err
	}
}

func (b *Backup) PerformBackup(srcPath, password string) (string, error) {
	RunBackup(srcPath, password)
	//TODO: 后端错误处理
	return "Backup: Function call success!", nil
}

func (b *Backup) PerformRestore(srcPath, password string) (string, error) {
	RunRestore(srcPath, password)
	//TODO: 后端错误处理
	return "Restore: Function call success!", nil
}

//TODO: 相关逻辑待实现
func (b *Backup) LoadList() (string, error) {
	filePath := "./ui/backupHistory.json"
	b.logger.Infof("Loading list from: %s", filePath)
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		err = fmt.Errorf("Unable to open list: %s", filePath)
	}
	return string(bytes), err
}

func (b *Backup) WailsInit(runtime *wails.Runtime) error {
	b.runtime = runtime
	b.logger = b.runtime.Log.New("Backup")
	b.logger.Info("I'm here")
	return nil
}
