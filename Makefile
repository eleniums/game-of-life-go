EXECUTABLE=gameoflife.exe

all: $(EXECUTABLE)

deps:
	dep ensure

$(EXECUTABLE):
	go build -o $(EXECUTABLE) ./cmd/game/main.go

clean:
	rm $(EXECUTABLE)