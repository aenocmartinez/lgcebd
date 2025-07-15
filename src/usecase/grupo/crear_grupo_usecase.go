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

func NewCrearGrupoUseCase(
	repoGrupo domain.GrupoRepository,
	celebracionRepo domain.CelebracionRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository,
	maestroRepo domain.MaestroRepository,
) *CrearGrupoUseCase {
	return &CrearGrupoUseCase{
		grupoRepo:        repoGrupo,
		celebracionRepo:  celebracionRepo,
		cursoPeriodoRepo: cursoPeriodoRepo,
		maestroRepo:      maestroRepo,
	}
}

func (uc *CrearGrupoUseCase) Execute(datos dto.GuardarGrupoDto) shared.APIResponse {

	servicio := newGrupoService(uc.grupoRepo, uc.celebracionRepo, uc.cursoPeriodoRepo)

	grupoExistente := uc.grupoRepo.FindByCursoPeriodoYCelebracion(datos.CursoPeriodoID, datos.CelebracionID)
	if grupoExistente.Existe() {
		return shared.NewAPIResponse(409, "Ya existe un grupo para esta celebración", nil)
	}

	celebracion, cursoPeriodo, errMsg := servicio.validarDatosBasicos(datos)
	if errMsg != nil {
		return shared.NewAPIResponse(404, *errMsg, nil)
	}

	grupo := domain.NewGrupo(uc.grupoRepo)
	grupo.SetCelebracion(celebracion)
	grupo.SetCursoPeriodo(cursoPeriodo)

	if err := grupo.Crear(); err != nil {
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
