![](https://raw.githubusercontent.com/marcusolsson/gophers/master/gopherdata-gopher.png | width=100)

# GoFinance

GoFinance un proyecto de software libre cuyo objetivo es implementar herramientas para expertos en áreas de índole económica.

## Introducción

**Go** es necesario. Para instalarlo visite [https://golang.org/doc/install]. Despues de instalarlo, coloque el proyecto dentro de `$GOPATH/src` ya sea clonando el proyecto por medio de **git** o descargandolo como zip.

Una vez instalado es necesario correr `go get` para instalar los módulos externos.

### Valuación de bonos - Duración.

Por el momento solo la implementación de la duración de valuación de bonos, la misma se encuentra en main.go. Para usarlo edite el archivo excel TablaDuracion.xlsx e incluye tantas filas como problemas se deseen resolver.

Una vez editado el archivo solo basta con correr `go run main.go` para ver los resultados en pantalla.

---

Gopher Artwork por https://github.com/marcusolsson/gophers
