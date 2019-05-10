package main

import (
	"fmt"
	"time"
)

func main() {
        start := time.Date(1985, 10, 9, 12, 30, 0, 0, time.UTC)

	fmt.Println(start)
	fmt.Println(time.Now())
	
        // calculare ani, luni, zile si timp dintre date
        an, luna, zi, ora, min, sec := diff(start, time.Now())

        fmt.Printf("Diferenta: %d ani, %d luni, %d zile, %d ore, %d minute and %d secunde.", an, luna, zi, ora, min, sec)
        fmt.Printf("")

        // calculare total numar de zile
	duration := time.Now().Sub(start)
	fmt.Printf("\nDiferenta in zile: %d", int(duration.Hours()/24) )
}

func diff(a, b time.Time) (an, luna, zi, ora, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	an = int(y2 - y1)
	luna = int(M2 - M1)
	zi = int(d2 - d1)
	ora = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalizarea valoriilor negative
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		ora--
	}
	if ora < 0 {
		ora += 24
		zi--
	}
	if zi < 0 {
		// zile in luna:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		zi += 32 - t.Day()
		luna--
	}
	if luna < 0 {
		luna += 12
		an--
	}

	return
}