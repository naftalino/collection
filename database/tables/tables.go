package tables

type User struct {
	ID               string
	Fichas           int
	Moedas           int
	Banido           bool
	QntCartas        int
	SpinsDisponiveis int
	ContaCriada      string
	Admin            bool
}

type Category struct {
}

type SubCategory struct {
}

type Collection struct {
	UserID       string
	QntAcumulada int
	IdCarta      int
}

type Card struct {
}

type Serie struct {
}

type Shop struct {
}

type Events struct {
}

type RankingBiggestCollection struct {
}

type Admins struct {
}

type Owner struct {
}

type Gifts struct {
}
