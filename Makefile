.PHONY: all build install uninstall clean

BINARY_NAME=lab6
SERVICE_NAME=lab6.service
NGINX_FILE=lab6.conf
GO=go

all: check_deps build

build:
	# Crear los directorios de salida
	mkdir -p bin

	# Inicializar el modulo
	if [ ! -f go.mod ]; then \
		$(GO) mod init github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB; \
	fi

	# Descargar las dependencias de Go
	$(GO) mod tidy
	$(GO) mod download

	# Compilar el binario
	$(GO) build -o bin/$(BINARY_NAME) server/server.go

install: check_deps build
	# Copiar el binario a /usr/local/bin/
	cp bin/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

	# Copiar el archivo de servicio de systemd
	cp config/$(SERVICE_NAME) /etc/systemd/system/$(SERVICE_NAME)

	# Copiar el archivo de configuración de nginx
	cp config/$(NGINX_FILE) /etc/nginx/conf.d/$(NGINX_FILE)

	# Restart nginx
	systemctl restart nginx

	# Habilitar y arrancar el servicio
	systemctl enable $(SERVICE_NAME)
	systemctl start $(SERVICE_NAME)

uninstall:
	# Detener y deshabilitar el servicio
	systemctl stop $(SERVICE_NAME)
	systemctl disable $(SERVICE_NAME)

	# Eliminar el archivo de servicio de systemd
	rm -f /etc/systemd/system/$(SERVICE_NAME)

	# Eliminar el archivo de configuración de nginx
	rm -f /etc/nginx/conf.d/$(NGINX_FILE)

	# Restart nginx
	systemctl restart nginx

	#Eliminar los logs
	rm -rf /var/log/lab6

	# Eliminar el binario
	rm -f /usr/local/bin/$(BINARY_NAME)

clean:
	rm -rf build
	rm -rf bin

check_deps:
	# Verificar que Go esté instalado
	@command -v $(GO) >/dev/null 2>&1 || { echo "Go is not installed. Please install Go and try again."; exit 1; }

	# Verificar que systemd esté instalado
	@command -v systemctl >/dev/null 2>&1 || { echo "systemd is not installed. Please install systemd and try again."; exit 1; }

	# Verificar la instalación de SSH
	@command -v ssh >/dev/null 2>&1 || { echo >&2 "SSH is not installed. Please install SSH and try again."; exit 1; }

	# Verificar la instalación de nginx
	@command -v nginx >/dev/null 2>&1 || { echo >&2 "nginx is not installed. Please install nginx and try again."; exit 1; }