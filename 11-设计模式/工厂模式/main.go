package 工厂模式

import "fmt"

func getFactory(gunType string) (IGun, error) {
	switch gunType {
	case "sk47":
		return NewAk47(), nil
	case "musket":
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}

func main() {
	ak47, _ := getFactory("ak47")
	musket, _ := getFactory("musket")

	printDetails(ak47)
	printDetails(musket)

	/*
		打印输出:

		Gun: AK47 gun
		Power: 4
		Gun: Musket gun
		Power: 1
	*/

}
