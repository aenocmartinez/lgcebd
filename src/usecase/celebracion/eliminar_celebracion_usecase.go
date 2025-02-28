package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"log"
)

type EliminarCelebracionUseCase struct {
	celebracionRepo domain.CelebracionRepository
}

func NewEliminarCelebracionUseCase(repository domain.CelebracionRepository) *EliminarCelebracionUseCase {
	return &EliminarCelebracionUseCase{
		celebracionRepo: repository,
	}
}

func (uc *EliminarCelebracionUseCase) Execute(id int64) shared.APIResponse {

	celebracion := uc.celebracionRepo.FindByID(id)
	if !celebracion.Existe() {
		return shared.NewAPIResponse(404, "Celebración no encontrada", nil)
	}

	err := celebracion.Eliminar()
	if err != nil {
		log.Println(err)
		return shared.NewAPIResponse(500, "Ha ocurrido un error en el sistema", nil)
	}

	return shared.NewAPIResponse(200, "La celebración se ha eliminado exitosamente.", nil)
}
