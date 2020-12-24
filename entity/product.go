package entity

//Product is struct
type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

//StockStatus is a method
func (p Product) StockStatus() (status string) {
	if p.Stock < 3 {
		status = "Stock hampir habis"
	} else if p.Stock < 10 {
		status = "Stock terbatas"
	} else {
		status = "Ready banyak boss"
	}

	return
}
