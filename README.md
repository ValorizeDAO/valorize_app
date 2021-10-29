![Valorize App](valorize.png)

# Lunch a token and turn your community into a DAO

## Requirements
- Go 1.14 or above
- MySQL 8.0 or above

---

## Setup

  - [1. Clone repository](#1-clone-repository)
  - [2. Install Go modules](#2-install-go-modules)
  - [3. Create new MySQL database locally](#3-create-new-mysql-database-locally)
  - [4. Configure your local environment](#4-configure-your-local-environment)
  - [5. Run migrations](#5-run-migrations)
  - [6. Run app](#6-run-app)

### 1. Clone repository

```
git clone git@github.com:ValorizeDAO/valorize_app.git
```

### 2. Install Go modules

```
go get .
```

### 3. Create new MySQL database locally
You can create a new database from the command line or whatever GUI client you currently use

### 4. Configure your local environment
1. Duplicate `.env.example` and name it `.env`
2. After duplicating the example environment file, make sure to update the database credentials to your local credentials and use the name of database you created in step 3
3. You also need to get set up with an rpc provider and point the .env variables `TESTNET_NODE` and `MAINNET_NODE` to it. The provider can be either a truffle instance, Infura, or Alchemy url. This is necessary for the app to function in different capacities.
4. Services can be left blank or setup by getting your own keys.
5. You will also need to set up a keystore file for the operations that require accessing testnet. The one way to get one is to generate it by downloading geth and creating a new account instructions. Then copying the contents of your keystore json as the environment variable.
6. Lastly send some test `ETH` to your address through a faucet.

### 5. Run migrations

```
go run ./db/migrations/run_migrations.go
```


### 6. Run app

```
go run web.go
```

### That's it, you should be ready to go🚀

---

## License
Valorize App is open source under the [MIT License](LICENSE.md)

## Troubleshooting
If you run into any issues, please reach out to us in the `#dev-general` channel in our [Discord](https://discord.gg/3PRMWrH9DT)

## Contributing
Pull requests are welcome!
For meaningful changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate☑️
