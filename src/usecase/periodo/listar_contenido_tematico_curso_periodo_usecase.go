package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type ListarContenidoTematicoDeUnCursoPeriodoUseCase struct {
	contenidoTematicoRepo domain.ContenidoTematicoRepository
	cursoPeriodoRepo      domain.CursoPeriodoRepository
}

func NewListarContenidoTematicoDeUnCursoPeriodoUseCase(
	contenidoTematicoRepo domain.ContenidoTematicoRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository) *ListarContenidoTematicoDeUnCursoPeriodoUseCase {
	return &ListarContenidoTematicoDeUnCursoPeriodoUseCase{
		contenidoTematicoRepo: contenidoTematicoRepo,
		cursoPeriodoRepo:      cursoPeriodoRepo,
	}
}

func (uc *ListarContenidoTematicoDeUnCursoPeriodoUseCase) Execute(cursoPeriodoID int64) shared.APIResponse {

	cursoPeriodo := uc.cursoPeriodoRepo.FindByID(cursoPeriodoID)
	if !cursoPeriodo.Existe() {
		return shared.NewAPIResponse(404, "Curso periodo no encontrado", nil)
	}

	contenidos := []dto.ItemConteniodoTematicoDTO{}
	for _, contenido := range cursoPeriodo.ContenidoTematico() {

		contenidos = append(contenidos, dto.ItemConteniodoTematicoDTO{
			ID:          contenido.GetID(),
			Orden:       contenido.GetOrden(),
			Descripcion: contenido.GetDescripcion(),
		})
	}

	return shared.NewAPIResponse(200, "Contenido temático obtenidos correctamente", contenidos)
}
