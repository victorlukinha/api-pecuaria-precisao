package animals

import (
	"github.com/eucatur/go-toolbox/database"
	"github.com/victorlukinha/api-pecuaria-precisao/config"
)

func getAnimalsByFilter(filter Filter) ([]Animals, error) {
	return getAnimalsByFilterTx(filter)
}

func getAnimalsByFilterTx(filter Filter) ([]Animals, error) {

	tx := database.MustGetByFile(config.MYSQL_ENV).MustBegin()
	defer tx.Rollback()

	return getAnimalsByFilterDBTx(filter, tx)

}

func createAnimals(animal Animals) (Animals, error) {

	tx := database.MustGetByFile(config.MYSQL_ENV).MustBegin()

	animal, err := createAnimalsDBTx(animal, tx)
	if err != nil {
		tx.Rollback()
		return Animals{}, err
	}

	tx.Commit()

	return animal, nil

}
