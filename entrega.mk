
# Nombre del ejecutable final
BIN = netstats

# Directorio donde está main.go
DIR = internet

# Regla principal (dada en el enunciado)
netstats:
	cd $(DIR); go build -o ../$(BIN)

# Regla opcional de limpieza
clean:
	rm -f $(BIN)