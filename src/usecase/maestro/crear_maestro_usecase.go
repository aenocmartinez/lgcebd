package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type CrearMaestroUseCase struct {
	maestroRepo domain.MaestroRepository
}

func NewCrearMaestroUseCase(maestroRepo domain.MaestroRepository) *CrearMaestroUseCase {
	return &CrearMaestroUseCase{maestroRepo: maestroRepo}
}

func (u *CrearMaestroUseCase) Execute(maestroDTO dto.MaestroDTO) shared.APIResponse {

	maestro := domain.NewMaestro(u.maestroRepo)
	maestro.SetNombre(maestroDTO.Nombre)
	maestro.SetTelefono(maestroDTO.Telefono)
	maestro.SetFechaNacimiento(maestroDTO.FechaNacimiento)
	maestro.SetEstado("activo")

	err := u.maestroRepo.Save(maestro)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al crear el maestro", nil)
	}

	return shared.NewAPIResponse(201, "Maestro creado correctamente", maestro.ToDTO())
}
