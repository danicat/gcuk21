package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/spf13/viper"
)

type Input int

const (
	KeyNone Input = iota
	KeyDefaultOrGraveyard
	KeyOpponentOrGraveyard
	KeyGraveyard
	KeyLeft
	KeyRight
	KeyUp
	KeyDown
	KeyQuit
)

type InputHandler interface {
	Read() Input
	Cancel()
}

type KeyboardHandler struct {
	ch   chan Input
	done chan struct{}
}

func NewKBHandler() *KeyboardHandler {
	var ih KeyboardHandler
	ch := make(chan Input)
	done := make(chan struct{})

	go func(i KeyboardHandler) {
		for {
			select {
			case <-done:
				return
			default:
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
				ch <- KeyDefaultOrGraveyard
			}
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				ch <- KeyOpponentOrGraveyard
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyG) {
				ch <- KeyGraveyard
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
				ch <- KeyLeft
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
				ch <- KeyRight
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
				ch <- KeyQuit
			}
		}
	}(ih)

	ih.ch = ch
	ih.done = done
	return &ih
}

func (ih *KeyboardHandler) Cancel() {
	ih.done <- struct{}{}
}

func (ih *KeyboardHandler) Read() Input {
	var input Input
	select {
	case input = <-ih.ch:
	default:
		input = KeyNone
	}
	return input
}

type MockHandler struct {
	keys []Input
}

func NewMockHandler() *MockHandler {
	return &MockHandler{}
}

func (mh *MockHandler) AppendKeys(keys []Input) {
	mh.keys = append(mh.keys, keys...)
}

func (mh *MockHandler) Read() Input {
	if len(mh.keys) == 0 {
		return KeyNone
	}
	key := mh.keys[0]
	mh.keys = mh.keys[1:]
	return key
}

func (mh *MockHandler) Cancel() {

}

type DumbHandler struct {
}

func NewDumbHandler() *DumbHandler {
	return &DumbHandler{}
}

func (dh *DumbHandler) Read() Input {
	time.Sleep(time.Millisecond * time.Duration(viper.GetInt("cpu_delay_ms")))
	return KeyDefaultOrGraveyard
}

func (dh *DumbHandler) Cancel() {
}
