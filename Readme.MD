
# Chapter 1
Hexagonal Arch Folder Structure(this project) except {storage, auth}:

—— (/cmd)
main.go connect Port and Adapter and Routing(api)

—— Inside Hexagonal (/internal)

    —— (/adapters)
    Adapter
        -> gin, postgres, blaๆ. 
        in this layer for this project, follow me  ..V..below..V..
        for db(/postgres) ~cd postgres
            (postgres.go) <- database connecting
        
        for grpc(/grpc)

        for gin(/gin)


    ——  (/ports)
    Port Layer wait for Adapter plug in
        ->  interface for Service(usecase) & repositories(data access layer)
        
    —— (/infrastructure)
    Infrastructure(name it for easily for ME).
    It belike something belike sql query, nosql, handler for restapi and grpc.
    It is a piece of Adapter layer.

        —— (/handler)
        Handler (gin, fiber, mux, ...)

        —— (/grpc)
        Including grpc client and server (it's can be handler and repository in this project)

        —— (/repositories)
        Data acess layer(sql query blaๆ)

    —— (/application)
    Application Layer 

        —— (/core)
        Core of hex arch. 
        impl Usecase layer(service) and Entity layer(model).

        	Usecase Layer (/services)
        	Bussiness logic

        	Domain Layer (/entity)
        	Entity or model.

    ——

# TODO
- ssl(grpc) use same key (seperate later after done of drafting)
- whole services using postgres but changing to nosql, blaๆ later bcuz of this is fastest for me for drafting this project (dont worry, this is hex arch)
- distributed transaction (VERY IMPORTANT!!), currently Atomicity support only single service transaction.
- Phantom Phenomenon problem
- handler error by business logic
- seperate github repo
- caching on redis
- deploy on kube


# Rule
- dont change bussiness logic layer when switch adapter
- everything in bussiness logic layer must dont know any (dont call package belike grpc, gin, fiber, postgres, mysql and blaๆ in this layer)
- seperate grpc per service(not sharing)
- storage will use Storage services for getting image source
- in bussiness logic layer dont use(most case) same entity with db receiver entity (if u dont lazy), use {entityName}Response.

# Migration
migrate-up:
	cd internal/adapters/database && migrate -database $$DATABASE_URL -path migrations up
migrate-down:
	cd internal/adapters/database && migrate -database $$DATABASE_URL -path migrations down

# Command
docker run -p 8001:8080 --name user-con muniverse/user

docker run --name testdog2 --env-file ./.env -p 8001:8080 user_api

# migrate create -ext sql -dir internal/adapters/database/migrations -seq create_test    
# migrate create -ext sql -dir database/migrations -seq create_test    