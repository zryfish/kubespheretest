package simple

import (
    "github.com/gocraft/dbr"
    "github.com/zryfish/kubespheretest/pkg/simple/devops"
    "github.com/zryfish/kubespheretest/pkg/simple/mysql"
    "github.com/zryfish/kubespheretest/pkg/simple/redis"
)

type ClientSetOptions struct {
    mySQLOption  *mysql.MySQLOptions
    redisOption  *redis.RedisOptions
    devopsOption *devops.DevopsOptions
}

func NewClientSetOptions() *ClientSetOptions {
    return &ClientSetOptions{}
}

func (c *ClientSetOptions) SetMySQLOption(option *mysql.MySQLOptions) *ClientSetOptions {
    c.mySQLOption = option
    return c
}

func (c *ClientSetOptions) SetRedisOption(option *redis.RedisOptions) *ClientSetOptions {
    c.redisOption = option
    return c
}

func (c *ClientSetOptions) SetDevopsOption(option *devops.DevopsOptions) *ClientSetOptions {
    c.devopsOption = option
    return c
}

// ClientSet provide best of effort service to initialize clients,
// but there is no guarantee to return a valid client instance,
// so do validity check before use
type ClientSet struct {
    MySQL *mysql.MySQL
}

func NewClientSetFactory(c *ClientSetOptions, stopCh <-chan struct{}) *ClientSet {
    cs := &ClientSet{}

    // create clients one by one
    if c.mySQLOption != nil && c.mySQLOption.Host != "" {
        cs.MySQL = mysql.NewMySQL(c.mySQLOption, stopCh)
    }

    return cs
}

// no guarantee to return a valid connection
func (cs *ClientSet) GetMySQL() *dbr.Connection {
    if cs.MySQL != nil {
        return cs.MySQL.Client()
    }

    return nil
}
