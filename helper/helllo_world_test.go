package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldHakim(t *testing.T) {
	result := HelloWorld("Hakim")
	if result != "Hello Hakim" {
		t.Fail()
	}

	// t.Fail() will fail the test but continue rest of code inside this test block
	fmt.Println("This code will be executed")
}

func TestHelloWorldFauzi(t *testing.T) {
	result := HelloWorld("Fauzi")
	if result != "Hello Fauzi" {

		// t.FailNow() will fail the test and end code cyle here.
		t.FailNow()

		fmt.Println("This code will not be executed")

	}
}
