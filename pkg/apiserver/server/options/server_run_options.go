package options

import (
    "fmt"
    "github.com/spf13/pflag"
    "net"
    "time"
)

// ServerRunOptions contains the options while running a generic api server.
type ServerRunOptions struct {
    BindAddress net.IP

    InsecurePort int

    SecurePort int
    TlsCertFile string
    TlsPrivateKeyFile string

    RequestTimeout time.Duration
}

func NewServerRunOptions() *ServerRunOptions {
    return &ServerRunOptions{
        BindAddress:  net.IPv6loopback,
        InsecurePort:      0,
        SecurePort:        0,
        TlsCertFile:       "",
        TlsPrivateKeyFile: "",
        RequestTimeout:    time.Duration(30) * time.Second,
    }
}

func (s* ServerRunOptions) Validate() []error {
    errors := []error{}

    if s.InsecurePort < 0 || s.InsecurePort > 65535 {
        errors = append(errors, fmt.Errorf("--insecure-port is out of port range"))
    }

    if s.SecurePort < 0 || s.SecurePort > 65535 {
        errors = append(errors, fmt.Errorf("--secure-port is out of port range"))
    }

    return errors
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
    fs.IPVar(&s.BindAddress, "bind-address", s.BindAddress, ""+
        "The IP address on which to advertise the apiserver. If blank, the host's default "+
        "interface will be used.")

    fs.IntVar(&s.InsecurePort, "insecure-port", s.InsecurePort, ""+
        "The insecure listening port which apiserver will listen. --inscure-port and "+
        "--secure-port should not be left blank at the same time.")

    fs.IntVar(&s.SecurePort, "secure-port", s.SecurePort, ""+
        "")

    fs.StringVar(&s.TlsCertFile, "tls-cert-file", s.TlsCertFile, ""+
        "")

    fs.StringVar(&s.TlsPrivateKeyFile, "tls-private-key-file", s.TlsPrivateKeyFile, "")
}