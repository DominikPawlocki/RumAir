Setting up local environment:

1. Install GCC - its needed for MongoDb driver. If You have error like :

PS > go run main.go
# github.com/DataDog/zstd
exec: "gcc": executable file not found in %PATH% 

On Windows install TDM-GCC , with adding to PATH, then reboot or see point 2 !

2. Add 2 environment variables, then restart VS code 
