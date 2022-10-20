package owner

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func getOwnerHandler(c echo.Context) error {

	cpf := c.QueryParam("cpf")

	if cpf == "" {
		return c.JSON(400, "Informe o cpf do dono.")
	}
	filter := Filter{
		Cpf: cpf,
	}
	owner, err := getOwner(filter)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(http.StatusOK, owner)
}

func createOwnerHandler(c echo.Context) error {

	owner := Owner{}
	c.Bind(&owner)

	owner, err := createOwner(owner)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(http.StatusOK, owner)
}

func updateOwnerHandler(c echo.Context) error {

	owner := Owner{}
	c.Bind(&owner)

	owner, err := updateOwner(owner)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(http.StatusOK, owner)
}

func deleteOwnerHandler(c echo.Context) error {

	ownerId, _ := strconv.Atoi(c.Param("ownerId"))

	err := deleteOwner(int64(ownerId))
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(http.StatusOK, "Dono deletado com sucesso.")
}
