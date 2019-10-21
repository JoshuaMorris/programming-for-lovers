package primes

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/JoshuaMorris/programming-for-lovers/testutils"
)

var (
	folder = "lecture-1/primes"
)

func TestTrivialPrimeFinder(t *testing.T) {
	var tests = []struct {
		name    string
		n       int
		fixture string
	}{
		{
			name:    "Should return a boolean arry of five primes",
			n:       5,
			fixture: "PrimesTestTrivialPrimeFinderShouldReturnFivePrimes.golden",
		},
		{
			name:    "Should return a boolean array of the first 50 primes",
			n:       50,
			fixture: "PrimesTestTrivialPrimeFinderShouldReturnFiftyPrimes.golden",
		},
		{
			name:    "Should return a boolean array of the first 10000 primes",
			n:       10000,
			fixture: "PrimesTestTrivialPrimeFinderShouldReturnTenThousandPrimes.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrivialPrimeFinder(tt.n)

			t.Logf("result: %v", result)

			if *testutils.Update {
				t.Log("Writing fixtures...")

				testutils.WriteFixture(t, folder, tt.fixture, testutils.BoolsToBytes(result))
			}

			actual := testutils.BoolsToBytes(result)
			expected := []byte(testutils.LoadFixture(t, folder, tt.fixture))

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("actual = %v, expected = %v", actual, expected)
			}
		})
	}
}

func BenchmarkTrivialPrimeFinder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrivialPrimeFinder(b.N)
	}
}

func TestIsPrime(t *testing.T) {
	if *testutils.Update {
		t.Log("nothing to update")
	}
	var tests = []struct {
		name     string
		p        int
		expected bool
	}{
		{
			name:     "Should return true for '2' being a prime number",
			p:        2,
			expected: true,
		},
		{
			name:     "Should return false for '6' not being a prime number",
			p:        6,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPrime(tt.p)

			t.Logf("result: %v", result)

			assert.Equal(t, result, tt.expected)
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(b.N)
	}
}

func TestSieveOfEratosthenes(t *testing.T) {
	var tests = []struct {
		name    string
		n       int
		fixture string
	}{
		{
			name:    "Should return a boolean arry of five primes",
			n:       5,
			fixture: "PrimesTestSieveOfEratosthenesShouldReturnFivePrimes.golden",
		},
		{
			name:    "Should return a boolean array of the first 50 primes",
			n:       50,
			fixture: "PrimesTestSieveOfEratosthenesShouldReturnFiftyPrimes.golden",
		},
		{
			name:    "Should return a boolean array of the first 10000 primes",
			n:       10000,
			fixture: "PrimesTestSieveOfEratosthenesShouldReturnTenThousandPrimes.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SieveOfEratosthenes(tt.n)

			t.Logf("result: %v", result)

			if *testutils.Update {
				t.Log("Writing fixtures...")
				testutils.WriteFixture(t, folder, tt.fixture, testutils.BoolsToBytes(result))
			}

			actual := testutils.BoolsToBytes(result)
			expected := []byte(testutils.LoadFixture(t, folder, tt.fixture))

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("actual = %v, expected = %v", actual, expected)
			}
		})
	}
}

func BenchmarkSieveOfEratosthenes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SieveOfEratosthenes(b.N)
	}
}

func TestCrossOffMultiples(t *testing.T) {
	var tests = []struct {
		name       string
		primeArray []bool
		p          int
		fixture    string
	}{
		{
			name:       "Should return a bool array with multiples of two set to false",
			primeArray: []bool{true, true, true, true, true, true, true, true, true, true},
			p:          2,
			fixture:    "PrimesTestCrossOffMultiplesShouldReturnPrimeArrayWithMultiplesOfTwoFalse.golden",
		},
		{
			name:       "Should return a bool array with multiples of three set to false",
			primeArray: []bool{true, true, true, true, true, true, true, true, true, true},
			p:          3,
			fixture:    "PrimesTestCrossOffMultiplesShouldReturnPrimeArrayWithMultiplesOfThreeFalse.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CrossOffMultiples(tt.primeArray, tt.p)

			t.Logf("result: %v", result)

			if *testutils.Update {
				t.Log("Writing fixtures...")
				testutils.WriteFixture(t, folder, tt.fixture, testutils.BoolsToBytes(result))
			}

			actual := testutils.BoolsToBytes(result)
			expected := []byte(testutils.LoadFixture(t, folder, tt.fixture))

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("actual = %v, expected = %v", actual, expected)
			}
		})
	}
}

func BenchmarkCrossOffMultiples(b *testing.B) {
	primeArray := make([]bool, b.N)

	for j := 0; j < b.N; j++ {
		primeArray[j] = true
	}

	for i := 0; i < b.N; i++ {
		CrossOffMultiples(primeArray, b.N)
	}
}
