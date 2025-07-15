package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type GuardarGrupoUseCase struct {
	grupoRepo        domain.GrupoRepository
	celebracionRepo  domain.CelebracionRepository
	cursoPeriodoRepo domain.CursoPeriodoRepository
	maestroRepo      domain.MaestroRepository
}

func NewGuardarGrupoUseCase(
	grupoRepo domain.GrupoRepository,
	celebracionRepo domain.CelebracionRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository,
	maestroRepo domain.MaestroRepository,
) *GuardarGrupoUseCase {
	return &GuardarGrupoUseCase{
		grupoRepo:        grupoRepo,
		celebracionRepo:  celebracionRepo,
		cursoPeriodoRepo: cursoPeriodoRepo,
		maestroRepo:      maestroRepo,
	}
}

func (uc *GuardarGrupoUseCase) Execute(grupoID *int64, datos dto.GuardarGrupoDto) shared.APIResponse {

	if grupoID == nil {
		crearUseCase := NewCrearGrupoUseCase(
			uc.grupoRepo,
			uc.celebracionRepo,
			uc.cursoPeriodoRepo,
			uc.maestroRepo,
		)
		return crearUseCase.Execute(datos)
	}

	actualizarUseCase := NewActualizarGrupoUseCase(
		uc.grupoRepo,
		uc.celebracionRepo,
		uc.cursoPeriodoRepo,
		uc.maestroRepo,
	)

	return actualizarUseCase.Execute(*grupoID, datos)
}
