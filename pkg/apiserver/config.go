package apiserver

import (
    "fmt"
    "github.com/zryfish/kubespheretest/pkg/simple/devops"
    "github.com/zryfish/kubespheretest/pkg/simple/kubernetes"
    "github.com/zryfish/kubespheretest/pkg/simple/mysql"
    "github.com/zryfish/kubespheretest/pkg/simple/redis"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

var config *Config

// Configuration for apiserver
type Config struct {
    MySQL      *mysql.MySQLOptions           `json:"mysql,omitempty" yaml:",omitempty"`
    Devops     *devops.DevopsOptions         `json:"devops,omitempty" yaml:",omitempty"`
    Redis      *redis.RedisOptions           `json:"redis,omitempty" yaml:",omitempty"`
    Kubernetes *kubernetes.KubernetesOptions `json:"kubernetes,omitempty" yaml:",omitempty"`
}

func NewConfig() *Config {
    return &Config{}
}

func Get() *Config {
    return config
}

func Set(c *Config) {
    config = c
}

func Marshal(c *Config) (string, error) {
    yamlBytes, err := yaml.Marshal(&c)
    if err != nil {
        return "", fmt.Errorf("failed to produce yamk. error=%v", err)
    }
    yamlString := string(yamlBytes)
    return yamlString, nil
}

func Unmarshal(yamlString string) (*Config, error) {
    conf := NewConfig()
    err := yaml.Unmarshal([]byte(yamlString), &conf)
    if err != nil {
        return nil, fmt.Errorf("failed to parse yaml data. error=%v", err)
    }
    return conf, nil
}

func LoadFromFile(filename string) (*Config, error) {
    fileContent, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to load config file [%v]. error=%v", filename, err)
    }
    return Unmarshal(string(fileContent))
}

func SaveToFile(filename string, conf *Config) error {
    fileContent, err := Marshal(conf)
    if err != nil {
        return fmt.Errorf("failed to save config file [%v]. error=%v", filename, err)
    }

    err = ioutil.WriteFile(filename, []byte(fileContent), 0640)
    return err
}
