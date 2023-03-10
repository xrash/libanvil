package libanvil

import (
	"fmt"
	"io"
	"bytes"
	"bufio"
)

type intermediateWriter struct {
	userWriter  io.Writer
	alreadyread bool
}

func newIntermediateWriter(userWriter io.Writer) *intermediateWriter {
	return &intermediateWriter{
		userWriter: userWriter,
	}
}

func (w *intermediateWriter) Write(b []byte) (int, error) {
	// This guarantess that the user writer will be ran.
	if w.userWriter != nil {
		defer w.userWriter.Write(b)
	}

	fmt.Println(1)
	if !w.alreadyread {
		fmt.Println(2)
		w.readAccounts(b)
	}

	w.alreadyread = true

	return 0, nil
}

func (w *intermediateWriter) readAccounts(b []byte) {
	buffer := bytes.NewBuffer(b)
	scanner := bufio.NewScanner(buffer)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(">", s)
	}
}
