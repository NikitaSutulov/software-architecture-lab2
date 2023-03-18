package lab2

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsoleInputCorrect(t *testing.T) {
	stdout := os.Stdout
	read, write, _ := os.Pipe()
	os.Stdout = write
	handler := ComputeHandler{
		Input:  bytes.NewBufferString("+ 254 4"),
		Output: write,
	}
	err := handler.Compute()
	write.Close()
	out, _ := ioutil.ReadAll(read)
	os.Stdout = stdout
	if assert.Nil(t, err) {
		assert.Equal(t, "254 + 4", string(out[:]))
	} else {
		t.Errorf("Wrong result")
	}
}

func TestConsoleWrongChars(t *testing.T) {
	stdout := os.Stdout
	_, write, _ := os.Pipe()
	os.Stdout = write
	handler := ComputeHandler{
		Input:  bytes.NewBufferString("lorem ipsum"),
		Output: write,
	}
	err := handler.Compute()
	write.Close()
	os.Stdout = stdout
	if assert.NotNil(t, err) {
		assert.Equal(t, "Unable to convert", err.Error())
	} else {
		t.Errorf("Wrong result")
	}
}

func TestConsoleWrongNumberArgs(t *testing.T) {
	stdout := os.Stdout
	_, write, _ := os.Pipe()
	os.Stdout = write
	handler := ComputeHandler{
		Input:  bytes.NewBufferString("0 7 "),
		Output: write,
	}
	err := handler.Compute()
	write.Close()
	os.Stdout = stdout
	if assert.NotNil(t, err) {
		assert.Equal(t, "Unable to convert", err.Error())
	} else {
		t.Errorf("Wrong result")
	}
}
