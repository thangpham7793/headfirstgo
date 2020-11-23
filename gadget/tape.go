package gadget

import "fmt"

//Player definese methods needed to be a player
type Player interface {
	Play(string)
	Stop()
}

//TapePlayer a concrete type that satisfies SongPlayer
type TapePlayer struct {
	Batteries string
}

//Play concrete implementation
func (t *TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

//Stop concrete implementation
func (t *TapePlayer) Stop() {
	fmt.Println("Stopped")
}

//TapeRecorder a concrete type that satisfies SongPlayer
type TapeRecorder struct {
	Microphone int
}

//Record extra method outside of interface
func (t *TapeRecorder) Record() {
	fmt.Println("Recording")
}

//Play concrete implementation
func (t *TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

//Stop concrete implementation
func (t *TapeRecorder) Stop() {
	fmt.Println("Stopped")
}
