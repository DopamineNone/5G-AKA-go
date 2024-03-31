# 5G-AKA-go (Four Entities)
GO code for 5G-AKA protocol(Four entities: UE, AMF/SEAF, AUSF, UDM/ARPF)

## Directory Structure

```shell
5G-AKA-go
|-- client            # client program
|-- idl               # thrift files
|-- kitex_gen         # Kitex dependencies
|   `-- _5gAKA_go
|       |-- AUSF
|       |-- SEAF
|       |-- UDM
|       `-- UE
|-- log               # log files; *.log is ignored by .gitignore
`-- rpc               # server programs
    |-- AUSF          
    |-- SEAF
    |-- UDM
    `-- UE
```

## Run

1. Start servers in respective terminal:

    ```shell
    # UDM
    cd rpc/UDM
    go run .
    
    # AUSF
    cd rpc/AUSF
    go run .
    
    # SEAF
    cd rpc/SEAF
    go run .
    
    # UE
    cd rpc/UE
    go run .
    ```

2. Run client:

    ```shell
    go run ./client
    # expected result:
    # > Authentication was completed successfully!
    ```