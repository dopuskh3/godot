package dot

import (
  "fmt"
  "log"
  "os"
  "path/filepath"
  "time"
)

func backupIfNeeded(file string) error {
  _, err := os.Lstat(file)
  if err == nil {
    suffix := time.Now().Format("200601020304")
    err = os.Rename(file, fmt.Sprintf("%s.bak-%s", file, suffix))
    if err != nil {
      return err
    }
  }
  return nil
}

func InstallFile(src string, dst string) error {
  // prepare target directory
  targetbasename := filepath.Dir(dst)
  err := os.MkdirAll(targetbasename, 0700)
  if err != nil {
    return err
  }
  err = backupIfNeeded(dst)
  if err != nil {
    return err
  }
  abssrc, err := filepath.Abs(src)
  if err != nil {
    return err
  }
  err = os.Symlink(abssrc, dst)
  if err != nil {
    return err
  }
  log.Printf("%s -> %s installed", src, dst)
  return nil
}

func InstallDotFiles(conf *DotConfig, target string) error {

  for src, dst := range conf.Files {
    relsrc, err := filepath.Rel(conf.Root, src)
    if err != nil {
      return err
    }
    compiledsrc := filepath.Join(conf.CompileDir, relsrc)
    fullTarget := filepath.Join(target, dst)
    InstallFile(compiledsrc, fullTarget)

  }
  return nil
}
