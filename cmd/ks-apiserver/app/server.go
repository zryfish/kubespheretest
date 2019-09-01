package app

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/pflag"
    "github.com/zryfish/kubespheretest/cmd/ks-apiserver/app/options"
    cliflag "k8s.io/component-base/cli/flag"
    "k8s.io/klog"
    "os"
    "os/signal"
    "syscall"

    "k8s.io/apiserver/pkg/util/term"
)
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

var onlyOneSignalHandler = make(chan struct{})
var shutdownHandler chan os.Signal

func SetupSignalHandler() <-chan struct{} {
    close(onlyOneSignalHandler) // panics when called twice

    shutdownHandler = make(chan os.Signal, 2)

    stop := make(chan struct{})
    signal.Notify(shutdownHandler, shutdownSignals...)
    go func() {
        <-shutdownHandler
        close(stop)
        <-shutdownHandler
        os.Exit(1) // second signal. Exit directly.
    }()

    return stop
}

// PrintFlags logs the flags in the flagset
func PrintFlags(flags *pflag.FlagSet) {
    flags.VisitAll(func(flag *pflag.Flag) {
        klog.V(1).Infof("FLAG: --%s=%q", flag.Name, flag.Value)
    })
}

func NewAPIServerCommand() *cobra.Command {
    s := options.NewServerRunOptions()

    cmd := &cobra.Command{
        Use:  "apiserver",
        Long: `The apiserver validates and configures data 
for the api objects which include balabala.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            PrintFlags(cmd.Flags())

            return Run(s, SetupSignalHandler())
        },
    }

    fs := cmd.Flags()
    namedFlagSets := s.Flags()

    for _, f := range namedFlagSets.FlagSets {
        fs.AddFlagSet(f)
    }

    usageFmt := "Usage:\n  %s\n"
    cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
    /*cmd.SetUsageFunc(func(cmd *cobra.Command) error {
        fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
        cliflag.PrintSections(cmd.OutOrStderr(), namedFlagSets, cols)
        return nil
    })*/
    cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
        fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
        cliflag.PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
    })

    return cmd

}

func Run(serverRunOptions *options.ServerRunOptions, stopCh <-chan struct{}) error {
    return nil
}