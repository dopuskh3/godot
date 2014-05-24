package dot

import (
  "bytes"
  "fmt"
  "github.com/bmizerany/assert"
  "io/ioutil"
  "os"
  "path/filepath"
  "testing"
)

func TestInstallFile(t *testing.T) {
  td := createTestDir()
  defer deleteTestDir(td)
  b := bytes.NewBuffer([]byte("foobar"))
  target := filepath.Join(td, "baz")
  err := InstallFile(b, target)
  assert.Equal(t, err, nil)
  _, err = os.Stat(target)
  assert.Equal(t, err, nil)
  content, err := ioutil.ReadFile(target)
  assert.Equal(t, string(content), "foobar")
}

func TestBackupIfNeeded(t *testing.T) {
  td := createTestDir()
  defer deleteTestDir(td)
  err := backupIfNeeded(filepath.Join(td, "foo"))
  assert.Equal(t, err, nil)
}

func TestBackupWhenFileExists(t *testing.T) {
  td := createTestDir()
  defer deleteTestDir(td)
  fileToBackup := filepath.Join(td, "foo")
  ioutil.WriteFile(fileToBackup, []byte("foobar"), 0700)
  err := backupIfNeeded(fileToBackup)
  assert.Equal(t, err, nil)
  files, err := filepath.Glob(filepath.Join(td, "foo.bak-*"))
  assert.Equal(t, len(files), 1)
}

func TestInstallFileBackupFile(t *testing.T) {
  td := createTestDir()
  defer deleteTestDir(td)

  b := bytes.NewBuffer([]byte("foobar"))
  target := filepath.Join(td, "baz")
  ioutil.WriteFile(target, []byte("already exist"), 0700)
  err := InstallFile(b, target)
  assert.Equal(t, err, nil)
  _, err = os.Stat(target)
  assert.Equal(t, err, nil)
  content, err := ioutil.ReadFile(target)
  assert.Equal(t, string(content), "foobar")

  _, err = os.Stat(fmt.Sprintf("%s.bak", target))
  assert.Equal(t, err, nil)
}
