package shoppinglistrepo

import (
	"github.com/alexandrecodigo/gorestapimysql/internal/app/grocery/types"
	"github.com/alexandrecodigo/gorestapimysql/internal/app/grocery/util"
)

// Create list
func Create(shoppingList types.ShoppingList) (int64, error) {
	db := util.DBConn()
	defer db.Close()

	query := "INSERT INTO shopping_list (name, qty, unit) VALUES (?, ?, ?);"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	res, queryErr := stmt.Exec(shoppingList.Name, shoppingList.Qty, shoppingList.Unit)
	util.PanicError(queryErr)

	id, getLastInsertIDErr := res.LastInsertId()
	util.PanicError(getLastInsertIDErr)

	return id, queryErr
}

// FindByID Find the ID
func FindByID(id int64) (types.ShoppingList, error) {
	var shoppingList types.ShoppingList
	db := util.DBConn()
	defer db.Close()

	query := "SELECT id, name, qty, unit FROM shopping_list WHERE id = ?;"

	row := db.QueryRow(query, id)
	row.Scan(&shoppingList.ID, &shoppingList.Name, &shoppingList.Qty, &shoppingList.Unit)

	return shoppingList, nil
}

// Put update table
func Put(id int64, shoppingList types.ShoppingList) (types.ShoppingList, error) {
	db := util.DBConn()
	defer db.Close()

	query := "UPDATE shopping_list SET name = ?, qty = ?, unit = ? WHERE id = ?"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	_, queryErr := stmt.Exec(shoppingList.Name, shoppingList.Qty, shoppingList.Unit, id)
	util.PanicError(queryErr)

	shoppingList.ID = id
	return shoppingList, queryErr
}

// Delete from table
func Delete(shoppingList types.ShoppingList) error {
	db := util.DBConn()
	defer db.Close()

	query := "DELETE FROM shopping_list WHERE id = ?"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	_, queryErr := stmt.Exec(shoppingList.ID)
	util.PanicError(queryErr)

	return queryErr
}
