#!/bin/bash

set -e  # Detiene la ejecución si hay error

BUILD=false  # No reconstruye por defecto
COMPILE=false # No compila por defecto

# Verificar argumentos
for arg in "$@"; do
    case $arg in
        --build)
            BUILD=true
            ;;
        --compile)
            COMPILE=true
            ;;
        -h|--help)
            echo "Uso: ./scripts/run.sh [--build] [--compile]"
            echo "  --build    Fuerza la reconstrucción de la imagen"
            echo "  --compile  Fuerza la compilación del binario antes de levantar el contenedor"
            exit 0
            ;;
        *)
            echo "Opción desconocida: $arg"
            exit 1
            ;;
    esac
done

echo "🛠️  Verificando dependencias de Go..."
go mod tidy

# Detectar cambios en go.mod o go.sum
if [[ go.mod -nt .docker-last-build || go.sum -nt .docker-last-build ]]; then
    echo "⚠️  Cambios detectados en dependencias. Se requiere reconstrucción."
    BUILD=true
    COMPILE=true
fi

# Detectar cambios en el código fuente (*.go)
if find . -name "*.go" -newer .bin-last-build | grep -q .; then
    echo "⚠️  Cambios detectados en el código fuente. Se requiere recompilación."
    COMPILE=true
fi

touch .docker-last-build .bin-last-build

# 🛠️ **Compilar el binario solo si es necesario**
if [ "$COMPILE" = true ] || [ ! -f main ]; then
    echo "🔨 Compilando el binario para Linux..."
    GOOS=linux GOARCH=amd64 go build -o main .
    chmod +x main  # Asegurar permisos de ejecución
    echo "✅ Binario compilado correctamente."
fi

# **Si el contenedor ya está corriendo, decide si eliminarlo o dejarlo así**
if docker ps --filter "name=lgcebd-container" --format '{{.Names}}' | grep -q "lgcebd-container"; then
    if [ "$BUILD" = true ]; then
        echo "🚨 Deteniendo y eliminando el contenedor para reconstrucción..."
        docker compose down
    else
        echo "✅ El contenedor ya está en ejecución."
        exit 0
    fi
fi

# **Si el contenedor existe pero está detenido, solo reiniciarlo**
if docker ps -a --filter "name=lgcebd-container" --format '{{.Names}}' | grep -q "lgcebd-container"; then
    echo "♻️ Reiniciando contenedor..."
    docker start lgcebd-container
    exit 0
fi

# **Reconstruir imagen solo si es necesario**
if [ "$BUILD" = true ]; then
    echo "🚀 Construyendo la imagen..."
    docker compose build
fi

# **Levantar contenedor solo si es necesario**
echo "🚀 Iniciando contenedor..."
docker compose up -d

echo "✅ Contenedor ejecutándose correctamente."
