#!/bin/bash

set -e  # Detiene la ejecución si hay errores

CONTAINER_NAME="lgcebd-container"

# Función para compilar correctamente en la máquina host
compile() {
    echo "🔄 Compilando código Go para Linux..."
    GOOS=linux GOARCH=amd64 go build -o main .
}

# Función para reiniciar el contenedor sin reconstruir la imagen
restart_container() {
    echo "♻️ Reiniciando contenedor..."
    docker stop $CONTAINER_NAME >/dev/null 2>&1 || true
    docker start $CONTAINER_NAME || docker compose up -d
    echo "✅ Contenedor en ejecución."
}

# Verifica si se debe compilar o construir
case "$1" in
    --compile)
        compile
        restart_container
        ;;
    --build)
        echo "🔨 Construyendo imagen y reiniciando..."
        compile  # Asegurar que el binario sea correcto antes de reconstruir
        docker compose down
        docker compose up -d --build
        ;;
    *)
        restart_container
        ;;
esac
