package sonarqube

import "github.com/spf13/pflag"

type SonarqubeOptions struct {
    Host  string `json:"host,omitempty" yaml:",omitempty"`
    Token string `json:"token,omitempty" yaml:",omitempty"`
}

func NewSonarqubeOptions() *SonarqubeOptions {
    return &SonarqubeOptions{
        Host:  "",
        Token: "",
    }
}

func (s *SonarqubeOptions) AddFlags(fs *pflag.FlagSet) {
    fs.StringVar(&s.Host, "sonarqube-host", s.Host, ""+
        "Sonarqube service address if enabled.")

    fs.StringVar(&s.Token, "sonarqube-token", s.Token, ""+
        "Sonarqube service access token.")
}
