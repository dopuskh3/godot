package dot

import (
  "log"
  "os"
  "path/filepath"
  "text/template"
)

var COMPILE_DIR = ".godot"

func CreateDir(path string) error {
  _, err := os.Stat(path)
  if err != nil {
    log.Printf("%s already exists. Cleaning.", path)
    err = os.RemoveAll(path)
    if err != nil {
      return err
    }
  }
  log.Printf("Creating %s", path)
  return os.MkdirAll(path, 0770)
}

func Templatize(inpath string, outpath string, vars map[string]string) error {
  tpl, err := template.ParseFiles(inpath)
  if err != nil {
    return err
  }
  fd, err := os.Create(outpath)
  defer fd.Close()
  err = tpl.Execute(fd, vars)
  if err != nil {
    return err
  }
  log.Printf("Generated %s", outpath)
  return nil
}

func TemplatizeAll(root string, vars map[string]string) error {
  target := filepath.Join(root, COMPILE_DIR)

  x := func(path string, f os.FileInfo, err error) error {
    if f.IsDir() && f.Name() != "." && f.Name()[0:1] == "." {
      log.Printf("Skipping %s", path)
      return filepath.SkipDir
    }
    if !f.IsDir() && f.Name()[0:1] != "." {
      outdir := filepath.Join(target, filepath.Dir(path))
      outfile := filepath.Join(outdir, f.Name())
      log.Printf("Creating %s", outdir)
      err = os.MkdirAll(outdir, 0770)
      if err != nil {
        return err
      }
      err = Templatize(path, outfile, vars)
    }
    return nil
  }
  CreateDir(target)
  err := filepath.Walk(root, x)
  return err
}
