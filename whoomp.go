// whoomp
// author: Patrick Easters

package main

import (
  "fmt"
  "net/http"
  "log"
  "os"
  "github.com/mediocregopher/radix.v2/pool"
  "github.com/mediocregopher/radix.v2/redis"
)

var db *pool.Pool

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func countHandler(w http.ResponseWriter, r *http.Request){
  conn, err := db.Get()
  if err != nil {
  	http.Error(w, "Error getting connection to Redis", 500)
  }
  defer db.Put(conn)

  count, err := conn.Cmd("GET", "whoomps").Str()
  if err != nil {
  	http.Error(w, "Error getting count from Redis", 500)
  }
  fmt.Fprintf(w, "%s", count)
}

func incrHandler(w http.ResponseWriter, r *http.Request){
  if r.Method == "POST" {
    conn, err := db.Get()
    if err != nil {
    	http.Error(w, "Error getting connection to Redis", 500)
    }
    defer db.Put(conn)

    _, err = conn.Cmd("INCR", "whoomps").Str()
    if err != nil {
    	http.Error(w, "Error updating Redis", 500)
    }
  } else {
    http.Error(w, "Invalid request method", 405)
  }
}

func main() {
  redis_host := getEnv("APP_REDIS_HOST", "localhost")
  redis_port := getEnv("APP_REDIS_PORT", "6379")

  var err error

  df := func(network, addr string) (*redis.Client, error) {
  	client, err := redis.Dial(network, addr)
  	if err != nil {
  		return nil, err
  	}

    // silenty ignore error if we don't need auth
  	client.Cmd("AUTH", getEnv("APP_REDIS_AUTH", "") )

  	return client, nil
  }

  db, err = pool.NewCustom("tcp", fmt.Sprintf("[%s]:%s", redis_host, redis_port), 10, df)
  if err != nil {
    log.Println("Could not connect to redis:", err)
  }

  // handle index page
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
  })

  // handle static content
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  // the actual api functions
  http.HandleFunc("/api/count", countHandler)
  http.HandleFunc("/api/incr", incrHandler)

  log.Println("Starting up on port 3000")
  http.ListenAndServe(":3000", nil)
}
