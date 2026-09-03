package myfunction

import "fmt"

func doA() {
	fmt.Println("Hiii")
}

func doB(x int, y int) {
	fmt.Println(x + y)
}

func doC() int {
	fmt.Println("Ho...")
	return 100
}

func doD() (int, string) {
	fmt.Println("Hew")
	return 555, "Hew"
}

func doE(p1 int, p2 string, p3 bool) (int, string, bool) {
	fmt.Println("Helo")
	return 100, "Helo", true

}
