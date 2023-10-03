package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	porter "github.com/fukurose/sam/grpc"

	"github.com/gosuri/uiprogress"
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

		to, err := cmd.Flags().GetString("to")
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

		TransPort(from, to)
		return nil
	},
}

func TransPort(from string, to string) {
	req := &porter.LSRequest{
		Path: from,
	}

	ls, err := client.ListSegmentStream(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	total := 0
	uiprogress.Start()
	for {
		res, err := ls.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		rel, err := filepath.Rel(from, res.GetPath())
		if err != nil {
			fmt.Println(err)
		}
		toPath := filepath.Join(to, rel)
		total += 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			bar := uiprogress.AddBar(int(res.GetSize())).AppendCompleted()
			bar.PrependFunc(func(b *uiprogress.Bar) string {
				return toPath
			})
			Order(res, bar, toPath)
		}()
	}
	wg.Wait()
	uiprogress.Stop()
	fmt.Printf("\033[%dE\n", total)
}

func Order(file_info *porter.LSResponse, bar *uiprogress.Bar, toPath string) {
	req := &porter.OrderRequest{
		FilePath: file_info.GetPath(),
	}
	stream, err := client.OrderStream(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	current := 0
	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		err = os.MkdirAll(filepath.Dir(toPath), 0750)
		if err != nil {
			fmt.Println(err)
			break
		}

		file, err := os.Create(toPath)
		if err != nil {
			fmt.Println(err)
			break
		}
		defer file.Close()

		n, err := file.Write(res.GetData())
		if err != nil {
			fmt.Println(err)
			break
		}

		current += n
		bar.Set(current)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func init() {
	bringCmd.Flags().StringP("from", "f", ".", "Path to bring from the server")
	bringCmd.Flags().StringP("to", "t", ".", "Path to bring")
	rootCmd.AddCommand(bringCmd)
}
