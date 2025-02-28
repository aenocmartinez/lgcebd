package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type AgregarContenidoTematicoUseCase struct {
	contenidoTematicoRepo domain.ContenidoTematicoRepository
	cursoPeriodoRepo      domain.CursoPeriodoRepository
}

func NewAgregarContenidoTematicoUseCase(
	contenidoTematicoRepo domain.ContenidoTematicoRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository) *AgregarContenidoTematicoUseCase {
	return &AgregarContenidoTematicoUseCase{
		contenidoTematicoRepo: contenidoTematicoRepo,
		cursoPeriodoRepo:      cursoPeriodoRepo,
	}
}

func (uc *AgregarContenidoTematicoUseCase) Execute(cursoPeriodoID int64, descripcion string) shared.APIResponse {

	contenidoTematico := uc.contenidoTematicoRepo.FindByDescripcion(cursoPeriodoID, descripcion)
	if contenidoTematico.Existe() {
		return shared.NewAPIResponse(409, "El contenido temático ya existe", nil)
	}

	cursoPeriodo := uc.cursoPeriodoRepo.FindByID(cursoPeriodoID)
	if !cursoPeriodo.Existe() {
		return shared.NewAPIResponse(404, "Curso no encontrado", nil)
	}

	err := cursoPeriodo.AgregarContenidoTematico(descripcion)
	if err != nil {
		return shared.NewAPIResponse(500, "Ha ocurrido un error en el sistema", nil)
	}

	return shared.NewAPIResponse(201, "Contenido temático creado exitosamente", nil)
}
