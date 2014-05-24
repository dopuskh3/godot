package dot

import (
  "github.com/bmizerany/assert"
  "io/ioutil"
  "path/filepath"
  "testing"
)

var sample_config = `
config:
  foo: bar

files:
  bar: bazz
`
var sample_invalid_config = ` 
config:
  foo: bar
files:
  /foo/bar: foobar
`

func TestLoadConfig(t *testing.T) {
  td := createTestDir()
  defer deleteTestDir(td)
  inFile := filepath.Join(td, "bar")
  confFile := filepath.Join(td, "conf")
  ioutil.WriteFile(confFile, []byte(sample_config), 0700)
  ioutil.WriteFile(inFile, []byte("foobar"), 0700)

  conf, err := LoadConfigFromFile(confFile)
  assert.NotEqual(t, conf, nil)
  assert.Equal(t, err, nil)
}

func TestLoadConfigWithAbsolutePathFail(t *testing.T) {
  td := createTestDir()
  defer deleteTestDir(td)
  confFile := filepath.Join(td, "conf")
  ioutil.WriteFile(confFile, []byte(sample_invalid_config), 0700)

  conf, err := LoadConfigFromFile(confFile)
  assert.Equal(t, conf, (*DotConfig)(nil))
  assert.NotEqual(t, err, nil)
}

func TestLoadConfigFailWhenFileDoNotExists(t *testing.T) {
  td := createTestDir()
  defer deleteTestDir(td)
  confFile := filepath.Join(td, "conf")
  ioutil.WriteFile(confFile, []byte(sample_config), 0700)
  conf, err := LoadConfigFromFile(confFile)
  assert.Equal(t, conf, (*DotConfig)(nil))
  assert.NotEqual(t, err, nil)
}

func TestLoadInvalidYaml(t *testing.T) {
  td := createTestDir()
  defer deleteTestDir(td)

  confFile := filepath.Join(td, "conf")
  confYaml := `
  foo: @ bar
  `
  ioutil.WriteFile(confFile, []byte(confYaml), 0700)
  conf, err := LoadConfigFromFile(confFile)
  assert.Equal(t, conf, (*DotConfig)(nil))
  assert.NotEqual(t, err, nil)
}

func TestLoadConfigFailWhenConfigNotExists(t *testing.T) {
  _, err := LoadConfigFromFile("not-exists")
  assert.NotEqual(t, err, nil)
}
