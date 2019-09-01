package options

import (
    genericoptions "github.com/zryfish/kubespheretest/pkg/apiserver/server/options"
    "github.com/zryfish/kubespheretest/pkg/simple/devops"
    "github.com/zryfish/kubespheretest/pkg/simple/kubernetes"
    "github.com/zryfish/kubespheretest/pkg/simple/mysql"
    "github.com/zryfish/kubespheretest/pkg/simple/redis"
    cliflag "k8s.io/component-base/cli/flag"
)

type ServerRunOptions struct {
    GenericServerRunOptions *genericoptions.ServerRunOptions
    Redis                   *redis.RedisOptions
    MySQL                   *mysql.MySQLOptions
    Devops                  *devops.DevopsOptions
    Kubernetes *kubernetes.KubernetesOptions
}

func NewServerRunOptions() *ServerRunOptions {
    s := ServerRunOptions{
        GenericServerRunOptions: genericoptions.NewServerRunOptions(),
        Redis:                   redis.NewRedisOptions(),
        MySQL:                   mysql.NewMySQLOptions(),
        Devops:                  devops.NewDevopsOptions(),
        Kubernetes: kubernetes.NewKubernetesOptions(),
    }

    return &s
}

func (s *ServerRunOptions) Flags() (fss cliflag.NamedFlagSets) {
    // Add the generic flags
    s.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
    s.Kubernetes.AddFlags(fss.FlagSet("kubernetes"))
    s.Redis.AddFlags(fss.FlagSet("redis"))
    s.MySQL.AddFlags(fss.FlagSet("mysql"))
    s.Devops.AddFlags(fss.FlagSet("devops"))


    return fss
}
