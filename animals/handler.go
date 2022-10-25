package animals

import (
	"github.com/labstack/echo"
	"strconv"
)

func FindByFilterHandler(c echo.Context) error {

	donoId, _ := strconv.Atoi(c.QueryParam("donoId"))
	animalId, _ := strconv.Atoi(c.QueryParam("animalId"))

	if donoId == 0 && animalId == 0 {
		return c.JSON(400, "Informe o ID do dono ou do animal.")
	}

	filter := Filter{
		OwnerId:  donoId,
		AnimalId: animalId,
	}

	animais, err := getAnimalsByFilter(filter)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, animais)
}

func CreateHandler(c echo.Context) error {

	animal := Animals{}
	c.Bind(&animal)

	animal, err := createAnimals(animal)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, animal)
}

func UpdateHandler(c echo.Context) error {

	animal := Animals{}
	c.Bind(&animal)

	animal, err := updateAnimals(animal)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, animal)
}

func DeleteHandler(c echo.Context) error {

	animalId, _ := strconv.Atoi(c.Param("animalId"))

	err := deleteAnimals(animalId)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, "Animal deletado com sucesso.")
}

func UpdatePesoHandler(c echo.Context) error {

	peso := Peso{}
	c.Bind(&peso)

	err := updatePeso(peso.AnimalId, peso.Peso)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, "Peso atualizado com sucesso.")
}
