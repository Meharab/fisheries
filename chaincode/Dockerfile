# Build image
FROM golang:1.23 as build

WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download

# copy all code
COPY . .

# build server binary (adjust path to your cmd)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-w -s" -o /bin/fisheries ./

# Minimal runtime image
FROM gcr.io/distroless/static-debian11
COPY --from=build /bin/fisheries /bin/fisheries
EXPOSE 7052
ENTRYPOINT ["/bin/fisheries"]
