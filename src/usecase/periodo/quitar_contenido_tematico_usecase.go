package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type QuitarContenidoTematicoUseCase struct {
	contenidoTematicoRepo domain.ContenidoTematicoRepository
	cursoPeriodoRepo      domain.CursoPeriodoRepository
}

func NewQuitarContenidoTematicoUseCase(
	contenidoTematicoRepo domain.ContenidoTematicoRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository) *QuitarContenidoTematicoUseCase {
	return &QuitarContenidoTematicoUseCase{
		contenidoTematicoRepo: contenidoTematicoRepo,
		cursoPeriodoRepo:      cursoPeriodoRepo,
	}
}

func (uc *QuitarContenidoTematicoUseCase) Execute(cursoPeriodoID int64, contenidoTematicoID int64) shared.APIResponse {

	cursoPeriodo := uc.cursoPeriodoRepo.FindByID(cursoPeriodoID)
	if !cursoPeriodo.Existe() {
		return shared.NewAPIResponse(404, "Curso periodo no encontrado", nil)
	}

	contenidoTematico := uc.contenidoTematicoRepo.FindByID(contenidoTematicoID)
	if !contenidoTematico.Existe() {
		return shared.NewAPIResponse(409, "El contenido temático no encontrado", nil)
	}

	err := cursoPeriodo.QuitarContenidoTematico(contenidoTematico)
	if err != nil {
		return shared.NewAPIResponse(500, "Ha ocurrido un error en el sistema", nil)
	}

	return shared.NewAPIResponse(201, "Contenido temático eliminado exitosamente", nil)
}
