# Usa una imagen base de Go
FROM golang:1.23rc1-alpine3.20

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src/infrastructure-catalog-backend

# Copia los archivos del módulo Go
COPY go.mod go.sum ./

# Descarga las dependencias de Go
RUN go mod download

# Copia el código fuente de la aplicación
COPY . .

# Cambia al directorio de la aplicación
WORKDIR /go/src/infrastructure-catalog-backend/src

# Expone el puerto en el que se ejecutará la aplicación
EXPOSE 8000

# El comando por defecto para ejecutar la aplicación
CMD ["go", "run", "main.go"]