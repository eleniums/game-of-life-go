EXECUTABLE=gameoflife.exe

all: $(EXECUTABLE)

deps:
	dep ensure

$(EXECUTABLE):
	go build -o $(EXECUTABLE) main.go

clean:
	rm $(EXECUTABLE)