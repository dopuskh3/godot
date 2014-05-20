package dot

import (
  "fmt"
  "log"
  "os"
  "path/filepath"
)

func InstallFile(src, dst string) error {
  // prepare target directory
  targetbasename := filepath.Dir(dst)
  err := os.MkdirAll(targetbasename, 0700)
  if err != nil {
    return err
  }

  _, err = os.Stat(src)
  if err != nil {
    return err
  }
  fi, err := os.Lstat(dst)
  if err == nil {
    if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
      log.Printf("Removing symlink %s", dst)
      os.Remove(dst)
    } else {
      err = os.Rename(dst, fmt.Sprintf("%s.bak", dst))
    }
  } else {
    err = nil
  }
  src, err = filepath.Abs(src)
  if err != nil {
    return err
  }
  log.Printf("Installing %s -> %s", dst, src)
  err = os.Symlink(src, dst)
  return nil
}

func InstallFiles(srcroot, targetroot string, files map[string]string) error {

  for k, v := range files {
    sourcefile := filepath.Join(srcroot, k)
    targetfile := filepath.Join(targetroot, v)
    err := InstallFile(sourcefile, targetfile)
    if err != nil {
      return err
    }
  }
  return nil
}
