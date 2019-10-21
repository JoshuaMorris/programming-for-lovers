package gcd

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/JoshuaMorris/programming-for-lovers/testutils"
)

var (
	folder = "lecture-1/gcd"
)

func TestTrivialGCD(t *testing.T) {
	var tests = []struct {
		name    string
		x       int
		y       int
		fixture string
	}{
		{
			name:    "Should return an int of the TrivialGCD of two integers",
			x:       378,
			y:       273,
			fixture: "GCDTestTrivialGCDShouldReturnTwentyOne.golden",
		},
		{
			name:    "Should return an int of the TrivialGCD of two large integers",
			x:       378202873,
			y:       273147834,
			fixture: "GCDTestTrivialGCDShouldReturnOne.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrivialGCD(tt.x, tt.y)

			t.Logf("result: %v", result)

			if *testutils.Update {
				t.Log("Writing fixtures...")

				testutils.WriteFixture(t, folder, tt.fixture, []byte(strconv.Itoa(result)))
			}

			actual := []byte(strconv.Itoa(result))
			expected := []byte(testutils.LoadFixture(t, folder, tt.fixture))

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("actual = %v, epected = %v", actual, expected)
			}
		})
	}
}

func BenchmarkTrivialGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrivialGCD(b.N, b.N+1)
	}
}

func TestAnotherSumEven(t *testing.T) {
	var tests = []struct {
		name    string
		n       int
		fixture string
	}{
		{
			name:    "Should return the sum of ints up to 5",
			n:       5,
			fixture: "GCDTestAnotherSumEvenShouldReturnSix.golden",
		},
		{
			name:    "Should return the su of ints up to 350678",
			n:       350678,
			fixture: "GCDTestAnotherSumEvenShouldReturn30743940260.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AnotherSumEven(tt.n)

			t.Logf("result: %v", result)

			if *testutils.Update {
				t.Log("Writing fixtures...")

				testutils.WriteFixture(t, folder, tt.fixture, []byte(strconv.Itoa(result)))
			}

			actual := []byte(strconv.Itoa(result))
			expected := []byte(testutils.LoadFixture(t, folder, tt.fixture))

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("actual = %v, epected = %v", actual, expected)
			}
		})
	}
}

func BenchmarkAnotherSumEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AnotherSumEven(b.N)
	}
}

func TestMin2(t *testing.T) {
	var tests = []struct {
		name    string
		x       int
		y       int
		fixture string
	}{
		{
			name:    "Should return an int of the min of two integers",
			x:       378,
			y:       273,
			fixture: "GCDTestMin2ShouldReturnTwoHundredSeventyThree.golden",
		},
		{
			name:    "Should return an int of the min of two large integers",
			x:       21,
			y:       42,
			fixture: "GCDTestMin2ShouldReturnTwentyOne.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min2(tt.x, tt.y)

			t.Logf("result: %v", result)

			if *testutils.Update {
				t.Log("Writing fixtures...")

				testutils.WriteFixture(t, folder, tt.fixture, []byte(strconv.Itoa(result)))
			}

			actual := []byte(strconv.Itoa(result))
			expected := []byte(testutils.LoadFixture(t, folder, tt.fixture))

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("actual = %v, epected = %v", actual, expected)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	var tests = []struct {
		name    string
		x       int
		y       int
		fixture string
	}{
		{
			name:    "Should return an int of the GCD of two integers",
			x:       378,
			y:       273,
			fixture: "GCDTestGCDShouldReturnTwentyOne.golden",
		},
		{
			name:    "Should return an int of the GCD of two large integers",
			x:       378202873,
			y:       273147834,
			fixture: "GCDTestGCDShouldReturnOne.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GCD(tt.x, tt.y)

			t.Logf("result: %v", result)

			if *testutils.Update {
				t.Log("Writing fixtures...")

				testutils.WriteFixture(t, folder, tt.fixture, []byte(strconv.Itoa(result)))
			}

			actual := []byte(strconv.Itoa(result))
			expected := []byte(testutils.LoadFixture(t, folder, tt.fixture))

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("actual = %v, epected = %v", actual, expected)
			}
		})
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(b.N, b.N+1)
	}
}

func TestEuclidGCD(t *testing.T) {
	var tests = []struct {
		name    string
		x       int
		y       int
		fixture string
	}{
		{
			name:    "Should return an int of the EuclidGCD of two integers",
			x:       378,
			y:       273,
			fixture: "GCDTestEuclidGCDShouldReturnTwentyOne.golden",
		},
		{
			name:    "Should return an int of the GCD of two large integers",
			x:       378202873,
			y:       273147834,
			fixture: "GCDTestEuclidGCDShouldReturnOne.golden",
		},
		{
			name:    "Should return an int of the GCD of two equal integers",
			x:       42,
			y:       42,
			fixture: "GCDTestEuclidGCDShouldReturnFortyTwo.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EuclidGCD(tt.x, tt.y)

			t.Logf("result: %v", result)

			if *testutils.Update {
				t.Log("Writing fixtures...")

				testutils.WriteFixture(t, folder, tt.fixture, []byte(strconv.Itoa(result)))
			}

			actual := []byte(strconv.Itoa(result))
			expected := []byte(testutils.LoadFixture(t, folder, tt.fixture))

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("actual = %v, epected = %v", actual, expected)
			}
		})
	}
}

func BenchmarkEuclidGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EuclidGCD(b.N, b.N+1)
	}
}
