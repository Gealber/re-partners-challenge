package orders

type serviceOrder struct{}

func New() *serviceOrder {
	return &serviceOrder{}
}

// PackOrder performs the packing of an order according to the rules specified
// n: number of items to order
// dimensions: allowed pack dimensions
// we are assiming that the dimensions are provided in a decreasing order
// for example [5000, 2000, 1000, 500, 250]
// return: a map with the packing distribution, for example
// for order with 501 items will return {500:1, 250:1}
func (srv *serviceOrder) PackOrder(n int, dimensions []int) map[int]int {
	// will contain the packing distribution
	packing := make(map[int]int)
	currentItemsNumber := n

	for _, packSize := range dimensions {
		if currentItemsNumber >= packSize {
			numPacks := currentItemsNumber / packSize
			packing[packSize] = numPacks
			currentItemsNumber = currentItemsNumber % packSize
		}

		if currentItemsNumber == 0 {
			break
		}
	}

	// this case is when all the items fit on the first smaller packing
	if currentItemsNumber > 0 {
		smallerPackSize := dimensions[len(dimensions)-1]
		packing[smallerPackSize]++

		// edge case when items fall between the two smaller packs size
		// for example 251, falls between packs of 250 and 500 dimensions
		if len(dimensions) > 1 && packing[smallerPackSize]*smallerPackSize >= dimensions[len(dimensions)-2] {
			delete(packing, smallerPackSize)
			packing[dimensions[len(dimensions)-2]] = 1
		}
	}

	return packing
}
