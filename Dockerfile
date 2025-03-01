# 1️⃣ Usa una imagen base de Golang
FROM golang:1.21 AS builder

# 2️⃣ Establece el directorio de trabajo en /app
WORKDIR /app

# 3️⃣ Copia los archivos de dependencias
COPY go.mod go.sum ./

# 4️⃣ Descarga dependencias
RUN go mod download

# 5️⃣ Copia el código fuente
COPY . .

# 6️⃣ Compilar la aplicación correctamente
#RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -ldflags="-s -w -extldflags '-static'" -o main .


# 7️⃣ Crear una imagen más liviana sin Golang
FROM debian:bullseye-slim

WORKDIR /app

# 8️⃣ Copiar el binario compilado desde la imagen builder
COPY --from=builder /app/main .

# 🔟 Definir el comando de ejecución
CMD ["/app/main"]
