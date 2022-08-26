# Compilation
reaction-roles-bot-go: go.sum main.go */**/*.go .env */**/**/*.go
	./scripts/build.sh

reaction-roles-bot-go.exe : go.sum main.go */**/*.go .env */**/**/*.go
	env CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 ./scripts/build.sh

# fetch dépendencies
go.sum : go.mod
	go get
	go mod tidy

.PHONY: install deb docs win all clean manuals
# require write rights for $(DESTDIR)
install : reaction-roles-bot-go
	mkdir -p $(DESTDIR)/usr/sbin/
	cp ./reaction-roles-bot-go $(DESTDIR)/usr/sbin/reaction-roles-bot-go

all: deb win linux manuals
	./scripts/publish.sh
	
deb :
	./scripts/deb.sh
win :
	./scripts/win.sh
linux : 
	./scripts/linux.sh

docs : */**/*.go
	mkdir -p docs
	godoc-static --destination ./docs .

manuals :
	./scripts/manuals.sh

clean : 
	rm -f reaction-roles-bot-go go.sum
	rm -rf package/
	rm -rf publish/