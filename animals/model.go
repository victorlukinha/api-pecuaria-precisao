package animals

import "google.golang.org/genproto/googleapis/type/date"

type Animals struct {
	Id      int       `db:"id" json:"id"`
	Especie string    `db:"especie" json:"especie"`
	Raca    string    `db:"raca" json:"raca"`
	DtNasc  date.Date `db:"dt_nasc" json:"dtNasc"`
	OwnerId int       `db:"owner_id" json:"ownerId"`
	Peso    float64   `db:"peso" json:"peso"`
}

type Filter struct {
	OwnerId  int `json:"owner_id"`
	AnimalId int `json:"animal_id"`
	Especie  int `json:"especie"`
}

type Peso struct {
	AnimalId int     `db:"animal_id" json:"animalId"`
	Peso     float64 `db:"peso" json:"peso"`
}
