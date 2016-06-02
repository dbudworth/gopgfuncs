PGINC=/usr/include/postgresql/9.5/server/
DBNAME=postgres

export CGO_CFLAGS=-I$(PGINC)

PGMOD=$(abspath out/libpgmod.so)

all: $(PGMOD)

$(PGMOD): *.c *.go Makefile
	go build -buildmode=c-shared -o $(PGMOD)

install: $(PGMOD)
	psql $(DBNAME) --set=MOD=\'$(PGMOD)\' -f install.sql

clean:
	rm -rf out
