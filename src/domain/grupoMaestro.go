package domain

type GrupoMaestro struct {
	id         int64
	maestro    *Maestro
	grupo      *Grupo
	repository GrupoRepository
}

func NewGrupoMaestro(repository GrupoRepository) *GrupoMaestro {
	return &GrupoMaestro{
		repository: repository,
	}
}

func (g *GrupoMaestro) SetID(id int64) {
	g.id = id
}

func (g *GrupoMaestro) GetID() int64 {
	return g.id
}

func (g *GrupoMaestro) SetGrupo(grupo *Grupo) {
	g.grupo = grupo
}

func (g *GrupoMaestro) GetGrupo() *Grupo {
	return g.grupo
}

func (g *GrupoMaestro) SetMaestro(maestro *Maestro) {
	g.maestro = maestro
}

func (g *GrupoMaestro) GetMaestro() *Maestro {
	return g.maestro
}

func (g *GrupoMaestro) Existe() bool {
	return g.id > 0
}
