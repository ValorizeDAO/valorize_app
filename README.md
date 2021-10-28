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
Duplicate `.env.example` and name it `.env`

>After duplicating the example environment file, make sure to update the database credentials to your local credentials and use the name of database you created in step 3.

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
If you run into any issues, please email us at {{ ADMIN_EMAIL }} or reach out to us via our Discord

## Contributing
Pull requests are welcome!
For meaningful changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate☑️
