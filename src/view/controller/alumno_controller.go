package controller

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
	usecase "ebd/src/usecase/alumno"
	formrequest "ebd/src/view/formrequest/alumno"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarAlumnos(c *gin.Context) {
	alumnoRepo := di.GetContainer().GetAlumnoRepository()
	listarAlumnos := usecase.NewListarAlumnosUseCase(alumnoRepo)
	response := listarAlumnos.Execute()
	c.JSON(response.StatusCode, response)
}

func CrearAlumno(c *gin.Context) {
	var request formrequest.AlumnoFormRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", nil))
		return
	}

	if err := request.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	alumnoRepo := di.GetContainer().GetAlumnoRepository()
	crearAlumno := usecase.NewCrearAlumnoUseCase(alumnoRepo)
	response := crearAlumno.Execute(request)
	c.JSON(response.StatusCode, response)
}

func ActualizarAlumno(c *gin.Context) {
	var request formrequest.AlumnoFormRequest

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

	alumnoRepo := di.GetContainer().GetAlumnoRepository()
	actualizarAlumno := usecase.NewActualizarAlumnoUseCase(alumnoRepo)
	response := actualizarAlumno.Execute(id, request)
	c.JSON(response.StatusCode, response)
}

func EliminarAlumno(c *gin.Context) {
	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	alumnoRepo := di.GetContainer().GetAlumnoRepository()
	eliminarAlumno := usecase.NewEliminarAlumnoUseCase(alumnoRepo)
	response := eliminarAlumno.Execute(id)
	c.JSON(response.StatusCode, response)
}

func MatricularAlumno(c *gin.Context) {
	var request formrequest.MatriculaFormRequest
	AlumnoID, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", nil))
		return
	}

	if err := request.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	useCase := usecase.NewMatricularAlumnoUseCase(
		di.GetContainer().GetAlumnoRepository(),
		di.GetContainer().GetCursoPeriodoRepository(),
		di.GetContainer().GetMatriculaRepository(),
	)

	response := useCase.Execute(AlumnoID, request.CursoPeriodoID)
	c.JSON(response.StatusCode, response)
}
