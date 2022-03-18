# 20210317
Practice Cosmos-sdk based on HTTP request

### Requirements
* golang 1.16+
* air

### Set up Development

#### Install air
```
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
```
#### Running dev server
```bash
make dev
```
After running `make dev` command you can access http://localhost:8080/healthcheck then you see a `ok` message in body


#### ENDPOINT

| METHOD | URL | DESCRIPTION |
| -------- | --- | ----------- |
| GET | http://localhost:8080/query/bank/total | return total supply amount |
