package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type EliminarMaestroUseCase struct {
	maestroRepo domain.MaestroRepository
}

func NewEliminarMaestroUseCase(maestroRepo domain.MaestroRepository) *EliminarMaestroUseCase {
	return &EliminarMaestroUseCase{maestroRepo: maestroRepo}
}

func (u *EliminarMaestroUseCase) Execute(id int64) shared.APIResponse {
	maestro := u.maestroRepo.FindByID(id)
	if !maestro.Existe() {
		return shared.NewAPIResponse(500, "Maestro no encontrado", nil)
	}

	err := u.maestroRepo.Delete(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al eliminar el maestro", nil)
	}

	return shared.NewAPIResponse(200, "Maestro eliminado correctamente", nil)
}
