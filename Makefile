build:
	git pull
	docker-compose build
	docker-compose up -d

run-local:
	# cd auth/cmd && go run .
	# cd order/cmd && go run .
	# cd payment/cmd && go run .
	# cd seer/cmd && go run .
	# cd storage/cmd && go run .
	# cd user/cmd && go run .
	# cd voucher/cmd && go run .
	# cd wallet/cmd && go run .