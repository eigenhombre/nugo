.PHONY: all install clean

all: $(TARGET)

TARGET = nugo

$(TARGET): *.go
	go build -o $(TARGET) *.go

install:
	go install

test:
	rm -rf $(GOPATH)/nugotestproject
	make
	make install
	nugo nugotestproject
	cd $(GOPATH)/nugotestproject && make && ./nugotestproject && make test

docker:
	docker build -t nugo .

clean:
	rm -f $(TARGET)
