package main

import (
	"fmt"
	"github.com/nayutaya/geodelta-go/geodelta"
	"math/rand"
	"time"
)

/*
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
*/

func main() {
	rand.Seed(1)
	start_time := time.Now().UnixNano()

	ok, ng := 0, 0
	for i := 0; i < 50000; i++ {
		lat1 := rand.Float64()*(85*2) - 85
		lng1 := rand.Float64()*(180*2) - 180
		for level := byte(1); level <= 20; level++ {
			code1 := geodelta.GetDeltaCode(lat1, lng1, level)
			lat2, lng2 := geodelta.GetCenterFromDeltaCode(code1)
			code2 := geodelta.GetDeltaCode(lat2, lng2, level)
			// fmt.Printf("lat1=%f,lng1=%f\n", lat1, lng1)
			// fmt.Printf("lat2=%f,lng2=%f\n", lat2, lng2)
			// fmt.Printf("code1=%s\n", code1)
			// fmt.Printf("code2=%s\n", code2)
			if code1 == code2 {
				ok++
			} else {
				ng++
			}
		}
	}
	end_time := time.Now().UnixNano()
	fmt.Printf("time=%d\n", (end_time-start_time)/1000/1000)
	fmt.Printf("ok=%d\n", ok)
	fmt.Printf("ng=%d\n", ng)
}
