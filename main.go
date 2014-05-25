package main

import (
  "github.com/dopuskh3/godot/dot"
  "log"
)

var DEFAULT_CONFIG = "godot.yml"

func Update() error {
  conf, err := dot.LoadConfigFromFile(DEFAULT_CONFIG)
  if err != nil {
    log.Fatalf("Cannot load config file %s: %s", DEFAULT_CONFIG, err)
  }
  err = dot.TemplatizeAll("./", conf.CompileDir, conf.Config)
  if err != nil {
    log.Fatalf("Cannot compile templates: %s", err)
    return err
  }
  err = dot.InstallDotFiles(conf, ".installdir")
  if err != nil {
    log.Fatalf("Cannot install: %s", err)
  }
  return nil
}

func main() {
  Update()
}
