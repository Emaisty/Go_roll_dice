package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed     = flag.Int("seed", 0, "seed for random generator. unix(now) be default")
	start    = flag.Int("start", 1, "Left border of rolling")
	end      = flag.Int("end", 6, "Right border of rolling")
	n        = flag.Int("n", 1, "how much cubes will roll")
	norepeat = flag.Bool("norepeat", false, "dont repeat values")
)

func randInterval_norepeat(l, r int) []int {
	m := make([]int, r-l+1)
	for i := l; i < r; i++ {
		j := rand.Intn(i)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

func randInterval(l, r int) int {
	return rand.Intn(r-l+1) + l
}

func main() {
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}
	if *n > (*end-*start+1) && *norepeat {
		fmt.Println("IMPOSSIBLE")
	} else {
		if *norepeat {
			arr := randInterval_norepeat(*start, *end)
			for i := 0; i < *n; i++ {
				print(arr[i] + 1)
				if i != (*n - 1) {
					print(",")
				}
			}
		} else {
			rand.Perm(*n)
			for i := 0; i < *n; i++ {
				fmt.Print(randInterval(*start, *end))
				if i != (*n - 1) {
					print(",")
				}
			}
		}

	}
}
