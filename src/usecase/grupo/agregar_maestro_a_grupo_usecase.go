package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"log"
)

type AgregarMaestroAGrupoUseCase struct {
	grupoRepo   domain.GrupoRepository
	maestroRepo domain.MaestroRepository
}

func NewAgregarMaestroAGrupoUseCase(grupoRepo domain.GrupoRepository, maestroRepo domain.MaestroRepository) *AgregarMaestroAGrupoUseCase {
	return &AgregarMaestroAGrupoUseCase{
		grupoRepo:   grupoRepo,
		maestroRepo: maestroRepo,
	}
}

func (uc *AgregarMaestroAGrupoUseCase) Execute(grupoID int64, maestroID int64) shared.APIResponse {

	grupo := uc.grupoRepo.FindByID(grupoID)
	if !grupo.Existe() {
		return shared.NewAPIResponse(404, "Grupo no encontrado", nil)
	}

	maestro := uc.maestroRepo.FindByID(maestroID)
	if !maestro.Existe() {
		return shared.NewAPIResponse(404, "Maestro no encontrado", nil)
	}

	err := grupo.AgregarMaestro(maestro)
	if err != nil {
		log.Println(err)
		return shared.NewAPIResponse(500, "Ha ocurrido un error en el sistema", nil)
	}

	return shared.NewAPIResponse(201, "Se ha agregado el maestro al grupo con éxito", nil)

}
