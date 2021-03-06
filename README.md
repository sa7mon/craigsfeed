# craigsfeed

RSS feed generator for Craigslist searches

## Usage

The HTTP server will listen on port 8000 for requests to `/rss`
```
 -interval int
        Minutes to wait between scrapes (default 120)
  -url string
        URL of Craigslist search
```

**Build from source**
```shell
go build -o craigsfeed main.go
./craigsfeed -url <search URL here> -interval 60
```

**Run via Docker**

```shell
docker build . -t craigsfeed:latest
docker run -p 8000:8000 --rm craigsfeed -url "https://your-craigslist-search.url" -interval 90
```



## License

MIT