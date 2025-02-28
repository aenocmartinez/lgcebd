package formrequest

type AgregarContenidoTematicoFormRequest struct {
	Descripcion string `json:"descripcion" binding:"required,max=255"`
}
