# Scoreboard

## Structure

Scoreboard is separated in 3 modules - server, client,  and database.

- Database and server are configurated and have to run via docker-compose
- Client is a go CLI application which has to be compiled

## Server

- main.go - server and database initialization 
- db.go - database initialization and functions
- server-go - server initialization and handlers

RPC server with 3 requests:
- StoreScore - bidirectional
- GetLeaderboard - unary
- Seed - unary

Code is separated in three files, nothing special (no controllers etc).

### Authorization

Server uses interceptors which check request context for token, if token is incorrect, then appropriate error code will be returned(Unauthenticated).

### Store score 

Arguments:

name - players name
points - players points

Takes name and points as arguments.
Stores score if person with given name doesn not exist.
Updates score if person exists and new score is greater than previous.

Response:

rank - players runk in leaderboard.

### GetLeaderboard

Arguments:

- name - players name
- page_size - page size
- page - page number
- monthly - indicates whethe current month should be used as a global period

Makes score table base on page and page_size.  
If page + page_size is incorrect, then appropriate error code will be returned(InvalidArgument).  
If monthly is true then scores will be taken for current month.  
If name is passed and user is not in current or previous pages, then around_me section will be added.  
Around_me sections contains 0-5 players before and 0-5 players after player with given name.  

Response:

- next_page - next page, "0" if there is no.
- score - score board based on page_size and page.
- around_me - additional board with players around given player.

### Seed

Simple request which seeds data for testing purposes.

### Notes

Idea was simple - use single query which generates full score table, it would look something like this:

```SELECT scores.name, scores.points, RANK () OVER (PARTITION BY scores.name ORDER BY scores.points DESC) AS position FROM user_scores AS scores WHERE scores.name = ?```

Where ```RANK () OVER (PARTITION BY scores.name ORDER BY scores.points DESC) AS position``` would do the job, but there was problem for both MySQL and PostgreSQL adapters which would mark query as invalid because of "RANK..." part.

To fix this problem data manipulation is separated in two parts:

1. Use SQL to get scores sorted by points
2. Do other calculations, such as position or pagination in go runtime


## Client

Initial idea was to implement it on React, but it happens that grpc-web does not support bidirectional streaming (StoreScore is bidirectional),
so client is a simple go CLI application.

## Configuration

Configrration is predefined.

All configuration is made via environment variables:

server - ./configs/server.env
database - ./config/db.env

## How to use
### Build binary

1. ```git clone https://github.com/vasjaj/Scoreboard```
2. ```docker-compose up --build -d``` (server waits 10s after database startup)
3. ```cd client```
4. ```go build -o client .```
5. ```./client```

### Help
```
Usage:
   [command]

Available Commands:
  help        Help about any command
  save        Starts stream for score saving
  seed        Creates new records
  show        Shows table results

Flags:
  -h, --help           help for this command
      --token string   Authentication token (default "correct_token")
      --url string     Service url with port (default "localhost:50051")

Use " [command] --help" for more information about a command.
```

### Worflow

1. Seed data with ```./client seed```
2. Add new with  ```./client save```
3. See scoreboard ```./client show```