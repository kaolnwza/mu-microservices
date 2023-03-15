migrate-up:
	echo "-----[Migrations UP...]-----"
	cd database && migrate -database $$DATABASE_URL -path migrations up
migrate-down:
	echo "-----[Migrations UP...]-----"
	cd database && migrate -database $$DATABASE_URL -path migrations down

# migrate create -ext sql -dir database/migrations -seq create_test    