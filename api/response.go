package api

type GenericResponse struct {
	Status  int
	Message string
}

type GenericResponseWithData struct {
	Status  int
	Message string
	Data    any
}
