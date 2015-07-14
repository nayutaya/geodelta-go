package main

import (
	"fmt"
	"github.com/nayutaya/geodelta-go/geodelta/encoder"
	"math/rand"
	"time"
)

func main() {
	world := encoder.DeltaIds{0, 1, 2, 3, 4, 5, 6, 7}
	sub := encoder.DeltaIds{0, 1, 2, 3}
	rand.Seed(1)
	start_time := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		ids := encoder.DeltaIds{world[rand.Intn(len(world))]}
		for j := 0; j <= 20; j++ {
			ids = append(ids, sub[rand.Intn(len(sub))])
		}
		ids.Encode().Decode()
	}
	end_time := time.Now().UnixNano()
	fmt.Println("time=", (end_time-start_time)/1000/1000)
}
