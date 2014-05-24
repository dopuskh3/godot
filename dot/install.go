package dot

import (
  "bytes"
  "fmt"
  "log"
  "os"
  "path/filepath"
  "time"
)

func backupIfNeeded(file string) error {
  log.Printf("Will intall %s", file)
  _, err := os.Lstat(file)
  if err == nil {
    suffix := time.Now().Format("200601020304")
    err = os.Rename(file, fmt.Sprintf("%s.bak-", file, suffix))
    if err != nil {
      return err
    }
  }
  return nil
}

func InstallFile(buffer *bytes.Buffer, dst string) error {
  // prepare target directory
  targetbasename := filepath.Dir(dst)
  err := os.MkdirAll(targetbasename, 0700)
  if err != nil {
    return err
  }
  log.Printf("Will intall %s", dst)
  _, err = os.Lstat(dst)
  if err == nil {
    err = os.Rename(dst, fmt.Sprintf("%s.bak", dst))
    if err != nil {
      return err
    }
  } else {
    err = nil
  }
  log.Printf("Creating %s", dst)
  fd, err := os.Create(dst)
  log.Print("Created")
  if err != nil {
    return err
  }
  defer fd.Close()
  _, err = buffer.WriteTo(fd)
  if err != nil {
    return err
  }
  log.Printf("%s installed", dst)
  return nil
}
