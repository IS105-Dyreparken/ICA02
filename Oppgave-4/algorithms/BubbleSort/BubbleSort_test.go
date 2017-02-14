package algorithms

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// https://golang.org/doc/effective_go.html#init

//Lager et random seed ut ifra systemtiden
func init() {
	seed := time.Now().Unix()
	rand.Seed(seed)
}

//Oppretter en liste med random verdier som brukes i benchmarking funksjonen
func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}

// Skriv "benchmark"-tester for benchmarkBSortModified funksjonen
// Skriv en ny testfunksjon benchmarkBSortModified

func BenchmarkBubbleSort20(b *testing.B) {
	benchmarkBSort(20, b)
}

func BenchmarkBubbleSort80(b *testing.B) {
	benchmarkBSort(80, b)
}

func BenchmarkBubbleSort320(b *testing.B) {
	benchmarkBSort(320, b)
}

func benchmarkBSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		//Resetter timer slik at vi får riktig resultat
		b.StopTimer()
		//Lagrer den tilfeldig genererte listen som variablen values
		values := perm(i)
		//Restarter timer
		b.StartTimer()
		//Kjører selve BSort funksjonen med de tilfeldige verdiene
		benchmarkBSortModified(values)
	}
}

// Implementasjon av testfunksjoner for Quicksort algoritmen
func TestQSort(t *testing.T) {
	values := []int{9, 1, 20, 3, 6, 7}
	expected := []int{1, 3, 6, 7, 9, 20}

	QSort(values)

	if !reflect.DeepEqual(values, expected) {
		t.Fatalf("expected %d, actual is %d", 1, values[0])
	}
}

func BenchmarkQSort20(b *testing.B) {
	benchmarkQSort(20, b)
}

func BenchmarkQSort80(b *testing.B) {
	benchmarkQSort(80, b)
}

func BenchmarkQSort320(b *testing.B) {
	benchmarkQSort(320, b)
}

func benchmarkQSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		QSort(values)
	}
}
