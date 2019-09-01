package main

import (
    "github.com/zryfish/kubespheretest/cmd/ks-apiserver/app"
    "os"
)

func main() {
    command := app.NewAPIServerCommand()

    if err := command.Execute(); err != nil {
        os.Exit(1)
    }
}
