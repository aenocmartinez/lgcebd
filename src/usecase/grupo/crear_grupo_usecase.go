package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
	"log"
)

type CrearGrupoUseCase struct {
	grupoRepo        domain.GrupoRepository
	celebracionRepo  domain.CelebracionRepository
	cursoPeriodoRepo domain.CursoPeriodoRepository
	maestroRepo      domain.MaestroRepository
}

func NewCrearGrupoUseCase(repoGrupo domain.GrupoRepository,
	celebracionRepo domain.CelebracionRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository,
	maestroRepo domain.MaestroRepository) *CrearGrupoUseCase {
	return &CrearGrupoUseCase{
		grupoRepo:        repoGrupo,
		celebracionRepo:  celebracionRepo,
		cursoPeriodoRepo: cursoPeriodoRepo,
		maestroRepo:      maestroRepo,
	}
}

func (uc *CrearGrupoUseCase) Execute(datos dto.GuardarGrupoDto) shared.APIResponse {

	grupo := uc.grupoRepo.FindByCursoPeriodoYCelebracion(datos.CursoPeriodoID, datos.CelebracionID)
	if grupo.Existe() {
		return shared.NewAPIResponse(409, "Ya existe un grupo para esta celebración", nil)
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

	err := grupo.Crear()
	if err != nil {
		return shared.NewAPIResponse(500, "Ha ocurrido un error en el sistema", nil)
	}

	agregarMaestro := NewAgregarMaestroAGrupoUseCase(uc.grupoRepo, uc.maestroRepo)
	for _, maestroID := range datos.Maestros {
		response := agregarMaestro.Execute(grupo.GetID(), maestroID)
		if response.StatusCode != 201 {
			log.Println(response.Message)
		}
	}

	return shared.NewAPIResponse(201, "El grupo se ha creado exitosamente", nil)
}
