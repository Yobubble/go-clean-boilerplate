# Go-Clean-Boilerplate

_Simple Go web service clean architecture boilerplate_

## Author's Note

- An improved version of Go web service clean architecture from using Fiber to Gin.
- mostly inpired by ([Rayato159](https://github.com/Rayato159) and [Blue-Cheesecake](https://github.com/Blue-Cheesecake))

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

5. **Complete all TODO**
   - Please see the TODO list from the keyword search or any extension and complete all the tasks.
   - If you are already satisfied with the default settings, then feel free to remove them.

## Features

- **JWT Authentication and Authorization**
- **Basic Auth** (Sign-in, Sign-up)
- **User Table**
- **Gin Framework**
- **CORS**
- **Struct Validation**
- **Digestive Configuration** (`/config.yaml` - any additional configurations can be added here)
- **Automatically Build, Test, and Push to DockerHub when Pushed (Thanks to GitHub Action)**
- **Automatically Make Release when git tag is created and pushed**
