# env file for /env

```env
# These are all mandatory fields which are required to start a Postgres server
POSTGRES_PORT = 5432
POSTGRES_USER = "admin"
POSTGRES_PASSWORD = 123
POSTGRES_DB = "CRM"

# Defined to run Dockerfile scripts
POSTGRES_HOST = "db"
API_PORT = 3001
```

# env file for /backend/.env

```env
# Variable used in go
API_PORT = 3001

# variables for postgress connection
POSTGRES_PORT = 5432
POSTGRES_USER = "admin"
POSTGRES_PASSWORD = 123
POSTGRES_DB = "CRM"
POSTGRES_HOST = "db"
```
