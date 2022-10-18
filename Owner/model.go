package Owner

type Owner struct {
	Id   int    `db:"id"`
	Nome string `db:"nome"`
	CPF  string `db:"cpf"`
}
