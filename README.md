## Descripción

La siguiente api rest en golang no tiene utilidad solo usar go, devolver JSON, string, devolver un HTML, devolver 404 si el endpoint no existe, uso de logs.

## Endpoints

Ver un JSON {
    "response": "pong"
}
```
/ping
```

Ver un mensaje en string, "Hello World"
```
/
```

Ver en HTML un gráfico de barras
```
/bar/
```
Ver en HTML un gráfico de torta
```
/pie/
```

Proyecto de golang de ejemplo, para ver si cambio algo.
El archivo es:
```
app.go
```

## Comandos

Correr el programa
```
go run app.go
```

Se genera un dibujo en una web ubicada en el mismo directorio, un:
```
bar.html
```
Es un gráfico de barras usando echarts<img width="852" alt="Captura de pantalla 2024-08-03 a la(s) 01 00 48" src="https://github.com/user-attachments/assets/6a18a634-6e19-4544-beba-8bd9b8f3a13f">

Grafico de torta, ubicado en:
```
pie.html
```
<img width="798" alt="Captura de pantalla 2024-08-03 a la(s) 14 55 13" src="https://github.com/user-attachments/assets/7c065e13-d7b8-4631-a480-e4215deb8987">

buildear
```
go build app.go
```

El mod lo genere:
```
go mod init app
```

Bajar echarts
```
go get -u github.com/go-echarts/go-echarts/..
```

**Fuente de echarts**
https://blog.logrocket.com/visualizing-data-go-echarts/

Para usar google app engine, primero hay que bajar gcloud-cli y en un repo local hacer init:
```
gcloud init
```

Deployar en google app engine
```
gcloud app deploy app.yaml
```

Se necesita un archivo de configuración para google app engine:
```
app.yaml
```
Ubicado en la raíz del directorio

Donde ver el programa en funcionamiento
https://nube-gpbybwgp7q-uc.a.run.app

Comando que use para deployar
```
gcloud run deploy
```
Cuando pregunta por el código fuente, apretar ENTER
en Service puse nube
Cuando pregunta por la zona: 32 (us-central1)
