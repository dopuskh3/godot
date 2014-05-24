package dot

import (
  "bytes"
  "io"
  "log"
  "os"
  "path/filepath"
  "text/template"
)

func Templatize(inpath string, vars map[string]string) (*bytes.Buffer, error) {
  var buffer bytes.Buffer
  tpl, err := template.ParseFiles(inpath)
  if err != nil {
    log.Printf("Warning: Can't parce %s template. Installing it anyway.")
    f, err := os.Open(inpath)
    defer f.Close()
    if err != nil {
      return nil, err
    }
    io.Copy(&buffer, f)
    return &buffer, nil
  }
  err = tpl.Execute(&buffer, vars)
  if err != nil {
    return nil, err
  }
  return &buffer, nil
}

func TemplatizeAll(root string, output string, vars map[string]string) error {
  x := func(path string, f os.FileInfo, err error) error {
    if f.IsDir() && f.Name() != "." && f.Name()[0:1] == "." {
      log.Printf("Skipping %s", path)
      return filepath.SkipDir
    }
    if !f.IsDir() && f.Name()[0:1] != "." {
      outdir := filepath.Join(output, filepath.Dir(path))
      outfile := filepath.Join(outdir, f.Name())
      log.Printf("Creating %s", outdir)
      err = os.MkdirAll(outdir, 0770)
      if err != nil {
        return err
      }
      buff, err := Templatize(path, vars)
      if err != nil {
        return err
      }
      err = InstallFile(buff, outfile)
      if err != nil {
        return err
      }
    }
    return nil
  }
  err := filepath.Walk(root, x)
  return err
}
