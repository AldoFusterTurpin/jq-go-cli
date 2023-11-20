# jq-go-cli
Simple implementation of jq CLI in Go.

Code challenge from https://codingchallenges.substack.com/p/coding-challenge-34-jq

### Run the unit tests
```sh
go test ./... -cover -v
```
### Run the Step 1
```sh
go build && curl -s 'https://dummyjson.com/quotes?limit=2' | ./jq-go-cli
```

### Run the Step 2
```sh
go build && curl -s 'https://api.github.com/repos/CodingChallegesFYI/SharedSolutions/commits?per_page=3' | ./jq-go-cli ".[1]"
```
