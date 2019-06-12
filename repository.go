package plutus

type Repository interface {
	CreateNewDiscountCode() *DiscountCode
}
