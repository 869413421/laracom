build:
		docker build -t laracom/service service/
		docker build -t laracom/userservice user-service/
		docker build -t laracom/productservice product-service/
		docker build -t  laracom/cli cli/