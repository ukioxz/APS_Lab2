package lab2

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	bufRead := make([]byte, 128)
	_, err := ch.Input.Read(bufRead)
	if err != nil {
		return err
	}
	bufRead = bytes.Trim(bufRead, "\x00")

	expression := string(bufRead)
	trimmed := strings.Trim(expression, " \n")
	res, err := PostfixCalculation(trimmed)
	if err != nil {
		return err
	}

	resStr := fmt.Sprint(res)
	resByte := []byte(resStr)
	_, err = ch.Output.Write(resByte)
	if err != nil {
		return err
	}
	return nil
}
