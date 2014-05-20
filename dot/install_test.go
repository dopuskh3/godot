package dot

import (
  "github.com/bmizerany/assert"
  "io/ioutil"
  "os"
  "path/filepath"
  "testing"
)

func TestInstallFile(t *testing.T) {
  tdir, _ := ioutil.TempDir("", "godot-test")
  defer os.RemoveAll(tdir)
  _ = os.MkdirAll(filepath.Join(tdir, "src"), 0777)
  _ = os.MkdirAll(filepath.Join(tdir, "dst"), 0777)
  ioutil.WriteFile(filepath.Join(tdir, "src/config_file"), []byte("target file content"), 0777)
  err := InstallFile(filepath.Join(tdir, "src/config_file"), filepath.Join(tdir, "dst/foo/bar/bazz/config_file"))
  assert.Equal(t, err, nil)
  target, err := os.Readlink(filepath.Join(tdir, "dst/foo/bar/bazz/config_file"))
  assert.Equal(t, filepath.IsAbs(target), true)
  assert.Equal(t, filepath.Join(tdir, "src/config_file"), target)
}

func TestInstallNonExistentFile(t *testing.T) {
  err := InstallFile("not-exists", "./")
  assert.NotEqual(t, err, nil)
}

func TestBackupExistingTargetFile(t *testing.T) {
  tdir, _ := ioutil.TempDir("", "godot-test")
  defer os.RemoveAll(tdir)
  _ = os.MkdirAll(filepath.Join(tdir, "src"), 0777)
  _ = os.MkdirAll(filepath.Join(tdir, "dst"), 0777)

  ioutil.WriteFile(filepath.Join(tdir, "src/config_file"), []byte("target file content"), 0777)
  ioutil.WriteFile(filepath.Join(tdir, "dst/config_file"), []byte("target file content"), 0777)
  err := InstallFile(filepath.Join(tdir, "src/config_file"), filepath.Join(tdir, "dst/config_file"))
  assert.Equal(t, err, nil)
  _, err = os.Stat(filepath.Join(tdir, "dst/config_file.bak"))
  assert.Equal(t, err, nil)
}

func TestUpdateExistingSymlink(t *testing.T) {
  tdir, _ := ioutil.TempDir("", "godot-test")
  defer os.RemoveAll(tdir)
  _ = os.MkdirAll(filepath.Join(tdir, "src"), 0777)
  _ = os.MkdirAll(filepath.Join(tdir, "dst/"), 0777)
  ioutil.WriteFile(filepath.Join(tdir, "src/config_file1"), []byte("target file content"), 0777)
  ioutil.WriteFile(filepath.Join(tdir, "src/config_file2"), []byte("target file content"), 0777)
  InstallFile(filepath.Join(tdir, "src/config_file1"), filepath.Join(tdir, "dst/config_file"))
  InstallFile(filepath.Join(tdir, "src/config_file2"), filepath.Join(tdir, "dst/config_file"))
  target, _ := os.Readlink(filepath.Join(tdir, "dst/config_file"))
  srcpath2, _ := filepath.Abs(filepath.Join(tdir, "src/config_file2"))
  assert.Equal(t, srcpath2, target)
}

func TestInstallFiles(t *testing.T) {
  tdir, _ := ioutil.TempDir("", "godot-test")
  _ = os.MkdirAll(filepath.Join(tdir, "src"), 0777)
  _ = os.MkdirAll(filepath.Join(tdir, "dst"), 0777)
  conf1 := filepath.Join(tdir, "src/conf1")
  conf2 := filepath.Join(tdir, "src/conf2")
  ioutil.WriteFile(conf1, []byte("target file content"), 0777)
  ioutil.WriteFile(conf2, []byte("target file content"), 0777)
  files := map[string]string{
    "conf1": "dst/conf1",
    "conf2": "dst/conf2",
  }
  err := InstallFiles(filepath.Join(tdir, "src"), filepath.Join(tdir, "dst"), files)
  assert.Equal(t, err, nil)
}

func TestInstallFilesFail(t *testing.T) {
  tdir, _ := ioutil.TempDir("", "godot-test")
  _ = os.MkdirAll(filepath.Join(tdir, "src"), 0777)
  _ = os.MkdirAll(filepath.Join(tdir, "dst"), 0777)
  conf1 := filepath.Join(tdir, "src/conf1")
  ioutil.WriteFile(conf1, []byte("target file content"), 0777)
  files := map[string]string{
    "conf1":      "dst/conf1",
    "donotexist": "dst/donotexisttarget",
  }
  err := InstallFiles(filepath.Join(tdir, "src"), filepath.Join(tdir, "dst"), files)
  assert.NotEqual(t, err, nil)

}
