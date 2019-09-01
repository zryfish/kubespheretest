package redis

import (
    "fmt"
    "github.com/go-redis/redis"
    "github.com/spf13/pflag"
    "github.com/zryfish/kubespheretest/pkg/util/net"
    "k8s.io/klog"
)

type RedisOptions struct {
    Host     string
    Port     int
    Password string
    DB       int
}

// NewRedisOptions returns options points to nowhere,
// because redis is not required for some components
func NewRedisOptions() *RedisOptions {
    return &RedisOptions{
        Host:     "",
        Port:     0,
        Password: "",
        DB:       0,
    }
}

func (r *RedisOptions) Validate() []error {
    errors := make([]error, 0)

    if r.Host != "" {
        if !net.IsValidPort(r.Port) {
            errors = append(errors, fmt.Errorf("--redis-port is out of range"))
        }
    }

    return errors
}

func (r *RedisOptions) AddFlags(fs *pflag.FlagSet) {
    fs.StringVar(&r.Host, "--redis-host", r.Host, ""+
        "Redis service host address. If left blank, means redis is unnecessary, "+
        "redis is disabled")

    fs.IntVar(&r.Port, "--redis-port", r.Port, ""+
        "Redis service port number.")

    fs.StringVar(&r.Password, "--redis-password", r.Password, ""+
        "Redis service password if necessary, default to empty")

    fs.IntVar(&r.DB, "--redis-db", r.DB, ""+
        "Redis service database index, default to 0.")
}

type RedisClient struct {
    client *redis.Client
}

func NewRedisClient(option *RedisOptions, stopCh <-chan struct{}) *RedisClient{
    r := &RedisClient{}

    r.client = redis.NewClient(&redis.Options{
        Addr:               fmt.Sprintf("%s:%d", option.Host, option.Port),
        Password:           option.Password,
        DB:                 option.DB,
    })

    if err := r.client.Ping().Err(); err != nil {
        klog.Exit("unable to reach redis host", err)
    }

    go func(){
        <-stopCh
        if err := r.client.Close(); err != nil {
            klog.Error(err)
        }
    }()

    return r
}

func (r *RedisClient) Client() *redis.Client {
    return r.client
}
