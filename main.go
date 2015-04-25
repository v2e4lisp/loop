package main

import (
        "flag"
        "fmt"
        "os"
        "os/exec"
        "time"
)

var (
        n float64
)

func main() {
        flag.Float64Var(&n, "n", 2, "loop interval second")
        flag.Parse()
        args := flag.Args()

        span := time.Duration(n)
        var cmd *exec.Cmd

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
