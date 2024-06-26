## Nebula Backend

1. **Ensure Dependencies are Installed**

   Make sure all the dependencies are installed. If you haven't already done so, you can install them by running:
    ```bash
   go mod tidy
    ```
   This command will ensure that all the dependencies specified in your `go.mod` file are installed.

2. **Set Up the Database**

   Ensure that your PostgreSQL database is running and accessible. The `dsn` (Data Source Name) in `config/config.go` is already configured to use environment variables for database credentials. 

   Ensure the environment variables `DB_HOST`, `DB_USER`, `DB_PASSWORD`, and `DB_NAME` are set in your `.env` file with your actual database details.

3. **Manage Database and Migrations**
   If you need to manage your database, you can use the `manage_db.sh` script. This script helps with creating and dropping the database. To create the database, run:
   ```bash
   ./manage_db.sh --create
   ```

   To drop the database, run:
   ```bash
   ./manage_db.sh --drop
   ```

   Please note that dropping the database is irreversible and will delete all data.

   To create a new migration, use the `create_migrations.sh` script. This script generates a new migration file with a timestamp and the provided name. To create a migration, run:
   ```bash
   ./create_migrations.sh
   ```

   You will be prompted to enter the migration name. Use underscores for spaces in the migration name. For example, if you want to create a migration to add a new table, you might name it `add_new_table`.

   After creating the migration file, you can add your migration logic in the `Migrate` function and the rollback logic in the `Rollback` function within the generated file.

4. **Apply Migrations**

   The migrations are automatically applied to your database when you start the application using `air -c .air.toml` (with live reloading) or `go run cmd/nebula_backend/main.go` (without live reloading).

5. **Run the Application**

   To run the application, execute the following command from the root directory of your project:

   ```bash
   go run cmd/nebula_backend/main.go
   ```

6. **Verify the Application is Running**

   Once the application is running, you should see output indicating that the server is running on port 8080. You can verify this by opening a web browser or using a tool like `curl` or Postman to make a request to the `/users` endpoint:

   ```bash
   curl http://localhost:8080/users
   ```

   If everything is set up correctly, you should receive a JSON response (even if it's an empty array if there are no users in the database).

    **Example Output**
    
    When you run the application, you should see output similar to this:
     ```bash
    2023-06-13T12:45:00.123Z INFO zap/logger.go:38 Database connected successfully
    2023-06-13T12:45:00.123Z INFO zap/logger.go:38 Migration did run successfully
    [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
    
    [GIN-debug] GET /users --> myapp/controllers.GetUsers (3 handlers)
    [GIN-debug] Listening and serving HTTP on :8080
    ```
    This indicates that the database connection was successful, migrations were applied, and the server is running on port 8080.
    
## Docker
# For Development   
docker-compose -f docker-compose.dev.yml up --build 

# For Production
docker-compose -f docker-compose.prod.yml up --build

# To Create Docker Image
docker build -t nebula_backend:latest -f Dockerfile.prod .
docker tag nebula_backend:latest abhionlyone/nebula_backend:latest
docker push abhionlyone/nebula_backend:latest