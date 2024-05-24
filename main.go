package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"golang.ngrok.com/ngrok"
	logrok "golang.ngrok.com/ngrok/log"
)

type stdLogger struct{}

func (l stdLogger) Log(ctx context.Context, level int, msg string, data map[string]interface{}) {
	if level > logrok.LogLevelInfo {
		return
	}
	lvlString, _ := logrok.StringFromLogLevel(level)
	log.Printf("[%s] %s %v", lvlString, msg, data)
}

func main() {
	server := os.Getenv("NGROK_SERVER")
	if server == "" {
		server = "connect.ngrok-agent.com:443"
	}
	authToken := os.Getenv("NGROK_AUTHTOKEN")
	if authToken == "" {
		fmt.Println("please set NGROK_AUTHTOKEN")
		os.Exit(1)
	}

	var countArg string
	if len(os.Args) > 1 {
		countArg = os.Args[1]
	}

	count, _ := strconv.Atoi(countArg)
	if count == 0 {
		fmt.Println("usage: lotsagrok <count>")
		os.Exit(1)
	}

	g := sync.WaitGroup{}
	g.Add(count)

	ctx, cancel := context.WithCancel(context.Background())

	log.Println("starting", count, "sessions")

	for i := 0; i < count; i++ {
		go func() {
			_, err := ngrok.Connect(ctx,
				ngrok.WithAuthtoken(authToken),
				ngrok.WithServer(server),
				ngrok.WithTLSConfig(func(c *tls.Config) {
					c.InsecureSkipVerify = true
				}),
				ngrok.WithLogger(stdLogger{}),
			)
			if err != nil {
				cancel()
			}

			<-ctx.Done()
			g.Done()
		}()
	}

	g.Wait()
}
