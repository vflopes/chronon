package chronon

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var (
	storeListenAddr string
	storeTls        bool
	storeTlsKey     string
	storeTlsCert    string

	storeServerListener net.Listener
	storeServer         *grpc.Server

	storeCmd = &cobra.Command{
		Use:   "store",
		Short: "(Event/Snapshot)Store server commands",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			opts := []grpc.ServerOption{}

			if storeTls {

				cert, err := tls.LoadX509KeyPair(storeTlsCert, storeTlsKey)
				if err != nil {
					log.Fatalln("Failed load TLS certificate (", storeTlsCert, ") or key (", storeTlsKey, ") because", err)
				}

				config := &tls.Config{
					Certificates: []tls.Certificate{cert},
					ClientAuth:   tls.VerifyClientCertIfGiven,
				}

				opts = append(opts, grpc.Creds(credentials.NewTLS(config)))

			}

			lis, err := net.Listen("tcp", storeListenAddr)
			if err != nil {
				log.Fatalln("Failed to listen to store address", storeListenAddr, "because", err)
			}

			storeServerListener = lis

			storeServer = grpc.NewServer(opts...)

		},
	}
)

func serveStoreServer() {

	reflection.Register(storeServer)

	log.Println("Server listening at:", storeServerListener.Addr())

	if err := storeServer.Serve(storeServerListener); err != nil {
		log.Fatalln("Failed to serve because", err)
	}

}

func init() {

	rootCmd.AddCommand(storeCmd)

	storeCmd.PersistentFlags().StringVar(&storeListenAddr, "listenAddr", "localhost:57755", "Bind address to store gRPC server")
	storeCmd.PersistentFlags().BoolVar(&storeTls, "tls", false, "Enable TLS protocol only on gRPC server")
	storeCmd.PersistentFlags().StringVar(&storeTlsKey, "tlsKey", "", "PEM encoded private key file path")
	storeCmd.PersistentFlags().StringVar(&storeTlsCert, "tlsCert", "", "PEM encoded certificate file path")

}
