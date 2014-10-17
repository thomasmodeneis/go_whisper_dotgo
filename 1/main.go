package main

import (
        "bufio"
        "encoding/json"
        "log"
        "os"
)

type Message struct {
        Body string
}

func main() {
        s := bufio.NewScanner(os.Stdin)
        e := json.NewEncoder(os.Stdout)
        for s.Scan() {
                m := Message{Body: s.Text()}
                err := e.Encode(m)
                if err != nil {
                        log.Fatal(err)
                }
        }
        if err := s.Err(); err != nil {
                log.Fatal(err)
        }
}