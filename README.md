## Descripción

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
