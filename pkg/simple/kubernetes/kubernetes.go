package kubernetes

import (
    "fmt"
    "github.com/spf13/pflag"
    "os"
)

type KubernetesOptions struct {
    KubeConfig string
    QPS int
    Burst int
}

func NewKubernetesOptions() *KubernetesOptions {
    return &KubernetesOptions{
        KubeConfig: fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")),
        QPS:        0,
        Burst:      0,
    }
}

func (k *KubernetesOptions) AddFlags(fs *pflag.FlagSet) {
    fs.StringVar(&k.KubeConfig, "kubeconfig", k.KubeConfig, ""+
        "Path for kubernetes kubeconfig file, if left blank, will use "+
        "in cluster way.")
}