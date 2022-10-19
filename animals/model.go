package animals

import "google.golang.org/genproto/googleapis/type/date"

type Animals struct {
	Id      int       `db:"id"`
	Especie string    `db:"especie"`
	Raca    string    `db:"raca"`
	DtNasc  date.Date `db:"dt_nasc"`
	OwnerId int       `db:"owner_id"`
}

type Filter struct {
	OwnerId  int `json:"owner_id"`
	AnimalId int `json:"animal_id"`
	Especie  int `json:"especie"`
}
