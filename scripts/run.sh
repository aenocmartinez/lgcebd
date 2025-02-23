#!/bin/bash

set -e  # Detiene la ejecución si hay un error

# Verificar si Docker está instalado
if ! command -v docker &> /dev/null; then
    echo "Error: Docker no está instalado."
    exit 1
fi

# Detener y eliminar contenedores existentes antes de levantar nuevos
docker compose down

# Levantar los contenedores en modo detached (-d)
docker compose up -d --build

echo "El contenedor se está ejecutando."
