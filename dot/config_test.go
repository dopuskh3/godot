package dot

import (
  "github.com/bmizerany/assert"
  "io/ioutil"
  "os"
  "path/filepath"
  "testing"
)

var sample_config = `
config:
  foo: bar

files:
  bar: bazz
`

func TestLoadConfig(t *testing.T) {
  conf, _ := LoadConfig([]byte(sample_config))
  t.Logf("%#v", conf)
  assert.Equal(t, conf.Config["foo"], "bar")
}

func TestLoadConfigFromNonExistingFile(t *testing.T) {
  _, err := LoadConfigFromFile("non-existent-file")
  assert.NotEqual(t, err, nil)
}

func TestLoadInvalidConfig(t *testing.T) {
  invalidConfig := `
{ 
  "asdaidsasd
asdasdasda--d-adasd
  `
  _, err := LoadConfig([]byte(invalidConfig))
  assert.NotEqual(t, err, nil)
}

func TestLoadConfigFromFile(t *testing.T) {
  tdir, _ := ioutil.TempDir("", "godot-test")
  defer os.RemoveAll(tdir)
  confpath := filepath.Join(tdir, "conf.yml")
  ioutil.WriteFile(filepath.Join(tdir, "conf.yml"), []byte(sample_config), 0777)
  conf, _ := LoadConfigFromFile(confpath)
  assert.Equal(t, conf.Config["foo"], "bar")
}
