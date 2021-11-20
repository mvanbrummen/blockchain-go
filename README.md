# Blockchain Go

Creating a simple Blockchain in Golang with a HTTP API based on:

https://www.geeksforgeeks.org/create-simple-blockchain-using-python/

## Getting started

Test and build binary

```
make
```

Run app

```
./blockchain
```

Mine a block

```
curl http://localhost:8080/mine_block
{
  "index": 2,
  "message": "A block is mined",
  "previous_hash": "520e6a034dded45a979ed3b592b2ea0169ac68a2763bed64b69032a67357f362",
  "proof": 4216,
  "timestamp": "2021-11-20T20:50:03.804875+10:00"
}
```

Get the blockchain

```
curl http://localhost:8080/get_chain
{
  "chain": [
    {
      "index": 1,
      "timestamp": "2021-11-20T20:50:00.908664+10:00",
      "proof": 1,
      "previous_hash": "0"
    },
    {
      "index": 2,
      "timestamp": "2021-11-20T20:50:03.804875+10:00",
      "proof": 4216,
      "previous_hash": "520e6a034dded45a979ed3b592b2ea0169ac68a2763bed64b69032a67357f362"
    }
  ],
  "length": 2
}
```

Check if the blocks are valid

```
curl http://localhost:8080/valid
{
  "valid": true
}
```
