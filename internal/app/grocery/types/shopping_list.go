package types

// ShoppingList struct List
type ShoppingList struct {
	ID   int64  `json:"id"`
	Name string `json:"name" binding:"required"`
	Qty  int    `json:"qty" binding:"required,gte=0"`
	Unit string `json:"unit" binding:"required"`
}
