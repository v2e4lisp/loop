package main

import (
        "flag"
        "fmt"
        "os"
        "os/exec"
        "os/signal"
        "syscall"
        "time"
)

var (
        n float64
        s bool
)

func main() {
        flag.Usage = func() {
                fmt.Println("loop [-n [interval secs]] [-s] command")
                fmt.Println("\nOPTIONS:")
                flag.PrintDefaults()
        }
        flag.Float64Var(&n, "n", 2, "loop interval second")
        flag.BoolVar(&s, "s", false, "trap SIGTERM SIGINT SIGHUP")
        flag.Parse()
        args := flag.Args()

        span := time.Duration(n)
        var cmd *exec.Cmd

        if s {
                sigs := make(chan os.Signal, 1)
                signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
                go func() {
                        for {
                                fmt.Println("Received:", <-sigs)
                        }
                }()
        }

        for {
                if len(args) == 1 {
                        cmd = exec.Command("sh", "-c", args[0])
                        cmd.Stdout = os.Stdout
                        cmd.Stderr = os.Stderr
                        if err := cmd.Run(); err != nil {
                                fmt.Println(err)
                                os.Exit(1)
                        }
                }
                <-time.After(span * time.Second)
        }
}
