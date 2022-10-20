package owner

import (
	"github.com/eucatur/go-toolbox/database"
	"github.com/victorlukinha/api-pecuaria-precisao/config"
)

func getOwner(filter Filter) (Owner, error) {
	tx := database.MustGetByFile(config.MYSQL_ENV).MustBegin()
	defer tx.Rollback()

	return getOwnerDBTx(filter, tx)

}

func createOwner(owner Owner) (Owner, error) {

	tx := database.MustGetByFile(config.MYSQL_ENV).MustBegin()

	owner, err := createOwnerDBTx(owner, tx)
	if err != nil {
		tx.Rollback()
		return Owner{}, err
	}

	tx.Commit()

	return owner, nil

}

func updateOwner(owner Owner) (Owner, error) {

	tx := database.MustGetByFile(config.MYSQL_ENV).MustBegin()

	owner, err := updateOwnerDBTx(owner, tx)
	if err != nil {
		tx.Rollback()
		return Owner{}, err
	}

	tx.Commit()

	return owner, nil

}

func deleteOwner(ownerId int64) error {

	tx := database.MustGetByFile(config.MYSQL_ENV).MustBegin()

	err := deleteOwnerDBTx(ownerId, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil

}
