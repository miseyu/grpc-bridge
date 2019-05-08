package main

import (
	"context"
	"fmt"
	"os"

	"github.com/envoyproxy/envoy/examples/grpc-bridge/client/kv"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

type set struct {
	key   string
	value string
}

type get struct {
	key string
}

var host string
var port string
var client kv.KVClient

func init() {
	host = os.Getenv("GRPC_HOST")
	if host == "" {
		host = "localhost"
	}
	port = os.Getenv("GRPC_PORT")
	if port == "" {
		port = "9211"
	}
	c, err := createClient()
	if err != nil {
		panic(err)
	}
	client = c
}

func main() {
	r := gin.Default()
	r.GET("/set", func(c *gin.Context) {
		key := c.Query("key")
		value := c.Query("value")
		req := kv.SetRequest{Key: key, Value: value}
		_, err := client.Set(context.Background(), &req)
		if err != nil {
			c.String(500, fmt.Sprintf("error %v", err))
			return
		}
		c.String(200, "Success")
	})
	r.GET("/get", func(c *gin.Context) {
		key := c.Query("key")
		req := kv.GetRequest{Key: key}
		res, err := client.Get(context.Background(), &req)
		if err != nil {
			c.String(500, "Error")
			return
		}
		c.String(200, fmt.Sprintf("key=%s value=%s", key, res.GetValue()))
	})
	r.Run("0.0.0.0:8080")
}

func createClient() (kv.KVClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return kv.NewKVClient(conn), nil
}
