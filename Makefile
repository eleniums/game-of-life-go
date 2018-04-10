EXECUTABLE=gameoflife.exe

all: game

deps:
	dep ensure

game:
	go build -o $(EXECUTABLE) main.go

clean:
	rm $(EXECUTABLE)