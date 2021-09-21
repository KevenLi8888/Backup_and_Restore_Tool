package ui

import (
	"fmt"
	"github.com/wailsapp/wails"
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
		// UI逻辑
		// 后端逻辑
		return src, nil
	} else {
		b.logger.Info("No directory selected!")
		err := fmt.Errorf("No directory selected!")
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

func (b *Backup) WailsInit(runtime *wails.Runtime) error {
	b.runtime = runtime
	b.logger = b.runtime.Log.New("Backup")
	b.logger.Info("I'm here")
	return nil
}
