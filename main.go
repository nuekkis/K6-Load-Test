package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
)

type Config struct {
    VUs       int    `json:"vus"`
    Duration  string `json:"duration"`
    TargetURL string `json:"target_url"`
}

func main() {
    data, err := ioutil.ReadFile("config.json")
    if err != nil {
        fmt.Fprintf(os.Stderr, "config.json could not be read: %v\n", err)
        os.Exit(1)
    }

    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil {
        fmt.Fprintf(os.Stderr, "config.json could not be parsed: %v\n", err)
        os.Exit(1)
    }

    cmd := exec.Command(
        "k6", "run",
        "--vus", fmt.Sprint(cfg.VUs),
        "--duration", cfg.Duration,
        "loadtest.js",
    )
    cmd.Env = append(os.Environ(), "TARGET_URL="+cfg.TargetURL)

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "k6 could not be run: %v\n", err)
        os.Exit(1)
    }
}
