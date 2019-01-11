run:
		docker-compose up

down:
		docker-compose down

clean:
		docker rm db svc-coupon
build:	
		GOOS=linux GOARCH=amd64 go build -o ./svc-coupon -i ./*.go
		docker build -t db ./db
		docker build -t svc-coupon .
		rm ./svc-coupon
