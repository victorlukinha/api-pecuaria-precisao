package animals

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func getAnimalsByFilterDBTx(filter Filter, tx *sqlx.Tx) ([]Animals, error) {

	query :=
		`SELECT 
    animal_id,
    especie,
    raca,
    dt_nasc,
    owner_id
    FROM animals
    WHERE true`

	args := []interface{}{}

	if filter.OwnerId != 0 {
		query += ` AND owner_id = ?`
		args = append(args, filter.OwnerId)
	}

	if filter.AnimalId != 0 {
		query += ` AND animal_id = ?`
		args = append(args, filter.AnimalId)
	}

	if filter.Especie != 0 {
		query += ` AND especie = ?`
		args = append(args, filter.Especie)
	}

	var animais []Animals

	err := tx.Select(&animais, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return animais, nil
}

func createAnimalsDBTx(animal Animals, tx *sqlx.Tx) (Animals, error) {

	query :=
		`INSERT INTO animals (especie, raca, dt_nasc, owner_id)
		VALUES (?, ?, ?, ?)`

	result, err := tx.Exec(query, animal.Especie, animal.Raca, animal.DtNasc, animal.OwnerId)
	if err != nil {
		return Animals{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Animals{}, err
	}

	animal.Id = int(id)

	return animal, nil
}

func updateAnimalsDBTx(animal Animals, tx *sqlx.Tx) (Animals, error) {

	query :=
		`UPDATE animals
		SET especie = ?,
		raca = ?,
		dt_nasc = ?,
		owner_id = ?
		WHERE animal_id = ?`

	_, err := tx.Exec(query, animal.Especie, animal.Raca, animal.DtNasc, animal.OwnerId, animal.Id)
	if err != nil {
		return Animals{}, err
	}

	return animal, nil
}

func deleteAnimalsDBTx(animalId int, tx *sqlx.Tx) error {

	query := `DELETE FROM animals WHERE animal_id = ?`

	_, err := tx.Exec(query, animalId)
	if err != nil {
		return err
	}

	return nil
}
