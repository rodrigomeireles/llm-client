# LLM Client

## Overview

This project is a personal LLM client. The website is built using HTMX and Templ for serving HTML templates in Golang. The backend is written in Go, which serves the static content and the templates for the front-end.

## Project Structure

The project is organized into the following directories and files:

- `/web`: Contains all front-end related files.
  - `/static`: Holds static files like CSS for styling.
    - `/css`: Directory for CSS stylesheets.
  - `/templates`: Contains HTML templates served by the Go backend.
  - `/images`: Directory for image files.
- `/backend`: Contains all Go backend logic.
  - `/handlers`: Route handlers that respond to HTTP requests.
  - `/models`: Data models used throughout the application.
  - `main.go`: The main entry point for the Golang application.

## How to Build

To build the project, ensure you have Go installed on your system. Navigate to the `/backend` directory and run:

```sh
go build
```

## How to Run

### Development Mode

To run the project in development mode with Tailwind CSS in watch mode, use the `run_dev.sh` script. This script will start your Go server and also start Tailwind CSS in watch mode, which automatically rebuilds your CSS file whenever you make changes to your styles.

In the project root directory, execute:

```sh
./run_dev.sh
```

Make sure to make the script executable before running it. You can do this with the following command:

```sh
chmod +x run_dev.sh
```

This will start the server and Tailwind CSS watcher. The website will be accessible at `http://localhost:8080` or the configured port.

### Production Mode

For production, use the `run.sh` script which starts the server without Tailwind CSS in watch mode, ensuring that your application is running with the optimized and compiled CSS files.

In the project root directory, execute:

```sh
./run.sh
```

Similarly, ensure the script is executable:

```sh
chmod +x run.sh
```


## How to Test

To run tests, navigate to the `/backend` directory and execute:

```sh
go test ./...
```

This command will recursively run all tests in the backend directories.

## Contributing

If you wish to contribute to this project, please follow the below steps:

1. Fork the repository.
2. Create a new branch for your feature or fix.
3. Commit your changes with clear, descriptive messages.
4. Create a pull request against the main branch.

For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT]


#TODO
- Write tests
- Better session handling
- Add model options
- Write a Dockerfile and a compose.yaml 
