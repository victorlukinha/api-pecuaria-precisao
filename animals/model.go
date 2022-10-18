package animals

import "google.golang.org/genproto/googleapis/type/date"

type Animals struct {
	Id      int       `db:"id"`
	Especie string    `db:"especie"`
	Raca    string    `db:"raca"`
	DtNasc  date.Date `db:"dt_nasc"`
	OwnerId int       `db:"owner_id"`
}
