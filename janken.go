package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var sc = bufio.NewScanner(os.Stdin)
var handMap = map[string]int{
	"g": 0,
	"c": 1,
	"p": 2,
}

func main() {

	done := false

	for !done {

		fmt.Println("Please input your hand(G, C, P).")

		var player string
		if sc.Scan() {
			player = strings.ToLower(sc.Text())
		}

		rand.Seed(time.Now().UnixNano())
		cpuNum := rand.Intn(3)
		cpuHand := getHand(cpuNum)

		fmt.Println("your hand: " + player + " cpu hand: " + cpuHand)

		res := match(player, cpuNum)
		if res != 0 {
			done = true
		}
	}
}

func match(player string, cpu int) int {
	_, ok := handMap[player]
	if !ok {
		fmt.Println("invalid")
		return 0
	}

	switch handMap[player] - cpu {
	case 0:
		fmt.Println("even")
		return 0
	case 1, -2:
		fmt.Println("lose")
		return -1
	case 2, -1:
		fmt.Println("win")
		return 1
	}
	return 0
}

func getHand(num int) string {
	for s, n := range handMap {
		if n == num {
			return s
		}
	}
	return "invalid"
}
