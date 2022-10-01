package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type TDoor struct {
	Win      int
	Opened   bool
	Selected bool
}

func rnd2() int {
	return rand.Intn(2)
}

func rnd3() int {
	return rand.Intn(3)
}

func game(change bool) int {
	win := rnd3()
	first_sel := rnd3()

	if !change {
		if win == first_sel {
			return 1
		} else {
			return 0
		}
	}

	doors := make([]TDoor, 3)
	doors[win].Win = 1
	doors[first_sel].Selected = true

	var fst, fst_ bool
	if rnd2() == 1 {
		fst = true
	}

	for k := range doors {
		if (doors[k].Win == 0) && !doors[k].Selected {
			if win != first_sel {
				doors[k].Opened = true
				break
			} else {
				if fst || fst_ {
					doors[k].Opened = true
					break
				} else {
					fst_ = true
				}
			}
		}
	}

	for n := 0; n < 3; n++ {
		if !doors[n].Opened && !doors[n].Selected {
			return doors[n].Win
		}
	}

	log.Fatalln("restricted point")
	return 0
}

func main() {

	rand.Seed(time.Now().UnixNano())

	n := 100_000_000
	sum_A := 0
	sum_B := 0

	ch_A := make(chan int)
	ch_B := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			sum_A += game(true)
		}
		ch_A <- 0
	}()

	go func() {
		for i := 0; i < n; i++ {
			sum_B += game(false)
		}
		ch_B <- 0
	}()

	<-ch_A
	<-ch_B

	win_A := float64(sum_A) / float64(n)
	win_B := float64(sum_B) / float64(n)

	fmt.Println("A:", win_A)
	fmt.Println("B:", win_B)

}
