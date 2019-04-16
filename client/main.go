package main

import (
	"context"
	"fmt"
	"os"

	"github.com/envoyproxy/envoy/examples/grpc-bridge/client/kv"

	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

type set struct {
	key   string
	value string
}

type get struct {
	key string
}

var setOpt set
var getOpt get
var host string
var port string

var rootCmd = &cobra.Command{
	Use:           "kv",
	Short:         "key value memory store",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(
		setCmd(),
		getCmd(),
	)
	host = os.Getenv("GRPC_HOST")
	if host == "" {
		host = "localhost"
	}
	port = os.Getenv("GRPC_PORT")
	if port == "" {
		port = "8081"
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}
}

func setCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set key value",
		Short: "set keyValue",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := createClient()
			if err != nil {
				fmt.Fprintf(os.Stderr, "grpc connect error: %v\n", err)
				return
			}
			req := kv.SetRequest{Key: setOpt.key, Value: setOpt.value}
			_, err = c.Set(context.Background(), &req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "grpc set: %v\n", err)
				return
			}
			fmt.Print("grpc set finished\n")
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&setOpt.key, "key", "k", "", "key")
	flags.StringVarP(&setOpt.value, "value", "v", "", "value")
	return cmd
}

func getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get key",
		Short: "get keyValue",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := createClient()
			if err != nil {
				fmt.Fprintf(os.Stderr, "grpc connect error: %v\n", err)
				return
			}
			req := kv.GetRequest{Key: getOpt.key}
			res, err := c.Get(context.Background(), &req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "grpc get: %v\n", err)
				return
			}
			fmt.Printf("%s = %s", getOpt.key, res.GetValue())
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&getOpt.key, "key", "k", "", "key")
	return cmd
}

func createClient() (kv.KVClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return kv.NewKVClient(conn), nil
}
