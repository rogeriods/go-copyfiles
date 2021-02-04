# Copy Files

This is a simple routine to copy files from source path to a destination path.

In my case I'm using to copy my databases backup files to external HD. It is one of my daily routines, because of that
I'm trying to automate this service and put the build file on my schedule tasks to run every day at 7AM and 1PM.

Please if this code help you in some ways let me know!

## Getting Started

Running and build the code:

```bash
$ go run main.go

$ go build main.go
```

## Result

In my case the result of running this code is the following:

```bash
C:\Developer\go-copyfiles>go run main.go
Start to backup MySQL files
- - - - - - - - - - - - - - - - - - - - - - - - - - - -
graficos_manutencao.sql 12096 bytes copied
graficos_pa.sql 81515 bytes copied
graficos_survey.sql 12436 bytes copied
grupo_indemetal.sql 7960654 bytes copied
grupoindemetal_apps.sql 652279 bytes copied
indeme49_etiquetas.sql 19815568 bytes copied
indeme49_graficos.sql 6694515 bytes copied
indeme49_localapps.sql 64123 bytes copied
All files copied to 'D:\_MySQL\' with success!

graficos_manutencao.sql 12096 bytes copied
graficos_pa.sql 81515 bytes copied
graficos_survey.sql 12436 bytes copied
grupo_indemetal.sql 7960654 bytes copied
grupoindemetal_apps.sql 652279 bytes copied
indeme49_etiquetas.sql 19815568 bytes copied
indeme49_graficos.sql 6694515 bytes copied
indeme49_localapps.sql 64123 bytes copied
All files copied to '\\192.168.0.15\global_db-mysql\' with success!

Start to backup Oracle files
- - - - - - - - - - - - - - - - - - - - - - - - - - - -
SAPIENS.DMP 8743727104 bytes copied
SAPIENS_NFE.DMP 2949132288 bytes copied
TELEMATICA.DMP 47005696 bytes copied
VETORH.DMP 1714118656 bytes copied
All files copied to 'D:\_Oracle\' with success!

Start to backup SQL Server files
- - - - - - - - - - - - - - - - - - - - - - - - - - - -
etiquetas.bak 6131793408 bytes copied
graficos.bak 4357996032 bytes copied
All files copied to 'D:\_SQLServer\' with success!

```
