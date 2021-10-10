package lib

import (
	"fmt"

	"github.com/wailsapp/wails"
)

type Backup struct {
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
		b.sourcePath = src
		return src, nil
	} else {
		b.logger.Info("No directory selected!")
		err := fmt.Errorf("未选择目录！")
		return "", err
	}
}

func (b *Backup) SelectDestDir() (string, error) {
	src := b.runtime.Dialog.SelectDirectory()
	if src != "" {
		b.logger.Info("Destination directory selected:" + src)
		return src, nil
	} else {
		b.logger.Info("No directory selected!")
		err := fmt.Errorf("未选择目录！")
		return "", err
	}
}

func (b *Backup) SelectRestoreFile() (string, error) {
	src := b.runtime.Dialog.SelectFile()
	if src != "" {
		b.logger.Info("Restore file selected:" + src)
		b.sourcePath = src
		return src, nil
	} else {
		b.logger.Info("No file selected!")
		err := fmt.Errorf("未选择文件！")
		return "", err
	}
}

func (b *Backup) PerformBackup(srcPath, desPath, password, filename string) (string, error) {
	err := RunBackup(srcPath, desPath, password, filename)
	if err != nil {
		b.logger.Info(err.Error())
		return "", fmt.Errorf(err.Error())
	}
	return "✅ 备份完成", nil
}

func (b *Backup) PerformRestore(srcPath, password string) (string, error) {
	err := RunRestore(srcPath, password)
	if err != nil {
		b.logger.Info(err.Error())
		return "", fmt.Errorf(err.Error())
	}
	return "✅ 恢复完成", nil
}

func (b *Backup) WailsInit(runtime *wails.Runtime) error {
	b.runtime = runtime
	b.logger = b.runtime.Log.New("Backup")
	b.logger.Info("I'm here")
	return nil
}
