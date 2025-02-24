package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type EliminarPeriodoUseCase struct {
	repo domain.PeriodoRepository
}

func NewEliminarPeriodoUseCase(repo domain.PeriodoRepository) *EliminarPeriodoUseCase {
	return &EliminarPeriodoUseCase{repo: repo}
}

func (u *EliminarPeriodoUseCase) Execute(id int64) shared.APIResponse {

	periodo, err := u.repo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el periodo", nil)
	}

	if !periodo.Existe() {
		return shared.NewAPIResponse(404, "El periodo no existe", nil)
	}

	if err := u.repo.Delete(id); err != nil {
		return shared.NewAPIResponse(500, "Error al eliminar el periodo", nil)
	}

	return shared.NewAPIResponse(200, "Periodo eliminado exitosamente", nil)
}
