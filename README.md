# La Liga Tracker - Backend

La aplicación **La Liga Tracker** es un sistema de gestión de partidos de La Liga que permite a los usuarios visualizar, agregar, actualizar y eliminar partidos, así como gestionar estadísticas del partido como goles, tarjetas amarillas, tarjetas rojas y tiempo extra. Este proyecto contiene el backend necesario para interactuar con los datos de los partidos.

## Tecnologías utilizadas

- **Go**: Lenguaje de programación utilizado para el backend.
- **Gorilla Mux**: Router HTTP para la gestión de rutas y solicitudes RESTful.
- **SQLite**: Base de datos utilizada para almacenar los datos de los partidos.
- **CORS**: Configuración de CORS para permitir que el frontend acceda a la API sin restricciones.
  
## Endpoints

A continuación se detallan los endpoints disponibles en la API:

- `GET /api/matches`: Obtiene una lista de todos los partidos.
- `GET /api/matches/{id}`: Obtiene un partido específico por su ID.
- `POST /api/matches`: Crea un nuevo partido.
- `PUT /api/matches/{id}`: Actualiza un partido existente.
- `DELETE /api/matches/{id}`: Elimina un partido por su ID.
- `PATCH /api/matches/{id}/goals`: Actualiza los goles de un partido.
- `PATCH /api/matches/{id}/yellowcards`: Actualiza las tarjetas amarillas de un partido.
- `PATCH /api/matches/{id}/redcards`: Actualiza las tarjetas rojas de un partido.
- `PATCH /api/matches/{id}/extratime`: Actualiza el tiempo extra de un partido.

## Instalación

### Prerrequisitos

- Tener instalado **Go** (versión 1.24.1 o superior).
- Tener instalado **Docker** si deseas utilizar la versión en contenedor.
  
### Instrucciones para la instalación local

1. Clona el repositorio:

    ```bash
    git clone https://github.com/Ricgo01/liga_api
    cd liga_api
    ```

2. Instala las dependencias de Go:

    ```bash
    go mod tidy
    ```

3. Compila y ejecuta el servidor:

    ```bash
    go run main.go
    ```

El servidor se ejecutará en el puerto **8080**.

### Instrucciones para la instalación con Docker

1. Construir la imagen Docker:

    ```bash
    docker build -t liga-api .
    ```

2. Levantar el contenedor:

    ```bash
    docker-compose up
    ```
Esto pondrá en marcha el servidor en el puerto **8080**.

3. Corre el contenedor 
    ```bash
    docker run -p 8080:8080 liga_api:v1.0
    ```
## Imagenes del front funcionando con el api

![Captura de pantalla](/assets/1.png)
![Captura de pantalla](/assets/2.png)
![Captura de pantalla](/assets/3.png)
![Captura de pantalla](/assets/4.png)
![Captura de pantalla](/assets/5.png)
![Captura de pantalla](/assets/6.png)
![Captura de pantalla](/assets/7.png)
![Captura de pantalla](/assets/8.png)

## Link a ejemlos de peticiones con postman

https://www.postman.com/research-meteorologist-68792412/my-workspace/collection/i31z0cy/la-liga-api?share=true&origin=tab-menu

