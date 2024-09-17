# RSS Aggregator.

It is a [RSS](https://en.wikipedia.org/wiki/RSS) feed aggregator in Go! It's a web server that allows clients to:

- Add RSS feeds to be collection
- Follow and unfollow RSS feeds that other users have added
- Fetch all of the latest posts from the RSS feeds they follow

Credits: [Boot.Dev](https://github.com/bootdotdev)

### Prerequisites

**go version go1.21.3**

### Cloning the repository

```shell
git clone https://github.com/nayak-nirmalya/rssagg.git
```

### Setup .env file

```js
PORT=8080
DB_URL=
```

### Setup DB

Migrate DB: Run following command inside `sql/schema` directory

```shell
goose postgres <YOUR_DB_URL> up
```

Generate Queries

```shell
sqlc generate
```

### Build & Start the Server

```shell
go build && ./rssagg
```
