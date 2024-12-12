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
	ID           int
	Nome         int
	ObrasLigadas []int
}

type Collection struct {
	UserID       string
	QntAcumulada int
	IdCarta      int
}

type Card struct {
	ID       int
	Foto     string
	Nome     string
	IdObra   int
	Raridade int // 1: comum, 2:raro, 3:lendário
}

type Serie struct {
	ID        int
	Foto      string
	Nome      string
	QntCartas int
	Descricao string
}

type Shop struct {
	ItemTipo   string
	Valor      int
	Disponivel bool
}

type Events struct {
	Nome      string
	Descricao string
	Presentes []int
}

type Admins struct {
	UserID string
}

type Owner struct {
	UserID string
}

type Reward struct {
	Presente string
}
