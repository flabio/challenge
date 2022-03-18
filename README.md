# challenge

1) Archivo readme con los pasos para poder ejecutar la aplicación:
    docker compose build --no-cache
    docker compose up 
    docker compose down
 2) Documentación que explique cómo se podría autenticar y autorizar el consumo de los distintos endpoints 
    para distintos usuarios teniendo en consideración por ejemplo su criticidad.
    nota: en el postman seleccionamos la opcion de headers, en key =Authorization y en values el valor del del token del usuario values=eyJhbGciOiJIUzI
    user la aplicacion de postman para hacer prueba.
        1) login method post
           localhost:8080/api/v1/login
           {
           "name": "yorbax",
           "owner": "Venomoth"
           }
           me desvuelve 
            {
            "status": "200",
            "message": "OK!",
            "errors": null,
            "data": {
                "id": 49,
                "name": "yorbax",
                "criticidad": "high",
                "owner": "Venomoth",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQ5IiwiY3JpdGljaWRhZCI6ImhpZ2giLCJleHAiOjE2NDc2NDEzNDEsImp0aSI6IjQ5IiwiaWF0IjoxNjQ3NjQwNDQxLCJpc3MiOiJ0ZXN0In0.BvuT9gzpRcuykbkez_-isjwFnp3tNe11GU-0QhZ30B8"
            }
            }
         2) la enpoint de administrador:
            se valido por el criticidad por ejemplo
            criticidad high:GET /administrador, POST /administrador,PUT /administrador/:id, GET /administrador/:id y DELETE /administrador/:id
            criticidad medium:GET /administrador, POST /administrador,PUT /administrador/:id, GET /administrador/:id
            criticidad medium:GET /administrador
   3) Documentación de la aplicación realizada, problemas y soluciones con los que se encontró al realizar esta API,
      el problema que se queriria una aplicacion para controlar los acceso y los permison del usuario.
      la solucion: es crear una api para que los usuario tuviera accesos a la informacion que le corresponder por el criticidad,
      y se puede manipular de una manera mas facil.
      

