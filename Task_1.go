package main

import (
	"errors"
	"fmt"
)

func main() {
	productList := productsMap()
	history := makeSlice()
	statistic := statistick()

	for {
		option, err := getOption()

		if err != nil {
			fmt.Println(err)
			statistic["Errors count"] += 1
			continue
		}

		if option == 7 {
			fmt.Println("\nGoodBye!")
			break
		}

		switch option {
		case 1:
			printProducts(productList)

		case 2:
			err, history = addProduct(productList, history, statistic)
			if err != nil {
				fmt.Println(err)
				statistic["Errors count"] += 1
				continue
			}

		case 3:
			err, history = sellProduct(productList, history, statistic)

			if err != nil {
				fmt.Println(err)
				statistic["Errors count"] += 1
				continue
			}

		case 4:
			err, history = restockProduct(productList, history, statistic)

			if err != nil {
				fmt.Println(err)
				statistic["Errors count"] += 1
				continue
			}

		case 5:
			printHistory(history)

		case 6:
			printStatistic(statistic)
		default:
			fmt.Println("Invalid menu option")
			statistic["Errors count"] += 1
		}

		if len(history) > 10 {
			history = history[1:]
		}
	}

}

func productsMap() map[string]int {
	productList := make(map[string]int)
	return productList
}

func getOption() (int, error) {
	var option int
	fmt.Print("===== Warehouse Manager =====\n1. Show products\n2. Add product\n3. Sell product\n4. Restock product\n5. Show operation history\n6. Show statistics\n7. Exit\nChoose option: ")
	_, err := fmt.Scanln(&option)

	if err != nil {
		return 0, errors.New("invalid option type")
	}

	return option, nil
}

func printProducts(prodList map[string]int) {
	if len(prodList) > 0 {
		fmt.Println("\n\n===== Products =====")
		for productName, amount := range prodList {
			fmt.Println(productName, ": ", amount)
		}
		fmt.Println("")

	} else {
		fmt.Println("\nWarehouse is empty")
	}
}

func addProduct(prodList map[string]int, history []string, stat map[string]int) (error, []string) {
	var productName string
	var quantity int

	fmt.Print("Enter product name: ")
	fmt.Scanln(&productName)
	fmt.Print("Enter quantity: ")
	_, err := fmt.Scanln(&quantity)

	if err != nil {
		return errors.New("invalid quantity type"), history
	}

	if productName == "" {
		return errors.New("name must not be empty"), history
	}

	if quantity <= 0 {
		return fmt.Errorf("quantity cannot be less than zero or zero as your input %d", quantity), history
	}

	_, ok := prodList[productName]

	if ok == true {
		return errors.New("Product already exists"), history
	}

	prodList[productName] = quantity

	addMessage := fmt.Sprintf("Added %q: +%d", productName, quantity)
	history = append(history, addMessage)
	stat["Products added"] += 1

	return nil, history
}

func sellProduct(prodList map[string]int, history []string, stat map[string]int) (error, []string) {
	productName, quantity, err := askProductAndQuantity(prodList)

	if err != nil {
		return err, history
	}

	//_, ok := prodList[productName]

	/*if !ok {
		return errors.New("Product not found")
	}*/

	if prodList[productName] < quantity {
		return fmt.Errorf("Not enough %d stock", quantity-prodList[productName]), history
	}

	prodList[productName] = prodList[productName] - quantity

	if prodList[productName] == 0 {
		delete(prodList, productName)
		stat["Products removed"] += 1
		removeHistoryMessage := fmt.Sprintf("Removed %q after selling all units", productName)
		history = append(history, removeHistoryMessage)
	}

	fmt.Println("Product sold successfully")

	sellHistoryMessage := fmt.Sprintf("Sold %q: -%d", productName, quantity)
	history = append(history, sellHistoryMessage)
	stat["Products sold"] += 1

	return nil, history
}

func askProductAndQuantity(prodList map[string]int) (string, int, error) {
	var productName string
	var quantity int

	fmt.Print("Enter product name: ")
	fmt.Scanln(&productName)

	_, ok := prodList[productName]

	if !ok {
		return "", 0, errors.New("Product not found")
	}
	fmt.Print("Enter product quantity: ")
	_, err := fmt.Scanln(&quantity)

	if err != nil {
		return "", 0, errors.New("\nincorrect quantity type")
	}

	if productName == "" {
		return "", 0, errors.New("\ninput must not be empty")
	}

	if quantity <= 0 {
		return "", 0, errors.New("\nInvalid quantity")
	}

	return productName, quantity, nil
}

/*func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}*/

func restockProduct(prodList map[string]int, history []string, stat map[string]int) (error, []string) {
	productName, quantity, err := askProductAndQuantity(prodList)

	if err != nil {
		return fmt.Errorf("cannot restore nothing, reason %w", err), history
	}

	/*_, ok := prodList[productName]

	if !ok {
		return errors.New("Product not found")
	}*/

	prodList[productName] = prodList[productName] + quantity
	fmt.Println("Product restocked successfully")

	restockHistoryMessage := fmt.Sprintf("Restocked %q +%d", productName, quantity)
	history = append(history, restockHistoryMessage)
	stat["Products restocked"] += 1
	return nil, history
}

func makeSlice() []string {
	var slice []string
	return slice
}

func statistick() map[string]int {
	statistic := map[string]int{
		"Products added":     0,
		"Products sold":      0,
		"Products restocked": 0,
		"Products removed":   0,
		"Errors count":       0,
	}
	return statistic
}

func printStatistic(stat map[string]int) {
	fmt.Println("\n===== Statistics =====")
	for productNames, quantity := range stat {
		fmt.Println(productNames, ": ", quantity)
	}
	fmt.Println("")
}

func printHistory(history []string) {
	if len(history) > 0 {
		fmt.Println("\n===== Operation History =====")
		for _, amount := range history {
			fmt.Println(": ", amount)
		}
		fmt.Println()
	} else {
		fmt.Println("History is empty")
	}
}
