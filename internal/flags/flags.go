package flags

import (
	"errors"
	"flag"
	"log"
	"playercontrol/internal/client"
	"playercontrol/internal/config"
)

type Connect struct {
	client client.Client
}

func NewConnect() *Connect {
	return &Connect{
		client: client.New(config.Server),
	}
}

func Parse() {
	if err := Perform(); err != nil {
		log.Fatal(err)
	}
}

type Flag struct {
	play  bool
	pause bool
	next  bool
	prev  bool
	add   string
	del   string
	upd   string
}

func NewFlags() Flag {
	return Flag{
		play:  false,
		pause: false,
		next:  false,
		prev:  false,
		add:   "",
		del:   "",
		upd:   "",
	}
}

var flags = NewFlags()

func Perform() error {
	connect := NewConnect()
	flag.BoolVar(&flags.play, "play", false, "play song")
	flag.BoolVar(&flags.pause, "pause", false, "pause playing")
	flag.BoolVar(&flags.next, "next", false, "next song")
	flag.BoolVar(&flags.prev, "prev", false, "prev song")
	flag.StringVar(&flags.add, "add", "", "add song")
	flag.StringVar(&flags.del, "del", "", "delete song")
	flag.StringVar(&flags.upd, "upd", "", "update song")
	flag.Parse()

	switch {
	case flags.play:
		connect.client.Play()
		return nil
	case flags.pause:
		connect.client.Pause()
		return nil
	case flags.next:
		connect.client.Next()
		return nil
	case flags.prev:
		connect.client.Prev()
		return nil
	case len(flags.add) > 0:
		if err := connect.client.AddSong(flags.add); err != nil {
			return err
		}
	case len(flags.del) > 0:
		if err := connect.client.DeleteSong(flags.del); err != nil {
			return err
		}
	case len(flags.upd) > 0:
		if err := connect.client.UpdateSong(flags.upd); err != nil {
			return err
		}
	default:
		return errors.New("operation not allowed")
	}
	return nil

}
