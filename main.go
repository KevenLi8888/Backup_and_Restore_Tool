package main

import (
  _ "embed"
  "github.com/wailsapp/wails"
  "gui1/ui"
  "log"
)

func basic() string {
  return "Hello World!"
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

  myBackup, err := ui.NewBackup()
  if err != nil {
    log.Fatal(err)
  }

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "gui1",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })

  app.Bind(myBackup)
  app.Bind(basic)
  app.Run()
}
