package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"log"
)

type CrearCelebracionUseCase struct {
	celebracionRepo domain.CelebracionRepository
}

func NewCrearCelabracionUseCase(celebracionRepo domain.CelebracionRepository) *CrearCelebracionUseCase {
	return &CrearCelebracionUseCase{
		celebracionRepo: celebracionRepo,
	}
}

func (uc *CrearCelebracionUseCase) Execute(nombre string) shared.APIResponse {
	celebracion := uc.celebracionRepo.FindByNombre(nombre)
	if celebracion.Existe() {
		return shared.NewAPIResponse(409, "Ya existe una celebración con este nombre", nil)
	}

	celebracion = domain.NewCelebracion(uc.celebracionRepo)
	celebracion.SetNombre(nombre)

	err := celebracion.Crear()
	if err != nil {
		log.Println(err)
		return shared.NewAPIResponse(500, "Ha ocurrido un error en el sistema", nil)
	}

	return shared.NewAPIResponse(201, "Celebración creada exitosamente", celebracion.ToDTO())
}
