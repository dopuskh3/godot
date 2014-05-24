package dot

import (
  "errors"
  "gopkg.in/yaml.v1"
  "io/ioutil"
  "os"
  "path/filepath"
)

type DotConfig struct {
  Config map[string]string
  Files  map[string]string
  Root   string
}

func LoadConfigFromFile(path string) (*DotConfig, error) {
  fi, err := ioutil.ReadFile(os.ExpandEnv(path))
  if err != nil {
    return nil, err
  }
  conf, err := LoadConfig(fi)
  if err != nil {
    return nil, err
  }
  absPath, err := filepath.Abs(path)
  if err != nil {
    return nil, err
  }
  conf.Root = filepath.Dir(absPath)
  err = validateConfig(conf)
  if err != nil {
    return nil, err
  }
  return conf, err
}

func validateConfig(config *DotConfig) error {
  newFiles := make(map[string]string)
  for k, v := range config.Files {
    if filepath.IsAbs(k) {
      return errors.New("All file path must be absolute paths.")
    }
    absFile := filepath.Join(config.Root, k)
    _, err := os.Stat(absFile)
    if err != nil {
      return err
    }
    newFiles[absFile] = v
  }
  config.Files = newFiles
  return nil
}

func LoadConfig(conf []byte) (*DotConfig, error) {
  config := &DotConfig{}

  err := yaml.Unmarshal(conf, config)
  if err != nil {
    return nil, err
  }
  return config, nil
}
