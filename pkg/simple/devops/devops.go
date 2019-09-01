package devops

import (
    "github.com/spf13/pflag"
    "github.com/zryfish/kubespheretest/pkg/simple/devops/sonarqube"
)

type DevopsOptions struct {
    Host string
    Username string
    Password string
    MaxConnections int

    SonarqubeOption *sonarqube.SonarqubeOptions
}

func NewDevopsOptions() *DevopsOptions{
    return &DevopsOptions{
        Host:           "",
        Username:       "",
        Password:       "",
        MaxConnections: 0,
        SonarqubeOption: sonarqube.NewSonarqubeOptions(),
    }
}

func (s *DevopsOptions) AddFlags(fs *pflag.FlagSet) {
    fs.StringVar(&s.Host, "jenkins-host", s.Host, ""+
        "Jenkins service host address. If left blank, means Jenkins "+
        "is unnecessary.")

    fs.StringVar(&s.Username, "jenkins-username", s.Username, ""+
        "Username for access to Jenkins service. Leave it blank if there isn't any.")

    fs.StringVar(&s.Password, "jenkins-password", s.Password, ""+
        "Password for access to Jenkins service, used pair with username.")

    fs.IntVar(&s.MaxConnections, "jenkins-max-connections", s.MaxConnections, ""+
        "Maximum allowed connections to Jenkins. ")

    s.SonarqubeOption.AddFlags(fs)

}