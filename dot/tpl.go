package dot

import (
  "log"
  "os"
  "path/filepath"
  "text/template"
)

func Templatize(inpath string, outpath string, vars map[string]string) error {
  tpl, err := template.ParseFiles(inpath)
  if err != nil {
    log.Printf("Warning: Can't parse %s template..", inpath)
    _, err = os.Stat(outpath)
    if err == nil {
      os.Remove(outpath)
    }
    err := os.Link(inpath, outpath)
    if err != nil {
      return err
    }
    return nil
  }
  f, err := os.Create(outpath)
  defer f.Close()
  err = tpl.Execute(f, vars)
  if err != nil {
    return err
  }
  return nil
}

func TemplatizeAll(input string, output string, vars map[string]string) error {
  install := func(path string, f os.FileInfo, err error) error {
    if f.IsDir() && f.Name() != "." && f.Name()[0:1] == "." {
      log.Printf("Skipping %s", path)
      return filepath.SkipDir
    }
    if !f.IsDir() && f.Name()[0:1] != "." {
      outdir := filepath.Join(output, filepath.Dir(path))
      outfile := filepath.Join(outdir, f.Name())
      err = os.MkdirAll(outdir, 0770)
      if err != nil {
        return err
      }
      err := Templatize(path, outfile, vars)
      log.Printf("Compiled %s", outfile)
      if err != nil {
        return err
      }
    }
    return nil
  }
  err := filepath.Walk(input, install)
  return err
}
