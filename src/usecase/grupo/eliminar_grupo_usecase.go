package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type EliminarGrupoUseCase struct {
	grupoRepo domain.GrupoRepository
}

func NewEliminarGrupoUseCase(grupoRepo domain.GrupoRepository) *EliminarGrupoUseCase {
	return &EliminarGrupoUseCase{
		grupoRepo: grupoRepo,
	}
}

func (uc *EliminarGrupoUseCase) Execute(grupoID int64) shared.APIResponse {

	grupo := uc.grupoRepo.FindByID(grupoID)
	if !grupo.Existe() {
		return shared.NewAPIResponse(404, "Grupo no encontrado", nil)
	}

	grupo.QuitarMaestros()
	err := grupo.Eliminar()
	if err != nil {
		return shared.NewAPIResponse(500, "Ha ocurrido un error en el sistema.", nil)
	}

	return shared.NewAPIResponse(200, "El grupo se ha eliminado exitosamente.", nil)
}
