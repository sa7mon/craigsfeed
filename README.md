# craigsfeed

RSS feed generator for Craigslist searches

## Usage

```bash
go build -o craigsfeed main.go
./craigsfeed -url <search URL here> -interval 60
```

```
 -interval int
        Minutes to wait between scrapes (default 120)
  -url string
        URL of Craigslist search
```

## License

MIT