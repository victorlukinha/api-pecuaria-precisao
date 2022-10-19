package animals

import (
	"github.com/labstack/echo"
	"strconv"
)

func FindByFilterHandler(c echo.Context) error {

	donoId, _ := strconv.Atoi(c.QueryParam("donoId"))
	//if err != nil {
	//	return c.JSON(400, "Informe o ID do dono.")
	//}

	animalId, _ := strconv.Atoi(c.QueryParam("animalId"))
	//if err != nil {
	//	return c.JSON(400, "Informe o ID do animal.")
	//}

	if donoId == 0 && animalId == 0 {
		return c.JSON(400, "Informe o ID do dono ou do animal.")
	}

	filter := Filter{
		OwnerId:  donoId,
		AnimalId: animalId,
	}

	animais := getAnimalsByFilter(filter)

	return c.JSON(200, animais)
}
