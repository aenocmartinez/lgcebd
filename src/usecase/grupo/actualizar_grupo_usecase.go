package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
	"fmt"
	"log"
)

type ActualizarGrupoUseCase struct {
	grupoRepo        domain.GrupoRepository
	celebracionRepo  domain.CelebracionRepository
	cursoPeriodoRepo domain.CursoPeriodoRepository
	maestroRepo      domain.MaestroRepository
}

func NewActualizarGrupoUseCase(
	grupoRepo domain.GrupoRepository,
	celebracionRepo domain.CelebracionRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository,
	maestroRepo domain.MaestroRepository,
) *ActualizarGrupoUseCase {
	return &ActualizarGrupoUseCase{
		grupoRepo:        grupoRepo,
		celebracionRepo:  celebracionRepo,
		cursoPeriodoRepo: cursoPeriodoRepo,
		maestroRepo:      maestroRepo,
	}
}

func (uc *ActualizarGrupoUseCase) Execute(grupoID int64, datos dto.GuardarGrupoDto) shared.APIResponse {

	fmt.Println("Ejecutando ActualizarGrupoUseCase")

	servicio := newGrupoService(uc.grupoRepo, uc.celebracionRepo, uc.cursoPeriodoRepo)

	grupo := uc.grupoRepo.FindByID(grupoID)
	if !grupo.Existe() {
		return shared.NewAPIResponse(404, "Grupo no encontrado", nil)
	}

	grupoConMismaClave := uc.grupoRepo.FindByCursoPeriodoYCelebracion(datos.CursoPeriodoID, datos.CelebracionID)
	if grupoConMismaClave.Existe() && grupoConMismaClave.GetID() != grupoID {
		return shared.NewAPIResponse(409, "Ya existe otro grupo para esta celebración", nil)
	}

	celebracion, cursoPeriodo, errMsg := servicio.validarDatosBasicos(datos)
	if errMsg != nil {
		return shared.NewAPIResponse(404, *errMsg, nil)
	}

	grupo.SetCelebracion(celebracion)
	grupo.SetCursoPeriodo(cursoPeriodo)

	if err := grupo.Actualizar(); err != nil {
		return shared.NewAPIResponse(500, "Ha ocurrido un error en el sistema", nil)
	}

	grupo.QuitarMaestros()
	agregarMaestro := NewAgregarMaestroAGrupoUseCase(uc.grupoRepo, uc.maestroRepo)
	for _, maestroID := range datos.Maestros {
		response := agregarMaestro.Execute(grupo.GetID(), maestroID)
		if response.StatusCode != 201 {
			log.Println(response.Message)
		}
	}

	return shared.NewAPIResponse(200, "El grupo se ha actualizado exitosamente", nil)
}
