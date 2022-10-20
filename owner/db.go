package owner

import "github.com/jmoiron/sqlx"

func getOwnerDBTx(filter Filter, tx *sqlx.Tx) (Owner, error) {
	var owner Owner
	query := `SELECT
    	owner_id,
    	nome,
    	cpf
		FROM owner
		WHERE true`

	args := []interface{}{}

	if filter.Cpf != "" {
		query += ` AND cpf = ?`
		args = append(args, filter.Cpf)
	}

	err := tx.Get(&owner, query, args...)
	if err != nil {
		return Owner{}, err
	}
	return owner, nil
}

func createOwnerDBTx(owner Owner, tx *sqlx.Tx) (Owner, error) {
	query := `INSERT INTO owner (nome, cpf)
		VALUES (?, ?)`

	result, err := tx.Exec(query, owner.Nome, owner.Cpf)
	if err != nil {
		return Owner{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Owner{}, err
	}

	owner.Id = int(id)

	return owner, nil
}

func updateOwnerDBTx(owner Owner, tx *sqlx.Tx) (Owner, error) {
	query := `UPDATE owner
		SET nome = ?, cpf = ?
		WHERE owner_id = ?`

	_, err := tx.Exec(query, owner.Nome, owner.CPF, owner.Id)
	if err != nil {
		return Owner{}, err
	}

	return owner, nil
}

func deleteOwnerDBTx(id int64, tx *sqlx.Tx) error {
	query := `DELETE FROM owner
		WHERE owner_id = ?`

	_, err := tx.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
