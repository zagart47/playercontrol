package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
	"playercontrol/internal/config"
	pb "playercontrol/proto"
	"time"
)

type Client struct {
	client pb.PlayerServiceClient
}

func New(conn grpc.ClientConnInterface) Client {
	return Client{client: pb.NewPlayerServiceClient(conn)}
}

func (c Client) Play() (pb.PlayerService_PlayClient, error) {
	playerClient := New(config.Server)
	stream, err := playerClient.client.Play(context.Background(), &pb.PlayRequest{})
	if err != nil {
		return nil, err
	}
	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Println(req.Info)
	}
	return nil, err
}

func (c Client) Pause() {
	//TODO implement me
	panic("implement me")
}

func (c Client) Next() {
	//TODO implement me
	panic("implement me")
}

func (c Client) Prev() {
	//TODO implement me
	panic("implement me")
}

func (c Client) AddSong(file string) error {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("cannot open file (%w)", err)
	}
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}

	fileName := stat.Name()

	md := metadata.Pairs("filename", fileName)
	mdCtx := metadata.NewOutgoingContext(context.Background(), md)

	buffer := make([]byte, 1024*1024)
	fileUpload := pb.NewPlayerServiceClient(config.Server)
	stream, err := fileUpload.AddSong(mdCtx)
	if err != nil {
		return err
	}

	for {
		bytes, err := f.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("buffer reading error (%w)", err)
		}
		req := &pb.AddSongRequest{Chunk: buffer[:bytes]}
		if err := stream.Send(req); err != nil {
			log.Fatal(err.Error())
		}
	}
	if _, err = stream.CloseAndRecv(); err != nil {
		return fmt.Errorf("cannot send file to server (%w)", err)
	}

	fmt.Println("file uploaded:", fileName)
	return nil
}

func (c Client) UpdateSong(song string) error {
	//TODO implement me
	panic("implement me")
}

func (c Client) DeleteSong(song string) error {
	//TODO implement me
	panic("implement me")
}
