package main

import (
        "flag"
        "fmt"
        "os"
        "os/exec"
        "time"
)

var (
        t float64
)

func main() {
        flag.Float64Var(&t, "t", 5, "loop interval second")
        flag.Parse()
        args := flag.Args()

        span := time.Duration(t)
        var cmd *exec.Cmd

        for {
                <-time.After(span * time.Second)
                if len(args) > 0 {
                        cmd = exec.Command(args[0], args[1:]...)
                        cmd.Stdout = os.Stdout
                        cmd.Stderr = os.Stderr
                        if err := cmd.Run(); err != nil {
                                fmt.Println(err)
                                os.Exit(1)
                        }
                }
        }
}
