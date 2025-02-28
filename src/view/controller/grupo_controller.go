package controller

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
	usecase "ebd/src/usecase/grupo"
	"ebd/src/view/dto"
	formrequest "ebd/src/view/formrequest/grupo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrearGrupo(c *gin.Context) {
	var req formrequest.GrupoFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada no válidos", nil))
		return
	}

	datos := dto.GuardarGrupoDto{
		CelebracionID:  req.CelebracionID,
		CursoPeriodoID: req.CursoPeriodoID,
		Maestros:       req.Maestros,
	}

	crearGrupo := usecase.NewCrearGrupoUseCase(di.GetContainer().GetGrupoRepository(),
		di.GetContainer().GetCelebracionRepository(),
		di.GetContainer().GetCursoPeriodoRepository(),
		di.GetContainer().GetMaestroRepository())

	response := crearGrupo.Execute(datos)

	c.JSON(response.StatusCode, response)

}

func ListarGrupos(c *gin.Context) {
	listarGrupos := usecase.NewListarGruposUseCase(di.GetContainer().GetGrupoRepository(),
		di.GetContainer().GetCelebracionRepository(),
		di.GetContainer().GetCursoPeriodoRepository(),
		di.GetContainer().GetMaestroRepository())
	response := listarGrupos.Execute()
	c.JSON(response.StatusCode, response)
}
