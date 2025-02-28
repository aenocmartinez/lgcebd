package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type ActualizarCelebracionUseCase struct {
	celeracionRepo domain.CelebracionRepository
}

func NewActualizarCelebracionUseCase(repository domain.CelebracionRepository) *ActualizarCelebracionUseCase {
	return &ActualizarCelebracionUseCase{
		celeracionRepo: repository,
	}
}

func (uc *ActualizarCelebracionUseCase) Execute(id int64, nombre string) shared.APIResponse {

	celebracion := uc.celeracionRepo.FindByID(id)
	if !celebracion.Existe() {
		return shared.NewAPIResponse(404, "Celebración no encontrado", nil)
	}

	celebracion.SetNombre(nombre)
	err := celebracion.Actualizar()
	if err != nil {
		return shared.NewAPIResponse(500, "Ha ocurrido un error en el sistema.", nil)
	}

	return shared.NewAPIResponse(200, "La celebración se ha actualizado exitosamente.", nil)
}
