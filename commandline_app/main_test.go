package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestIsPrime(t *testing.T) {
	// iniciamos con los test cases
	testCases := []struct {
		name string
		got  int
		want bool
		msg  string
	}{
		{"when is zero", 0, false, "0 is not prime, by definition!"},
		{"Negative", -3, false, "Negative numbers are not prime, by definition!"},
		{"Not prime", 4, false, "4 is not a primer number because it is divisible by 2"},
		{"Prime", 7, true, "7 is a prime number!"},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			prime, msg := isPrime(tC.got)

			if msg != tC.msg && prime != tC.want {
				t.Errorf("got %s and %t, \nwant %s and %t", msg, prime, tC.msg, tC.want)
			}

		})
	}
}

func Test_prompt(t *testing.T) {
	// Save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of out promt() func from our read pipe
	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Errorf("Incorrect prompt: expect '-> ' but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// Save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of out promt() func from our read pipe
	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Introduce un número") {
		t.Errorf("'intro' text not correct; got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{"Empty", "", "Por favor ingresa un numero entero"},
		{"Decimal", "1.1", "Por favor ingresa un numero entero"},
		{"With quit", "q", ""},
		{"Using zero", "0", "0 is not prime, by definition!"},
		{"Negative", "-2", "Negative numbers are not prime, by definition!"},
		{"Not prime", "8", "8 is not a primer number because it is divisible by 2"},
		{"Real prime", "7", "7 is a prime number!"},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// simulte the input from terminal
			input := strings.NewReader(tC.input)
			reader := bufio.NewScanner(input)
			res, _ := checkNumbers(reader)

			if res != tC.want {
				t.Errorf("Expected '%s'; got '%s'", tC.want, res)
			}
		})
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this function, we need a channel and an instance of an io.Reader
	doneChan := make(chan bool)

	// create a referece to a bytes.Buffer
	var stdin bytes.Buffer

	// simulates when we test the entire program with 1 and then quit (q)
	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)

	<-doneChan

	close(doneChan)
	
}