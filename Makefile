start:
	docker build -t backend-go .
	docker run -d -p "8081:8080" backend-go
