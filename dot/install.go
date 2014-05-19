package dot

import (
  "fmt"
  "log"
  "os"
  "path/filepath"
)

func InstallFiles(srcroot, targetroot string, files map[string]string) error {

  for k, v := range files {
    sourcefile := filepath.Join(srcroot, k)
    targetfile := filepath.Join(targetroot, v)
    _, err := os.Stat(sourcefile)
    if err != nil {
      return err
    }
    fi, err := os.Lstat(targetfile)
    if err == nil {
      if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
        log.Printf("Removing symlink %s", targetfile)
        os.Remove(targetfile)
      } else {
        err = os.Rename(targetfile, fmt.Sprintf("%s.bak", targetfile))
      }
    } else {
      err = nil
    }
    src, err := filepath.Abs(sourcefile)
    if err != nil {
      return err
    }
    log.Printf("Installing %s -> %s", targetfile, src)
    err = os.Symlink(src, targetfile)
    if err != nil {
      return err
    }
  }
  return nil
}
