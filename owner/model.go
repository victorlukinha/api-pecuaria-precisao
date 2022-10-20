package owner

type Owner struct {
	Id   int    `db:"id"`
	Nome string `db:"nome"`
	CPF  string `db:"cpf"`
}

type Filter struct {
	Nome string
	Cpf  string
}
