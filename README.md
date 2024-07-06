# Go-Clean-Boilerplate

*Simple Go web service clean architecture boilerplate*

## Author's Note
- An improved version of Go web service clean architecture from using Fiber to Gin.
- mostly inpired by (by [Rayato159](https://github.com/Rayato159))

## Setup Instructions (Ready to Use)

1. **Rename and Configure Environment File**
    - Change the file named `env` to `.env` and fill in the `SECRET_KEY`, which is the JWT signature key.

2. **Start Database with Docker**
    ```sh
    docker compose up -d
    ```
    This command opens the connection to the PostgreSQL database.

3. **Database Connection**
    - Connect to the database via any database IDE with the following credentials:
      - **Host:** localhost
      - **Port:** 5432
      - **Name:** mydb
      - **User:** myuser
      - **Password:** mypassword

4. **Customize Your Database**
    - Mock Data: `/pkg/utils/data/`
    - Entities (Tables): `/pkg/entities/`

## Features

- **JWT Authentication and Authorization**
- **Basic Auth** (Sign-in, Sign-up)
- **User Table**
- **Gin Framework**
- **Struct Validation** (by [github.com/go-playground/validator/v10](https://github.com/go-playground/validator))
- **Digestive Configuration** (`/config.yaml` - any additional configurations can be added here)
