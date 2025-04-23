package postgres

type PostgresDb struct {
}

func New() *PostgresDb {
	return &PostgresDb{}
}

func (p *PostgresDb) Get()    {}
func (p *PostgresDb) Update() {}
