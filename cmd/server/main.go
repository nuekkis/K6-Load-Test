package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/spf13/viper"
    "github.com/icelaterdc/K6-Load-Test/internal/handler"
)

func main() {
    viper.SetConfigName("config")
    viper.AddConfigPath("./config")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("error reading config: %v", err)
    }

    r := chi.NewRouter()
    r.Get("/health", handler.Health)

    srv := &http.Server{
        Addr:         ":" + viper.GetString("server.port"),
        Handler:      r,
        ReadTimeout:  viper.GetDuration("server.read_timeout"),
        WriteTimeout: viper.GetDuration("server.write_timeout"),
    }

    fmt.Printf("Server is listening on %s\n", srv.Addr)
    if err := srv.ListenAndServe(); err != nil {
        log.Fatalf("server error: %v", err)
    }
}
