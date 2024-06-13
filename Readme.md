To start the Go application that you have set up using the bootstrap script, follow these steps:

1. **Ensure Dependencies are Installed**

   Make sure all the dependencies are installed. If you haven't already done so, you can install them by running:

   go mod tidy

   This command will ensure that all the dependencies specified in your `go.mod` file are installed.

2. **Set Up the Database**

   Ensure that your PostgreSQL database is running and accessible. Update the `dsn` (Data Source Name) in `config/config.go` with your actual database credentials:

   dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"

   Replace `youruser`, `yourpassword`, `yourdb`, and other parameters with your actual database details.

3. **Run the Application**

   Navigate to the `cmd/nebula_backend` directory and run the main Go file. For example, if your project name is `myapp`, you would run:

   go run cmd/myapp/main.go

   Alternatively, you can run the application from the root directory of your project:

   go run cmd/nebula_backend/main.go

4. **Verify the Application is Running**

   Once the application is running, you should see output indicating that the server is running on port 8080. You can verify this by opening a web browser or using a tool like `curl` or Postman to make a request to the `/users` endpoint:

   curl http://localhost:8080/users

   If everything is set up correctly, you should receive a JSON response (even if it's an empty array if there are no users in the database).

**Example Output**

When you run the application, you should see output similar to this:

2023-06-13T12:45:00.123Z INFO zap/logger.go:38 Database connected successfully
2023-06-13T12:45:00.123Z INFO zap/logger.go:38 Migration did run successfully
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] GET /users --> myapp/controllers.GetUsers (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080

This indicates that the database connection was successful, migrations were applied, and the server is running on port 8080.
