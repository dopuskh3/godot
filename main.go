package main

import (
  "github.com/dopuskh3/godot/dot"
  "log"
)

var DEFAULT_CONFIG = "godot.yml"

func main() {
  conf, err := dot.LoadConfigFromFile("godot.yml")
  if err != nil {
    log.Fatal("Cannot load config file %s", DEFAULT_CONFIG)
  } else {
    log.Printf("Loaded %s", DEFAULT_CONFIG)
  }
  err = dot.TemplatizeAll("./", conf.Config)
  if err != nil {
    log.Fatal(err)
  }
}
