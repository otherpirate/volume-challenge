# volume-challenge
Senior Software Engineer Take-Home Programming Assignment for Golang

## How to run
```shell
PORT=8080 go run cmd/api/main.go
```

## Endpoints

##### `/calculate` 
* Find the origin and destination airports of a list of flights
* Method: POST
* Payload: `{"flights": [["airport1", "airport2"], ["airport2", "airport3"]]`
* Status OK(200): `{"path": "airport1->airport3"}`
* Status Bad Request(400): `{"error": "could not define path"}`
* Example: `curl -X POST localhost:8080/calculate -d '{"flights": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]}'`

##### `/paths` (BONUS)
* Find all possibles paths from start to end
* Method: POST
* Payload: `{"flights": [["someAirport", "fakeAirport"], ["fakeAirport", "lastAirport"]], "start": "someAirport", "end": "lastAirport"}`
* Status OK(200): `{"paths": ["someAirport->fakeAirport->lastAirport"]}`
* Example: `curl -X POST localhost:8080/paths -d '{"flights": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]], "start": "SFO", "end": "EWR"}'`

## The solution

Find the first and the last airport of a flights list.

My strategy was improve the performance, when we have O(N) of time complexy (n is the number of flights) and usage space O(N) in the worse case.

I had ths `pather` variable where I keep all in/out flights, so:
* The destinaion airport will miss (-1) a "take off"
* The origin airport will have (1) a "take off" more

I tried to map all case in the unit tests of `pather_test.go` file

Bonus: My first idea was find all possible paths that user can travel, like searching for flights. Then I did read one more time the challenge and I decided that I'm wrong. I let the graph code in the repo because I liked it :)


## Tests
The API layer and business logic layer has unit tests
