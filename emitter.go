package fsmcli

import (
	"errors"
	"fmt"
)

type CommandLineEmitter struct{}

func (c *CommandLineEmitter) Emit(i interface{}) error {
	switch v := i.(type) {
	case string:
		fmt.Println(v)
		return nil
	}
	return errors.New("CommandLineEmitter can only handle strings")
}
