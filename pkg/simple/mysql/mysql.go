package mysql

import (
    "fmt"
    "github.com/gocraft/dbr"
    "github.com/spf13/pflag"
    "k8s.io/klog"
    "time"
)

type MySQLOptions struct {
    Host                  string
    Port                  int
    Username              string
    Password              string
    MaxIdleConnections    int
    MaxOpenConnections    int
    MaxConnectionLifeTime time.Duration
}

func NewMySQLOptions() *MySQLOptions {
    return &MySQLOptions{
        Host:                  "",
        Port:                  3306,
        Username:              "",
        Password:              "",
        MaxIdleConnections:    100,
        MaxOpenConnections:    100,
        MaxConnectionLifeTime: time.Duration(10) * time.Second,
    }
}

func (m *MySQLOptions) AddFlags(fs *pflag.FlagSet) {

    fs.StringVar(&m.Host, "mysql-host", m.Host, ""+
        "MySQL service host address, address should be reachable.")

    fs.IntVar(&m.Port, "mysql-port", m.Port, ""+
        "MySQL service port number, default to 3306.")

    fs.StringVar(&m.Username, "mysql-username", m.Username, ""+
        "Username for access to mysql service.")

    fs.StringVar(&m.Password, "mysql-password", m.Password, ""+
        "Password for access to mysql, should be used pair with password.")
}

type MySQL struct {
    *dbr.Connection
}

func NewMySQL(option *MySQLOptions, stopCh <-chan struct{}) *MySQL {
    m := &MySQL{}

    conn, err := dbr.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/devops?parseTime=1&multiStatements=1&charset=utf8mb4&collation=utf8mb4_unicode_ci", option.Username, option.Password, option.Host, option.Port), nil)
    if err != nil {
        klog.Fatal("unable to connect to mysql", err)
    }

    conn.SetMaxIdleConns(option.MaxIdleConnections)
    conn.SetConnMaxLifetime(option.MaxConnectionLifeTime)
    conn.SetMaxOpenConns(option.MaxOpenConnections)

    m.Connection = conn

    go func() {
        <-stopCh
        if err := conn.Close(); err != nil {
            klog.Warningf("error happened during closing mysql connection", err)
        }
    }()

    return m
}

func (m *MySQL) Client() *dbr.Connection {
    return m.Connection
}
