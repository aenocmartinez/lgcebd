package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type ListarGruposUseCase struct {
	grupoRepo        domain.GrupoRepository
	celebracionRepo  domain.CelebracionRepository
	cursoPeriodoRepo domain.CursoPeriodoRepository
	maestroRepo      domain.MaestroRepository
}

func NewListarGruposUseCase(grupoRepo domain.GrupoRepository,
	celebracionRepo domain.CelebracionRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository,
	maestroRepo domain.MaestroRepository) *ListarGruposUseCase {
	return &ListarGruposUseCase{
		grupoRepo:        grupoRepo,
		celebracionRepo:  celebracionRepo,
		cursoPeriodoRepo: cursoPeriodoRepo,
		maestroRepo:      maestroRepo,
	}
}

func (uc *ListarGruposUseCase) Execute() shared.APIResponse {

	grupos := uc.grupoRepo.List()

	lista := []dto.GrupoDto{}
	for _, grupo := range grupos {
		maestros := []dto.MaestroDTO{}
		for _, grupoMaestro := range grupo.Maestros() {
			maestros = append(maestros, *grupoMaestro.GetMaestro().ToDTO())
		}

		grupoDTO := grupo.ToDTO()
		grupoDTO.Maestros = maestros

		lista = append(lista, grupoDTO)
	}

	return shared.NewAPIResponse(200, "Grupos obtenidos correctamente", lista)
}
