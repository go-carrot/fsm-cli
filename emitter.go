package fsmcli

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-carrot/fsm-emitable"
)

type CommandLineEmitter struct{}

func (c *CommandLineEmitter) Emit(i interface{}) error {
	switch v := i.(type) {

	case string:
		fmt.Println(v)
		return nil

	case emitable.Audio:
		fmt.Println("Audio:", v.URL)
		return nil

	case emitable.File:
		fmt.Println("File:", v.URL)
		return nil

	case emitable.Image:
		fmt.Println("Image:", v.URL)
		return nil

	case emitable.Video:
		fmt.Println("Video:", v.URL)
		return nil

	case emitable.QuickReply:
		fmt.Print("Type one of: [ ")
		for i, reply := range v.Replies {
			fmt.Print("'" + reply + "'")
			if i+1 < len(v.Replies) {
				fmt.Print(", ")
			}
		}
		fmt.Println(" ]")
		return nil

	case emitable.Typing:
		if v.Enabled {
			fmt.Println("...")
		}
		return nil
	}
	return errors.New("CommandLineEmitter cannot handle " + reflect.TypeOf(i).String())
}
