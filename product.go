package plutus

// Cost represents a economic cost
type Cost struct {
	Amount   float64
	Currency *Currency
}

// Product is a product description
type Product struct {
	Name    string
	Details string
	Cost    Cost
}

// Quantity is an int representing a quantity
type Quantity int

// ProductList is a quantized list of products
type ProductList map[Quantity]*Product
