package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) int {
    if avg != 0 {
        return len(layers) * avg
    }
	return len(layers) * 2
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    noodles = 0
    sauce = 0.0
    for _, layer := range layers {
        if layer == "sauce" {
            sauce += 0.2
        } else if layer == "noodles" {
        	noodles += 50
        }
    }
	return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
    lastIndex1 := len(friendList) - 1
    lastIndex2 := len(myList) - 1
    myList[lastIndex2] = friendList[lastIndex1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    scale := float64(portions) * 0.5
    shoppingList := []float64{}
    for i := 0; i < len(quantities); i++ {
        shoppingList = append(shoppingList,quantities[i]*scale)
    }
	return shoppingList
}