package test

import (
	"api-golang/models"
	"testing"
)

func TestConnection(t *testing.T) {
	connection := models.GetConnection()
	if connection == nil {
		t.Error("No se pudo establecer la conexion con la base de datos")
	}
}
