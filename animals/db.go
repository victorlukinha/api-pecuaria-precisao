package animals

import "github.com/jmoiron/sqlx"

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
	if err != nil && err != sqlx.ErrNoRows {
		return nil, err
	}

	return animais, nil
}
