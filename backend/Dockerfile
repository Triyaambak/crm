FROM golang:alpine

WORKDIR /app

# Build-time args
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_HOST
ARG POSTGRES_PORT
ARG POSTGRES_DB
ARG API_PORT

# Construct DB_URL at build time
ENV DB_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

COPY go.mod go.sum /app/

# Downloading all external packages
RUN go mod download

# Installing all go binaries needed for POSTGRES version control
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ./ /app/

RUN sqlc generate

RUN go build ./main.go

EXPOSE ${API_PORT}

CMD ["sh", "-c", "migrate -path=./sql/migrations -database ${DB_URL} up && ./main"] 



