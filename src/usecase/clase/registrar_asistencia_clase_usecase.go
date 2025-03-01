package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
	"log"
)

type RegistrarAsistenciaUseCase struct {
	claseRepo             domain.ClaseRepository
	grupoRepo             domain.GrupoRepository
	contenidoTematicoRepo domain.ContenidoTematicoRepository
	matriculaRepo         domain.MatriculaRepository
}

func NewRegistrarAsistenciaUseCase(claseRepo domain.ClaseRepository,
	grupoRepo domain.GrupoRepository,
	conteniodoTematicoRepo domain.ContenidoTematicoRepository,
	matriculaRepo domain.MatriculaRepository) *RegistrarAsistenciaUseCase {
	return &RegistrarAsistenciaUseCase{
		claseRepo:             claseRepo,
		grupoRepo:             grupoRepo,
		contenidoTematicoRepo: conteniodoTematicoRepo,
		matriculaRepo:         matriculaRepo,
	}
}

func (uc *RegistrarAsistenciaUseCase) Execute(datos dto.GuardarClaseDTO) shared.APIResponse {

	clase := uc.claseRepo.FindByGrupoFecha(datos.GrupoID, datos.Fecha)
	if clase.Existe() {
		return shared.NewAPIResponse(409, "La clase para este grupo ya fue registrada", nil)
	}

	grupo := uc.grupoRepo.FindByID(datos.GrupoID)
	if !grupo.Existe() {
		return shared.NewAPIResponse(404, "Grupo no encontrado", nil)
	}

	contenidoTematico := uc.contenidoTematicoRepo.FindByID(datos.ContenidoTematicoID)
	if !contenidoTematico.Existe() {
		return shared.NewAPIResponse(404, "Contenido temático no encontrado", nil)
	}

	clase.SetFecha(datos.Fecha)
	clase.SetOfrenda(datos.Ofrenda)
	clase.SetGrupo(grupo)
	clase.SetContenidoTematico(contenidoTematico)

	err := clase.Crear()
	if err != nil {
		return shared.NewAPIResponse(500, "Ha ocurrido un error en el sistema", err.Error())
	}

	for _, alumnoMatriculadoID := range datos.AlumnosMatriculados {
		matricula, err := uc.matriculaRepo.FindByID(alumnoMatriculadoID)
		if err != nil {
			log.Println("Error al consultar matrícula alumno: ", err.Error())
			continue
		}

		err = clase.RegistrarAsistencia(matricula)
		if err != nil {
			log.Println("Error al agrega la asistencia: ", err.Error())
			continue
		}
	}

	return shared.NewAPIResponse(201, "Clase registrada correctamente.", nil)
}
