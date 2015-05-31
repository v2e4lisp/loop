package main

import (
        "flag"
        "fmt"
        "os"
        "os/exec"
        "os/signal"
        "strings"
        "syscall"
        "time"
)

var (
        interval time.Duration
        s        bool
)

func main() {
        flag.Usage = func() {
                fmt.Println("loop [-n interval] [-s] command")
                fmt.Println("\nOPTIONS:")
                flag.PrintDefaults()
        }
        flag.DurationVar(&interval, "n", 2*time.Second, "Interval between command execution")
        flag.BoolVar(&s, "s", false, "trap SIGTERM SIGINT SIGHUP")
        flag.Parse()
        if flag.NArg() < 1 {
                flag.Usage()
                os.Exit(1)
        }
        cmd := strings.Join(flag.Args(), " ")

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
                c := exec.Command("sh", "-c", cmd)
                c.Stdout = os.Stdout
                c.Stderr = os.Stderr
                if err := c.Run(); err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
                <-time.After(interval)
        }
}
