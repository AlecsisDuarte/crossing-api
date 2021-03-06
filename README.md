# Crossing API

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-0%25-brightgreen.svg?longCache=true&style=flat)</a>

Simple, fast and reliable RESTful API to fetch information about U.S. border's wait times.

## Getting Started

This project requires a [Firebase account](https://console.firebase.google.com/) in order to store all ports related information.

1. Create [Firebase project](https://console.firebase.google.com/)
2. Create test database within your firebase project
3. Download the database private key using this [tutorial](https://firebase.google.com/docs/admin/setup/#initialize-sdk)
5. Create an account in [Currency Layer](https://currencylayer.com/quickstart) and copy the API access key
4. Create your `.env`  file with this keys:

```env
   PRODUCTION = <BOOL> //Set TRUE if running a production environment
   DATABASE_URL = <STRING> //URL towards your created Database
   SERVICE_ACCOUNT_KEY_PATH = <STRING> //Relative path to the private key
   CURRENCY_LAYER_ACCESS_KEY = <STRING> // Access key from currencylayer.com
```

5. Run with `go run main.go`
6. Go to [localhost:8080](localhost:8080)

## License

This project is licensed under the [MIT license](https://github.com/carllerche/tower-web/blob/master/LICENSE).

## Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in
`crossing-go` by you, shall be licensed as MIT, without any additional terms or conditions.
