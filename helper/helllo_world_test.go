package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These below are benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hanif")
	}
}
func BenchmarkHelloWorld2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Joko")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Hanif", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hani")
		}
	})
	b.Run("Fauzi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hani")
		}
	})
	b.Run("Hakim", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hanif")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Hanif",
			request: "Hanif",
		},
		{
			name:    "Fauzi",
			request: "Fauzi",
		},
		{
			name:    "Hakim",
			request: "Hakim",
		},
		{
			name:    "John Walllace Knight",
			request: "John Walllace Knight",
		},
	}

	for _, benchmarks := range benchmarks {
		b.Run(benchmarks.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmarks.request)
			}
		})
	}
}

func TestTableHelloWorld(t *testing.T) {

	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hanif",
			request:  "Hanif",
			expected: "Hello Hanif",
		},
		{
			name:     "John",
			request:  "John",
			expected: "Hello John",
		},
		{
			name:     "Edy",
			request:  "Edy",
			expected: "Hello Edy",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Hakim", func(t *testing.T) {
		result := HelloWorld("Hakim")
		require.Equal(t, "Hello Hakim", result, "Result must be 'Hello Hakim'")
	})

	t.Run("Joko", func(t *testing.T) {
		result := HelloWorld("Joko")
		require.Equal(t, "Hello Joko", result, "Result must be 'Hello Joko'")
	})
}

func TestMain(m *testing.M) {

	//Before run all unit test
	fmt.Println("BEFORE BEGIN UNIT TEST")

	m.Run() // run all test in this package (helper package)

	// After run all unit test
	fmt.Println("AFTER RUNNING UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on Mac OS") // will skip this unit test
	}

	result := HelloWorld("Johny")
	require.Equal(t, "Hello Johny", result, "Result must be Hello Johny")
}

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
