package formrequest

type AgregarContenidoTematicoFormRequest struct {
	Descripcion string `json:"descripcion" binding:"required,max=255"`
	Orden       int    `json:"orden" binding:"numeric,min=1"`
}
