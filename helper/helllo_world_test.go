package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Hanif")
	assert.Equal(t, "Hello Hanif", result, "Result must be Hello Hanif") // assert with Fail()
	fmt.Println("This still will be executed if error")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Hanif")
	require.Equal(t, "Hello Hanif", result, "Result must be Hello Hanif") // require with FailNow()
	fmt.Println("This will not be executed if error")
}

func TestHelloWorldHakim(t *testing.T) {
	result := HelloWorld("Hakim")
	if result != "Hello Hakim" {
		// t.Fail()

		// t.Error will Fail() with message description
		t.Error("Result must be Hello Hakim")
	}

	// t.Fail() will fail the test but continue rest of code inside this test block
	fmt.Println("This code will be executed")
}

func TestHelloWorldFauzi(t *testing.T) {
	result := HelloWorld("Fauzi")
	if result != "Hello Fauzi" {

		// t.FailNow() will fail the test and end code cyle here.
		// t.FailNow()

		// t.Fatal will FailNow() with message description
		t.Fatal("Result must be Hello Fauzi")
		fmt.Println("This code will not be executed")

	}
}
