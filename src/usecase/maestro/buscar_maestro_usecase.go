package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type BuscarMaestroUseCase struct {
	maestroRepo domain.MaestroRepository
}

func NewBuscarMaestroUseCase(maestroRepo domain.MaestroRepository) *BuscarMaestroUseCase {
	return &BuscarMaestroUseCase{maestroRepo: maestroRepo}
}

func (u *BuscarMaestroUseCase) Execute(id int64) shared.APIResponse {
	maestro := u.maestroRepo.FindByID(id)
	if !maestro.Existe() {
		return shared.NewAPIResponse(500, "Maestro no encontrado", nil)
	}

	return shared.NewAPIResponse(200, "Maestro encontrado", maestro.ToDTO())
}
