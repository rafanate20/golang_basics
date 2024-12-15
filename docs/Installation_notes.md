## 1.- Instalar Go desde alguna de estas webs

[go.dev/doc/install](https://go.dev/doc/install) 

[go.dev/dl/](https://go.dev/dl/)

## 2.- Una vez instalado agregar variables de entorno, editando el archivo .bashrc o .profile o .zshrc y agregar al final

```
export GOPATH=/home/rnatera/go
export GOBIN=$GOPATH/bin
export GOROOT=/usr/local/go
export PATH=$PATH:$GOBIN:$GOROOT/bin
```

## 3.- Compilar con

```bash
$ go build
```

## 4.- Correr temporalmente con

```bash
$ go run
```

## 5.- Como Instalar un paquete

```bash
# Instalar el go tour
$ go install -v -u golang.org/x/website/tour@latest
#luego para ejecutarla sobre el $GOPATH ejecutar tour
$ cd $GOPATH
$ tour
```

## 6.- Iniciar un proyecto

```bash
sobre el directorio del proyecto ejecutar
$ go mod init
#or
$ go mod init github.com/dev-account/repo-name
```

## 7.- Instalar un modulo con go get, por ejemplo:

```bash
$ go get -v -u github.com/labstack/echo/v4
```

## 8.- Verificar estado de los modulos instalados
Revisamos que todos los modulos y sus dependencias se referencien en el go.mod del proyecto.
```bash
#Ejecutamos en la raiz del proyecto el comando 
$ go mod verify
```

## 9.- Modificar un modulo desde el repositorio origen
Clonamos el repositorio del modulo en nuestro path_proyecto/src,
luego hacemos las modificaciones del codigo
Redireccionamos la nueva version con el comando.
```bash
$ go mod edit -replace github/labstack/examplepkg=./src/mynewpkg
```

## 10.- Volver a la version original del modulo
Usamos el comando.
```bash
$ go mod edit -dropreplace github/labstack/examplepkg
```

## 11.- Empaquetar todo el codigo de librerias de terceros en el proyecto
Sobre el path del proyecto Ejecutamos.
```bash
$ go mod vendor
```

## 12.- Limpiar modulos y libs en deshuso
Sobre el path del proyecto Ejecutamos.
```bash
$ go mod tidy
```

## 13.- Check all external packages downloaded (go get cache)
Sobre el path del proyecto Ejecutamos.
```bash
$ go mod download -json
```

## LISTADO DE ENLACES DE INTERES
#### Go Tour: https://tour.golang.org.
#### Cheat Sheet -> https://devhints.io/go.
#### Webs con listados de paquetes y librerias de GO -> http://awesome-go.com/.
#### Golang Weekly: https://golangweekly.com.
#### Play With Go: https://play-with-go.dev.
#### Go by example: https://gobyexample.com.
#### Go Time (en Spotify).
#### Canal de slack de Gophers -> gophers.slack.com.

## Extra

## Instalar VSCode desde Software Manager (flatpak version)

### Configurar vscode con los pasos descritos en:

[Install visualstudio.code](https://github.com/flathub/com.visualstudio.code#readme)

```bash
#Abrir vscode desde terminal con
$ FLATPAK_ENABLE_SDK_EXT=golang /usr/bin/flatpak run com.visualstudio.code
#Instalar la extension de GO
```
