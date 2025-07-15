package usecase

import (
	"ebd/src/domain"
	"ebd/src/view/dto"
)

type grupoService struct {
	grupoRepo        domain.GrupoRepository
	celebracionRepo  domain.CelebracionRepository
	cursoPeriodoRepo domain.CursoPeriodoRepository
}

func newGrupoService(gr domain.GrupoRepository, cr domain.CelebracionRepository, cpr domain.CursoPeriodoRepository) *grupoService {
	return &grupoService{
		grupoRepo:        gr,
		celebracionRepo:  cr,
		cursoPeriodoRepo: cpr,
	}
}

func (s *grupoService) validarDatosBasicos(dto dto.GuardarGrupoDto) (*domain.Celebracion, *domain.CursoPeriodo, *string) {
	celebracion := s.celebracionRepo.FindByID(dto.CelebracionID)
	if !celebracion.Existe() {
		msg := "Celebración no encontrada"
		return nil, nil, &msg
	}

	curso := s.cursoPeriodoRepo.FindByID(dto.CursoPeriodoID)
	if !curso.Existe() {
		msg := "Curso no encontrado"
		return celebracion, nil, &msg
	}

	return celebracion, curso, nil
}
