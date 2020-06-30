# strongDM Go SDK

This is the official [strongDM](https://www.strongdm.com/) SDK for the Go programming language.

## Installation

```bash
$ go get github.com/strongdm/strongdm-sdk-go
```

## Authentication

If you don't already have them you will need to generate a set of API keys, instructions are here: [API Credentials](https://www.strongdm.com/docs/admin-guide/api-credentials/)

Add the keys as environment variables; the SDK will need to access these keys for every request.
```bash
$ export SDM_API_ACCESS_KEY=<YOUR ACCESS KEY>
$ export SDM_API_SECRET_KEY=<YOUR SECRET KEY>
```

## List Users
The following code lists all registered users:

```go
func main() {
	log.SetFlags(0)
	accessKey := os.Getenv("SDM_API_ACCESS_KEY")
	secretKey := os.Getenv("SDM_API_SECRET_KEY")
	if accessKey == "" || secretKey == "" {
		log.Fatal("SDM_API_ACCESS_KEY and SDM_API_SECRET_KEY must be provided")
	}

	client, err := sdm.New(accessKey, secretKey)
	if err != nil {
		log.Fatal("failed to create strongDM client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	users, err := client.Accounts().List(ctx, "")
	if err != nil {
		log.Fatal("failed to query accounts:", err)
	}
	for users.Next() {
		user := users.Value()
		fmt.Println(user)
	}
	if err := users.Err(); err != nil {
		log.Fatal("error while iterating users:", err)
	}
}
```

## Useful Links

* Documentation:  [sdm package · pkg.go.dev](https://pkg.go.dev/github.com/strongdm/strongdm-sdk-go?tab=doc)
* Examples: [GitHub - strongdm/strongdm-sdk-go-examples](https://github.com/strongdm/strongdm-sdk-go-examples)

## License

[Apache 2](https://github.com/strongdm/strongdm-sdk-go/blob/master/LICENSE)

## Contributing 

Currently, strongDM does not accept pull requests for this repository. Please submit any feedback to <support@strongdm.com>.
