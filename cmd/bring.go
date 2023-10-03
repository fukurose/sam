package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	porter "github.com/fukurose/sam/grpc"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client porter.PorterServiceClient
)

var bringCmd = &cobra.Command{
	Use:   "bring",
	Short: "Bring the file here",

	RunE: func(cmd *cobra.Command, args []string) error {
		from, err := cmd.Flags().GetString("from")
		if err != nil {
			return err
		}

		address, err := cmd.Flags().GetString("address")
		if err != nil {
			return err
		}

		conn, err := grpc.Dial(
			address,

			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)

		if err != nil {
			log.Fatal("Connection failed.")
			return err
		}
		defer conn.Close()

		client = porter.NewPorterServiceClient(conn)

		TransPort(from)
		return nil
	},
}

func TransPort(from string) {
	req := &porter.LSRequest{
		Path: from,
	}

	ls, err := client.ListSegmentStream(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		res, err := ls.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res.GetPath())
	}
}

func init() {
	bringCmd.Flags().StringP("from", "f", ".", "Path to bring from the server")
	rootCmd.AddCommand(bringCmd)
}
