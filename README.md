# API REST en Go

El programa crea dos servicios web utilizando el framework Gin: un servicio de usuarios y un servicio de procesamiento. Estos servicios se ejecutan en puertos diferentes y tienen diferentes rutas y controladores asociados. El programa también configura archivos de registro y redirige la salida estándar y la salida de error a estos archivos.

### Autores:
- **Bottini, Franco Nicolas**

### ¿ Como compilar ?

Se puede compilar y ejecutar el programa utilizando el archivo Makefile de la siguiente manera:

```bash
$ git clone https://github.com/francobottini99/APIRESTGO-2023.git
$ cd APIRESTGO-2023
$ sudo make install
```

Esto genera el binario y lo ejcuta como un servicio utilizando `systemd`, ademas configura el programa `nginx` para que redirija las solicitudes a los servicios web correspondientes.

> [!NOTE]
> Para compilar el proyecto es necesario tener instalado `Go`, `systemd` y `nginx` en el equipo.

Para verificar que el programa se ejecuto correctamente se puede utilizar el comando:

```bash
$ systemctl status lab6
```

Finalmente, si se desea desinstalar el programa y eliminar todos los archivos generados, se puede utilizar el comando:

```bash
$ sudo make uninstall
```

## Servicios

Los servicios web se ejecutan sobre los puertos `8555` y `85556` y se puede acceder a ellos a traves de los dominios `dashboard.com` y `sensors.com` respectivamente. Para hacer uso de los *end points* expuestos por los servicios es necesario enviar un token `JWT` en el *header* *"Authorization"* de la solicitud. El token se obtiene al loguearse en el servicio de usuarios.

### Servicio de Usuarios

El servicio de usuarios se ejecuta sobre el puerto `8555` y se puede acceder a el atraves del dominio `dashboard.com`. Expone tres *end points* que permiten crear nuevos usuarios, loguearse y obtener un listado de los usuarios existentes. Los *end points* son: `api/users/createuser`, `api/users/login` y `api/users/listall`.

#### **api/users/login**

Este *end point* permite validar a un usuario existente en el sistema. Para ello, se debe enviar una solicitud *POST* con el siguiente formato:

```json
{
    "username": "user",
    "password": "pass"
}
```

Si las credenciales corresponden con un usuario del sistema obtendremos como respuesta un token `JWT` que nos permitirá hacer uso de los demas *end point* del servidor.

```json
{
    "message": "Login successful !",
    "token": "token jwt"
}
```

#### **api/users/createuser**

Este *end point* permite crear un nuevo usuario en el sistema. Para ello, se debe enviar una solicitud *POST* con el siguiente formato:

```json
{
    "username": "user",
    "password": "pass"
}
```

Si el resultado de la creación es exitoso, obtendremos como respuesta la información del nuevo usuario generado.

```json
{
    "id": 1000,
    "username": "user",
    "create_at": "2023-06-08 10:03:15"
}
```

Los usuarios creados por medio de este *end point* estan habilitados para acceder al servidor por medio de `ssh`. 

#### **api/users/listall**

Este *end point* permite obtener un listado de todos los usuarios registrados en el sistema. Para ello, se debe enviar una solicitud *GET* al mismo.

Obtenemos como respuesta un listado con la información de todos los usuarios registrados en el sistema.

```json
[
    {
        "id": 1000,
        "username": "user",
        "create_at": "2023-06-08 10:03:15"
    },
    {
        "id": 1001,
        "username": "user1",
        "create_at": "2023-06-08 10:03:15"
    }
]
```

### Servicio de Procesamiento

El servicio de procesamiento se ejecuta sobre el puerto `8556` y se puede acceder a el a traves del dominio `sensors.com`. Expone dos *end points* que permiten subscribir información de los sensores y obtener la información de los mismos. Los *end points* son: `api/sensors/submit` y `api/sensors/summary`.

#### **api/sensors/submit**

Este *end point* permite subscribir información de los sensores. Para ello, se debe enviar una solicitud *POST* con el siguiente formato:

```json
{
    "processing": 0.235,
    "free_memory": 546,
    "swap": 32532
}
```

Si el resultado de la subscripción es exitoso, obtendremos como respuesta un mensaje de confirmación.

```json
{
    "message": "Processing submitted !"
}
```

#### **api/sensors/summary**

Este *end point* permite obtener la información de los sensores. Para ello, se debe enviar una solicitud *GET* al mismo.

Obtenemos como respuesta un listado con la información de los sensores.

```json
[
    {
        "id": 1,
        "processing": 0.235,
        "free_memory": 546,
        "swap": 32532,
    },
    {
        "id": 2,
        "processing": 0.2443,
        "free_memory": 3432,
        "swap": 3532,
    }
]
```

> [!NOTE]
> Para hacer uso de este *end point* no es necesario el token de autentificación.