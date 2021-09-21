package ui

import (
	"fmt"
	"github.com/wailsapp/wails"
)

type Backup struct {
	savePath string
	sourcePath string
	runtime  *wails.Runtime
	logger   *wails.CustomLogger
}

func NewBackup() (*Backup, error) {
	result := &Backup{}
	return result, nil
}

func (b *Backup) SelectSourceDir() {
	src := b.runtime.Dialog.SelectDirectory()
	b.logger.Info("Source directory selected" + src)
}

func (b *Backup) WailsInit(runtime *wails.Runtime) error {
	b.runtime = runtime
	b.logger = b.runtime.Log.New("Backup")
	b.logger.Info("I'm here")
	return nil
}