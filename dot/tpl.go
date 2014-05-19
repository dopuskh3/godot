package dot

import (
  "errors"
  "fmt"
  "log"
  "os"
  "path/filepath"
  "text/template"
)

var TPL_EXT = ".gotpl"

func Templatize(path string, vars map[string]string) error {
  if filepath.Ext(path) != TPL_EXT {
    return errors.New(fmt.Sprintf("Invalid template file %s. Must have the %s file ext.", path, TPL_EXT))
  }
  tpl, err := template.ParseFiles(path)
  if err != nil {
    return err
  }
  outputfile := path[0 : len(path)-len(TPL_EXT)]
  fd, err := os.Create(outputfile)
  defer fd.Close()
  err = tpl.Execute(fd, vars)
  if err != nil {
    return err
  }
  log.Printf("Generated %s", outputfile)
  return nil
}

func TemplatizeAll(root string, vars map[string]string) error {
  x := func(path string, f os.FileInfo, err error) error {
    if !f.IsDir() {
      if filepath.Ext(path) == TPL_EXT {
        err = Templatize(path, vars)
        return err
      }
    }
    return nil
  }
  err := filepath.Walk(root, x)
  return err
}
