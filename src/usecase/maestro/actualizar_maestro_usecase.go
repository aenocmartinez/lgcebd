package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type ActualizarMaestroUseCase struct {
	maestroRepo domain.MaestroRepository
}

func NewActualizarMaestroUseCase(maestroRepo domain.MaestroRepository) *ActualizarMaestroUseCase {
	return &ActualizarMaestroUseCase{maestroRepo: maestroRepo}
}

func (u *ActualizarMaestroUseCase) Execute(id int64, maestroDTO dto.MaestroDTO) shared.APIResponse {

	maestro, err := u.maestroRepo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el maestro", nil)
	}

	if !maestro.Existe() {
		return shared.NewAPIResponse(404, "El maestro no existe", nil)
	}

	maestro.SetNombre(maestroDTO.Nombre)
	maestro.SetTelefono(maestroDTO.Telefono)
	maestro.SetFechaNacimiento(maestroDTO.FechaNacimiento)
	maestro.SetEstado(maestroDTO.Estado)

	err = u.maestroRepo.Update(maestro)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al actualizar el maestro", nil)
	}

	return shared.NewAPIResponse(200, "Maestro actualizado correctamente", maestro.ToDTO())
}
