# transaction-risk
Go Lang Application built to calculated the risk for a list of financial transactions, based on a set of rules.

# Tests
Used the native `go lang` test framework, to run all tests in the root folder of the project execute:

```
git test ./...
```

# Execution
The application run in the port 3000, for now it's fixed but can be easily changed to use envars. To start execution is just execute the `main.go` file with the command:
```
go run main
```

but if your setup is not able to run directly the file this project is on docker, just need these commands:
```
docker build -t transaction-risk -f dockerfile .

docker run -it -p 3000:3000 transaction-risk
```

# Example
Once with the application running with the port 3000, this is a cURL example for a valid request.

```
curl --location --request POST 'localhost:3000/transaction-risk' \
--header 'Content-Type: application/json' \
--data-raw '{
    "transactions": [
        {"id": 1, "user_id": 1, "amount_us_cents": 200000, "card_id": 1},
        {"id": 2, "user_id": 1, "amount_us_cents": 600000, "card_id": 1},
        {"id": 3, "user_id": 1, "amount_us_cents": 1100000, "card_id": 1},
        {"id": 4, "user_id": 2, "amount_us_cents": 100000, "card_id": 2},
        {"id": 5, "user_id": 2, "amount_us_cents": 100000, "card_id": 3},
        {"id": 6, "user_id": 2, "amount_us_cents": 100000, "card_id": 4}
    ]
}'
```