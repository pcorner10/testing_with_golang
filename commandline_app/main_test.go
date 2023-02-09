package main

import "testing"

func TestIsPrime(t *testing.T) {
	// iniciamos con los test cases
	testCases := []struct {
		name string
		got  int
		want bool
		msg string
	}{
		{"when is zero", 0, false,"0 is not prime, by definition!"},
		{"Negative", -3, false,"Negative numbers are not prime, by definition!"},
		{"Not prime", 4, false,"4 is not a primer number because it is divisible by 2"},
		{"Prime", 7, true,"7 is a prime number!"},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			prime, msg := isPrime(tC.got)

			if msg != tC.msg && prime != tC.want{
				t.Errorf("got %s and %t, \nwant %s and %t", msg,prime, tC.msg, tC.want)
			}

		})
	}
}
