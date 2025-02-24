#!/bin/bash

set -e  # Detiene la ejecución si hay errores

CONTAINER_NAME="lgcebd-container"

# Función para compilar correctamente en la máquina host
compile() {
    echo "🔄 Compilando código Go para Linux..."
    GOOS=linux GOARCH=amd64 go build -o main .
}

# Función para limpiar logs del contenedor
clean_logs() {
    LOG_PATH=$(docker inspect --format='{{.LogPath}}' $CONTAINER_NAME 2>/dev/null || echo "")

    if [ -n "$LOG_PATH" ]; then
        echo "🧹 Limpiando logs del contenedor..."
        sudo truncate -s 0 "$LOG_PATH" 2>/dev/null || true
    else
        echo "⚠️ No se encontraron logs para limpiar."
    fi
}

# Función para reiniciar el contenedor sin reconstruir la imagen
restart_container() {
    echo "♻️ Reiniciando contenedor..."
    docker stop $CONTAINER_NAME >/dev/null 2>&1 || true

    clean_logs  # 🔄 Llamar a la función de limpieza de logs antes de reiniciar

    docker start $CONTAINER_NAME || docker compose up -d
    echo "✅ Contenedor en ejecución."
}

# Verifica si se debe compilar o construir
case "$1" in
    --compile)
        compile
        clean_logs   # 🔄 Ahora también limpia los logs en `--compile`
        restart_container
        ;;
    --build)
        echo "🔨 Construyendo imagen y reiniciando..."
        compile  # Asegurar que el binario sea correcto antes de reconstruir
        docker compose down
        clean_logs   # 🔄 También limpia los logs antes de construir la imagen
        docker compose up -d --build
        ;;
    *)
        restart_container
        ;;
esac
