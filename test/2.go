package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(generateRandomNickname())
	fmt.Println(rand.Int31n(8))
}
func generateRandomNickname() string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100000000)
	return fmt.Sprintf("user%08d", num)
}
