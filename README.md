# Go-Clean-Boilerplate

_Simple Go web service clean architecture boilerplate_

## Author's Note

- The project itself might stil not that flexible for some usecases.
- open issues for any errors or bugs
- mostly inpired by ([Rayato159](https://github.com/Rayato159) and [Blue-Cheesecake](https://github.com/Blue-Cheesecake))

## Setup Instructions (Ready to Use)

1. **Rename and Configure Environment File**

- Edit `dev.yaml`, `stg.yaml` and `prd.yaml` in `configs` folder for development, staging and production configuration.
- You may edit `docker-compose.yaml` to match your config.
- Change the file named `env` to `.env` and fill in the `SECRET_KEY` and `NODE_ENV`, which is the JWT signature key and environment respectively for the project.

2. **Start Database with Docker**

   ```sh
   docker compose up -d
   ```

   This command opens the connection to the PostgreSQL database using docker image or further configuration for your own needs.

3. **Database Connection (default)**

   - Connect to the database via any database IDE with the following credentials:
   - **Host:** localhost
   - **Port:** 5432
   - **Name:** mydb
   - **User:** myuser
   - **Password:** mypassword

- I use Postgresql as default database, you may implement yours in `databases` folder

1. **Database Setup**

- Mock Data: `/pkg/utils/data/`
- Entities (Tables): `/pkg/entities/`

2. **Complete all TODO**

- Please see the TODO list from the keyword search or any extension and complete all the tasks.
- If you are already satisfied with the default settings, then feel free to remove them.

## Features

- **JWT Authentication and Authorization (gin)**
- **Basic Auth (gin)** (Sign-in, Sign-up)
- **CORS (gin)**
- **Digestive Configuration** (`/config.yaml` - any additional configurations can be added here)
- **Automatically Build and Push docker image to DockerHub when Pushed (GitHub Action)**
