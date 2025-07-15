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

func ActualizarGrupo(c *gin.Context) {
	var req formrequest.GrupoFormRequest

	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada no válidos", nil))
		return
	}

	datos := dto.GuardarGrupoDto{
		CelebracionID:  req.CelebracionID,
		CursoPeriodoID: req.CursoPeriodoID,
		Maestros:       req.Maestros,
	}

	actualizarGrupo := usecase.NewActualizarGrupoUseCase(di.GetContainer().GetGrupoRepository(),
		di.GetContainer().GetCelebracionRepository(),
		di.GetContainer().GetCursoPeriodoRepository(),
		di.GetContainer().GetMaestroRepository())

	response := actualizarGrupo.Execute(id, datos)

	c.JSON(response.StatusCode, response)

}

func EliminarGrupo(c *gin.Context) {
	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	eliminarGrupo := usecase.NewEliminarGrupoUseCase(di.GetContainer().GetGrupoRepository())

	response := eliminarGrupo.Execute(id)
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

func GuardarGrupo(c *gin.Context) {
	var req formrequest.GrupoFormRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada no válidos", nil))
		return
	}

	datos := dto.GuardarGrupoDto{
		CelebracionID:  req.CelebracionID,
		CursoPeriodoID: req.CursoPeriodoID,
		Maestros:       req.Maestros,
	}

	guardarGrupo := usecase.NewGuardarGrupoUseCase(
		di.GetContainer().GetGrupoRepository(),
		di.GetContainer().GetCelebracionRepository(),
		di.GetContainer().GetCursoPeriodoRepository(),
		di.GetContainer().GetMaestroRepository(),
	)

	response := guardarGrupo.Execute(req.ID, datos)

	c.JSON(response.StatusCode, response)
}
