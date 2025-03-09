# Quick start

rqlited -node-id=1 db/

migrate -path db/migrations -database "rqlite://localhost:4001?x-connect-insecure=true" up

go run connector.go

npm run dev

# Toolchain needed

https://github.com/rqlite/rqlite
https://github.com/gin-gonic/gin
https://github.com/rqlite/gorqlite

1. Get RQLite
1. Launch RQLite Node
1. Run upstream migrations
1. Launch Go Gin router
1. Launch SvelteKit
1. Api & App should now work - insert into posts table in RQLite and refresh :5173 index in browser.
