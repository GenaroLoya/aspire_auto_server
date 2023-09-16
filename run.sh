#!/bin/bash

# Directorio raíz del proyecto
project_root="$(pwd)"

# Crear directorio .logs si no existe
logs_dir="$project_root/.logs"
mkdir -p "$logs_dir"

# Directorio del servidor Go
server_dir="$project_root/apps/aspire.server"
cd "$server_dir"

# Nombre del archivo de log del servidor
server_log="$logs_dir/server.log"

# Ejecutar el servidor Go en segundo plano y redirigir la salida al archivo de log del servidor
gow run main.go > "$server_log" 2>&1 &

if [ $? -ne 0 ]; then
  echo "Error: No se pudo ejecutar el servidor Go."
  exit 1
fi

echo "Server HTTP is up... logs in file => $logs_dir/$server_log"

# Guardar el ID del proceso del servidor Go
server_pid=$!

# Directorio del cliente React
client_dir="$project_root/apps/aspire.client"
cd "$client_dir"

# Nombre del archivo de log del cliente React
client_log="$logs_dir/client.log"

# Ejecutar el cliente React en segundo plano y redirigir la salida al archivo de log del cliente React
npm run dev > "$client_log" 2>&1 &

if [ $? -ne 0 ]; then
  echo "Error: No se pudo ejecutar el servidor Vite.js."
  exit 1
fi
#Log para decir que el puerto esta levantado e indicar el archivo de logs
echo "Server HTTP/Vite.js is up... logs in file => $logs_dir/$client_log"

# Guardar el ID del proceso del cliente React
client_pid=$!

# Esperar a que ambos procesos terminen
wait $server_pid
wait $client_pid

