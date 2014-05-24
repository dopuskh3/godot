package main

import (
  "github.com/dopuskh3/godot/dot"
  "log"
  "os"
)

var DEFAULT_CONFIG = "godot.yml"
var COMPILE_OUTPUT_DIR = ".godot"
var INSTALL_DIRECTORY = os.ExpandEnv("$HOME")

func main() {
  conf, err := dot.LoadConfigFromFile("godot.yml")
  if err != nil {
    log.Fatalf("Cannot load config file %s: %s", DEFAULT_CONFIG, err)
  } else {
    log.Printf("Loaded %s", DEFAULT_CONFIG)
  }
  log.Printf("%#v", conf)
  log.Print("config root: ", conf.Root)
  err = dot.TemplatizeAll(conf.Root, "out", conf.Config)
  if err != nil {
    log.Fatal(err)
  }
  /*
     done, err := dot.WatchDir("./", "out", conf.Files)
     if err != nil {
       log.Fatal(err)
       done <- true
     }
     <-done
  */

}
