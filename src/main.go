package main

import "fmt"

func main() {
	var (
		priceApple = 599
		pricePear  = 700
		countApple = 9
		countPear  = 8
		debit      = 2300
	)

	fmt.Printf("Скільки грошей треба витратити, щоб купити %d яблук та %d груш?\n", countApple, countPear)
	hryvna, copiika := copiikaToHryvna(countPear*pricePear + countApple*priceApple)
	fmt.Printf("Треба витратити %d гривень та %d копійок.\n\n", hryvna, copiika)

	fmt.Println("Скільки груш ми можемо купити?")
	fmt.Printf("Ми можемо купити %d груші. \n\n", canBuy(debit, pricePear))

	fmt.Println("Скільки яблук ми можемо купити?")
	fmt.Printf("Ми можемо купити %d яблука. \n\n", canBuy(debit, priceApple))

	countApple = 2
	countPear = 2
	fmt.Printf("Чи ми можемо купити %d груші та %d яблука? (true/false)\n", countPear, countApple)
	fmt.Println("Відповідь:", countPear*pricePear+countApple*priceApple <= debit)
}

func canBuy(debit, priceGood int) int {
	return debit / priceGood
}

func copiikaToHryvna(copiika int) (int, int) {
	return copiika / 100, copiika % 100
}
