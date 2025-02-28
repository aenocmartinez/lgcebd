package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
	"log"
)

type ActualizarGrupoUseCase struct {
	grupoRepo        domain.GrupoRepository
	celebracionRepo  domain.CelebracionRepository
	cursoPeriodoRepo domain.CursoPeriodoRepository
	maestroRepo      domain.MaestroRepository
}

func NewActualizarGrupoUseCase(grupoRepo domain.GrupoRepository,
	celebracionRepo domain.CelebracionRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository,
	maestroRepo domain.MaestroRepository) *ActualizarGrupoUseCase {
	return &ActualizarGrupoUseCase{
		grupoRepo:        grupoRepo,
		celebracionRepo:  celebracionRepo,
		cursoPeriodoRepo: cursoPeriodoRepo,
		maestroRepo:      maestroRepo,
	}
}

func (uc *ActualizarGrupoUseCase) Execute(grupoID int64, datos dto.GuardarGrupoDto) shared.APIResponse {

	grupo := uc.grupoRepo.FindByCursoPeriodoYCelebracion(datos.CursoPeriodoID, datos.CelebracionID)
	if grupo.Existe() {
		return shared.NewAPIResponse(409, "Ya existe un grupo para esta celebración", nil)
	}

	grupo = uc.grupoRepo.FindByID(grupoID)
	if !grupo.Existe() {
		return shared.NewAPIResponse(404, "Grupo no encontrado", nil)
	}

	celebracion := uc.celebracionRepo.FindByID(datos.CelebracionID)
	if !celebracion.Existe() {
		return shared.NewAPIResponse(404, "Celebración no encontrada", nil)
	}

	cursoPeriodo := uc.cursoPeriodoRepo.FindByID(datos.CursoPeriodoID)
	if !cursoPeriodo.Existe() {
		return shared.NewAPIResponse(404, "Curso no encontrado", nil)
	}

	grupo.SetCelebracion(celebracion)
	grupo.SetCursoPeriodo(cursoPeriodo)

	err := grupo.Actualizar()
	if err != nil {
		return shared.NewAPIResponse(500, "Ha ocurrido un error en el sistema", nil)
	}

	agregarMaestro := NewAgregarMaestroAGrupoUseCase(uc.grupoRepo, uc.maestroRepo)
	grupo.QuitarMaestros()
	for _, maestroID := range datos.Maestros {
		response := agregarMaestro.Execute(grupo.GetID(), maestroID)
		if response.StatusCode != 201 {
			log.Println(response.Message)
		}
	}

	return shared.NewAPIResponse(200, "El grupo se ha actualizado con exitosamente", nil)
}
