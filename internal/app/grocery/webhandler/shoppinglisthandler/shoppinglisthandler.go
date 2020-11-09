package shoppinglisthandler

import (
	"fmt"
	"net/http"

	"github.com/alexandrecodigo/gorestapimysql/internal/app/grocery/repo/shoppinglistrepo"
	"github.com/alexandrecodigo/gorestapimysql/internal/app/grocery/types"
	"github.com/alexandrecodigo/gorestapimysql/internal/app/grocery/util"
	"github.com/gin-gonic/gin"
)

// CreateHandler handler
func CreateHandler(c *gin.Context) {
	var shoppingList types.ShoppingList

	// App level validation
	bindErr := c.BindJSON(&shoppingList)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	// Inserting data
	id, insertErr := shoppinglistrepo.Create(shoppingList)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something wrong on our server"))
		util.PanicError(insertErr)
	} else {
		shoppingList.ID = id
		c.JSON(http.StatusCreated, shoppingList)
	}
}

// ShowHandler show
func ShowHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	shoppingList, _ := shoppinglistrepo.FindByID(id)

	// Check if resource exist
	if shoppingList.ID == 0 {
		c.JSON(http.StatusNotFound, "Not found")
	} else {
		c.JSON(http.StatusOK, shoppingList)
	}
}
