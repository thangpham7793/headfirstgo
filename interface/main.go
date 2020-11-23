package main

import (
	"fmt"

	"github.com/headfirstgo/gadget"
)

func playList(p gadget.Player, songs []string) {

	for _, s := range songs {
		p.Play(s)
		//type assertion to get at type-specific method
		if p, ok := p.(*gadget.TapeRecorder); ok {
			p.Record()
		} else {
			fmt.Println("Not a recorder, skip recording!")
		}
	}
	p.Stop()

}

func main() {
	songs := []string{"We are the world", "Crying", "Hotel California"}

	//either declare the implementing type and pass in its address
	p := gadget.TapePlayer{}
	playList(&p, songs)

	//or declare a var of the interface type that takes the address value
	var r gadget.Player = &gadget.TapeRecorder{}
	playList(r, songs)

	//after all its *T, not T that implements the interface
}
