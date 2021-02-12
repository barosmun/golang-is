package main

import (
	"fmt"
	"math/rand"
	"time"
)

//This program simulates a bank heist, with a branching story using probability and conditionals

func main() {
	rand.Seed(time.Now().UnixNano())

	var isHeistOn bool = true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("We've made it passed the guards!")
	} else {
		isHeistOn = false
		fmt.Println("Couldn't get passed the guards! Mission failed")
	}

	if openedVault := rand.Intn(100); openedVault >= 60 && isHeistOn {
		fmt.Println("Vault is open!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Couldn't open the vault! Mission Failed")
	}

	if isHeistOn {
		switch leftSafely := rand.Intn(5); leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Couldn't get an Uber! Mission failed")
		case 1:
			isHeistOn = false
			fmt.Println("Got locked in the bank vault! Mission failed")
		case 2:
			isHeistOn = false
			fmt.Println("Gold bars too heavy for getaway! Mission failed")
		default:
			fmt.Println("Got away!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Printf("Score: $%v.00\n", amtStolen)
	}

	fmt.Printf("Heist ongoing: %v", isHeistOn)
}
