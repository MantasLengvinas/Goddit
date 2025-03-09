rqlited -node-id=1 db/

migrate -path db/migrations -database "rqlite://localhost:4001?x-connect-insecure=true" up

go run connector.go

npm run dev
