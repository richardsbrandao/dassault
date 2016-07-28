dassault: clean
	go get -u github.com/valyala/fasthttp	
	go get -u github.com/buaazp/fasthttprouter
	go get -u gopkg.in/pg.v4
	go get -u github.com/gocql/gocql
	go build
	mv dassault bin/dassault

clean:
	rm -f bin/*
        