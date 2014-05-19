package dot

import (
  "gopkg.in/yaml.v1"
  "io/ioutil"
  "os"
)

type DotConfig struct {
  Config map[string]string
  Files  map[string]string
}

func LoadConfigFromFile(path string) (*DotConfig, error) {
  fi, err := ioutil.ReadFile(os.ExpandEnv(path))
  if err != nil {
    return nil, err
  }
  conf, err := LoadConfig(fi)
  return conf, err
}

func LoadConfig(conf []byte) (*DotConfig, error) {
  config := &DotConfig{}

  err := yaml.Unmarshal(conf, config)
  if err != nil {
    return nil, err
  }
  return config, nil
}
