# Golang Api Rest

This is a project to create an API REST using Go.

### Entorno de pruebas
Por defecto el estorno es de pruebas, si se desea deshabilitar este ambiente, se debe ir al archivo config/config.go y cambiar el valor de la variable debug:

```Go
// database.Debug = gonv.GetBoolEnv("DEBUG", true) <- Antes
database.Debug = gonv.GetBoolEnv("DEBUG", false)
```