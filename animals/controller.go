package animals

func getAnimalsByFilter(filter Filter) ([]Animals, error) {
	return getAnimalsByFilterTx(filter)
}

func getAnimalsByFilterTx(filter Filter) ([]Animals, error) {

	tx := db.MustBegin()
	defer tx.Rollback()

	return getAnimalsByFilterDBTx(filter, tx)

}

func createAnimals(animal Animals) (Animals, error) {

	tx := db.MustBegin()

	animal, err := createAnimalsDBTx(animal, tx)
	if err != nil {
		tx.Rollback()
		return Animals{}, err
	}

	tx.Commit()

	return animal, nil

}
