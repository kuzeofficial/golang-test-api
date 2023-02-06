# golang-test-api
A simple API with CRUD made with Fiber⚡️ in Golang 🐨
![Group 3](https://user-images.githubusercontent.com/65286318/192130130-0c4633fe-2462-456b-8902-5af83707c294.png)

## Starting

-> to use this server is needed install dependencies with the command
```bash
go mod tidy
```

-> you need add your env variables to a file `.env` in root directory with the example of variable showed in the file `.env.development` assigned, for this example i use [Deta](https://www.deta.sh) this is a free service cloud with a lot of integrations and good features:
```txt
DETA_PROJECT_KEY=XXX
DETA_BASE_NAME=XXX
```

-> then you will need go to `cmd/` and execute:
```bash
go run main.go
```
-> then you can go to Postman or any client http and send a request to the endpoints declared in `main.go` example:

![CleanShot 2023-02-05 at 20 58 28](https://user-images.githubusercontent.com/65286318/216854018-c3c2ec3e-fb76-4b48-ad45-b19c1e564166.png)


