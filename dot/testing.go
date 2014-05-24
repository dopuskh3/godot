package dot

import (
  "io/ioutil"
  "os"
)

func createTestDir() (t string) {
  t, err := ioutil.TempDir("", "godot-test")
  if err != nil {
    panic(err)
  }
  return
}

func deleteTestDir(t string) {
  os.RemoveAll(t)
}
