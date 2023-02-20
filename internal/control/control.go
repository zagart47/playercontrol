package control

import (
	"google.golang.org/grpc"
	"log"
	"playercontrol/internal/config"
	"playercontrol/internal/flags"
)

func Run() {
	defer func(FileServer *grpc.ClientConn) {
		if err := FileServer.Close(); err != nil {
			log.Fatal(err)
		}
	}(config.Server)
	flags.Parse()
}
