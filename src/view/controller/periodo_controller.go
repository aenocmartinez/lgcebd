package controller

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
	usecase "ebd/src/usecase/periodo"
	formrequest "ebd/src/view/formrequest/periodo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarPeriodos(c *gin.Context) {
	useCase := usecase.NewListarPeriodosUseCase(di.GetContainer().GetPeriodoRepository())
	response := useCase.Execute()
	c.JSON(response.StatusCode, response)
}

func CrearPeriodo(c *gin.Context) {
	var request formrequest.PeriodoFormRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", nil))
		return
	}

	if err := request.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	useCase := usecase.NewCrearPeriodoUseCase(di.GetContainer().GetPeriodoRepository(), di.GetContainer().GetCursoRepository())
	response := useCase.Execute(request.ToDTO())
	c.JSON(response.StatusCode, response)
}

func ActualizarPeriodo(c *gin.Context) {
	var request formrequest.PeriodoFormRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", nil))
		return
	}

	if err := request.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	useCase := usecase.NewActualizarPeriodoUseCase(di.GetContainer().GetPeriodoRepository())
	response := useCase.Execute(id, request.ToDTO())
	c.JSON(response.StatusCode, response)
}

func EliminarPeriodo(c *gin.Context) {
	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	useCase := usecase.NewEliminarPeriodoUseCase(di.GetContainer().GetPeriodoRepository())
	response := useCase.Execute(id)
	c.JSON(response.StatusCode, response)
}

func ListarCursosDePeriodo(c *gin.Context) {

	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	useCase := usecase.NewListarCursosPeriodoUseCase(di.GetContainer().GetPeriodoRepository())
	response := useCase.Execute(id)
	c.JSON(response.StatusCode, response)
}

func ListarAlumnosMatriculados(c *gin.Context) {

	periodoID, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID de periodo inválido", nil))
		return
	}

	cursoID, err := shared.ConvertStringToID(c.Param("curso_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID de curso inválido", nil))
		return
	}

	useCase := usecase.NewListarAlumnosMatriculadosUseCase(
		di.GetContainer().GetMatriculaRepository(),
		di.GetContainer().GetCursoPeriodoRepository())

	response := useCase.Execute(periodoID, cursoID)
	c.JSON(response.StatusCode, response)
}

func AgregarContenidoTematico(c *gin.Context) {
	var req formrequest.AgregarContenidoTematicoFormRequest
	periodoID, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID de periodo inválido", nil))
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", nil))
		return
	}

	agregarContenidoTematico := usecase.NewAgregarContenidoTematicoUseCase(di.GetContainer().GetContenidoTematicoRepository(), di.GetContainer().GetCursoPeriodoRepository())
	response := agregarContenidoTematico.Execute(periodoID, req.Descripcion)

	c.JSON(response.StatusCode, response)
}

func QuitarContenidoTematico(c *gin.Context) {
	periodoID, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID de periodo inválido", nil))
		return
	}

	contenidoTematicoID, err := shared.ConvertStringToID(c.Param("contenido_tematico_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID de contenido temático inválido", nil))
		return
	}

	quitarContenidoTematico := usecase.NewQuitarContenidoTematicoUseCase(di.GetContainer().GetContenidoTematicoRepository(), di.GetContainer().GetCursoPeriodoRepository())
	response := quitarContenidoTematico.Execute(periodoID, contenidoTematicoID)

	c.JSON(response.StatusCode, response)
}

func ObtenerContenidoTematicoDeUnCursoPeriodo(c *gin.Context) {

	periodoID, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID de periodo inválido", nil))
		return
	}

	obtenerContenidoTematico := usecase.NewListarContenidoTematicoDeUnCursoPeriodoUseCase(di.GetContainer().GetContenidoTematicoRepository(), di.GetContainer().GetCursoPeriodoRepository())
	response := obtenerContenidoTematico.Execute(periodoID)

	c.JSON(response.StatusCode, response)
}
