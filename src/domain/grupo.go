package domain

import "ebd/src/view/dto"

type Grupo struct {
	id               int64
	celebracion      *Celebracion
	cursoPeriodo     *CursoPeriodo
	repository       GrupoRepository
	celebracionRepo  CelebracionRepository
	cursoPeriodoRepo CursoPeriodoRepository
}

func NewGrupo(repository GrupoRepository) *Grupo {
	return &Grupo{
		repository: repository,
	}
}

func (g *Grupo) SetCelebracionRepository(celebracionRepo CelebracionRepository) {
	g.celebracionRepo = celebracionRepo
}

func (g *Grupo) SetCursoPeriodoRepository(cursoPeriodoRepo CursoPeriodoRepository) {
	g.cursoPeriodoRepo = cursoPeriodoRepo
}

func (g *Grupo) SetID(id int64) {
	g.id = id
}

func (g *Grupo) SetCelebracion(celebracion *Celebracion) {
	g.celebracion = celebracion
}

func (g *Grupo) SetCursoPeriodo(cursoPeriodo *CursoPeriodo) {
	g.cursoPeriodo = cursoPeriodo
}

func (g *Grupo) GetID() int64 {
	return g.id
}

func (g *Grupo) GetCelebracion() *Celebracion {
	return g.celebracion
}

func (g *Grupo) GetCursoPeriodo() *CursoPeriodo {
	return g.cursoPeriodo
}

func (g *Grupo) Existe() bool {
	return g.id > 0
}

func (g *Grupo) Crear() error {
	return g.repository.Save(g)
}

func (g *Grupo) Actualizar() error {
	return g.repository.Update(g)
}

func (g *Grupo) Eliminar() error {
	return g.repository.Delete(g.id)
}

func (g *Grupo) AgregarMaestro(maestro *Maestro) error {
	return g.repository.AgregarMaestro(g.id, maestro.GetID())
}

func (g *Grupo) QuitarMaestros() error {
	return g.repository.QuitarMaestros(g.id)
}

func (g *Grupo) Maestros() []GrupoMaestro {
	return g.repository.ListarMaestros(g.id)
}

func (g *Grupo) ToDTO() dto.GrupoDto {

	return dto.GrupoDto{
		ID:           g.id,
		CursoPeriodo: g.GetCursoPeriodo().ToDTO(),
		Celebracion:  g.celebracion.ToDTO(),
	}
}
