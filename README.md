# Project lotery_viking

This project is an API for a lottery system for Viking. This will make the application more scalable and easier to maintain.

## Getting Started

Based on the Makefile, you can run the following commands to build and run the application.

You need to create a .env file from the .env.example file and fill in the necessary information.

Using the Framework Gin, you need to set the environment variable GIN_MODE=release to run the application in production mode.

```bash
# .env
GIN_MODE=release
```

In development mode, the application will run with the default configuration.

```bash
# .env
GIN_MODE=debug
```

> **Note:** The application uses a Mysql database, so you need to have a Mysql instance running.

### Prerequisites

-   Go 1.2X
-   Mysql
-   Docker

## MakeFile

run all make commands with clean tests

```bash
make all build
```

build the application

```bash
make build
```

run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB container

```bash
make docker-down
```

live reload the application

```bash
make watch
```

run the test suite

```bash
make test
```

clean up binary from the last build

```bash
make clean
```

## API

For interact with the application you need to run the command from the binary file using the following command:

```bash
./main
```

This will display the following message:

```bash
Usage: go run main.go [migrate|drop|seed|serve|init]
```

### Commands

-   **init**: This command will create the .env file with the necessary information. And the folder for the kiosk_images.
-   **migrate**: This command will create the tables in the database.
-   **drop**: This command will drop the tables in the database
-   **seed**: This command will insert the initial data in the database (useful for testing)
-   **serve**: This command will start the server and the application will be available at http://localhost:8080 by default.

### Init the API

To init the application you need to run the following command:

```bash
./main init
```

This will create the .env file and the folder for the kiosk_images.
Then after filling the .env file with the necessary information, you can run the following command to create the tables in the database with the following command:

```bash
./main migrate
```

> **Note:** The application uses a Mysql database, so you need to have a Mysql instance running.

You should only use this command once, to create the tables in the database.

For convenience, you can use the **seed** command to insert the initial data in the database for testing.

## Endpoints

Security is a priority, so you need to use the _API_KEY_ or the _Mac Address_ of the kiosk to access the endpoints.

### Headers

| Key          | Value            |
| ------------ | ---------------- |
| api-key      | 123456           |
| Content-Type | application/json |
| Accept       | application/json |

### Security

| Endpoint             | Method | API_KEY | Mac Address |
| -------------------- | ------ | ------- | ----------- |
| /                    | GET    | false   | false       |
| /test                | GET    | true    | false       |
| /health              | GET    | false   | false       |
| /images              | GET    | true    | false       |
| /images/:id          | GET    | true    | false       |
| /kiosks              | GET    | true    | false       |
| /kiosks/params       | GET    | true    | true        |
| /tickets             | POST   | true    | true        |
| /tickets/:code       | GET    | true    | true        |
| /tickets/claim/:code | GET    | true    | true        |

#### Errors

All errors are in the following format:

400, 401, 403, 404, 500

```json
{
  "error": string
}
```

### General

-   **GET** /: This endpoint is used to check if the application is running. Return "Hello World"
-   **GET** /test: This endpoint is used to test if you have the correct API_KEY.
-   **GET** /health: This endpoint is used to check if the application is running.
    -   Response:
        200
    ```json
    {
    	"idle": string,
    	"in_use": string,
    	"max_idle_closed": string,
    	"max_lifetime_closed": string,
    	"message": string,
    	"open_connections": string,
    	"status": string,
    	"wait_count": string,
    	"wait_duration": string
    }
    ```

### images

-   **GET** /images: This endpoint is used to get all the images from the database.

    -   Response:
        200

    ```json
    [
    	{
      "id": number,
      "created_at": string,
      "updated_at": string,
      "url": string,
      "name": string,
      "format": string
    },
    {
      "id": number,
      "created_at": string,
      "updated_at": string,
      "url": string,
      "name": string,
      "format": string
    },
    ...
    ]
    ```

-   **GET** /images/:id: This endpoint is used to get an image by id.
    -   Response:
        200
    ```json
    {
    	"id": number,
    	"created_at": string,
    	"updated_at": string,
    	"url": string,
    	"name": string,
    	"format": string
    }
    ```

### kiosks

-   **GET** /kiosks: This endpoint is used to get all the kiosks from the database.

    -   Response:
        200

    ```json
    [
    {
      "id": number,
      "created_at": string,
      "updated_at": string,
      "name": string,
      "macadress_wifi": string,
      "macadress_ethernet": string,
      "location": string,
      "id_parameters": number
    },
    ...
    ]
    ```

    > Usefull for check all the settings of the kiosk

-   **GET** /kiosks/params: This endpoint is used to get the parameters of the kiosks from the database.

    -   Response:
        200

    ```json
    {
	    "ID": number,
	     "Name": string,
	     "MacadressWifi": string,
	     "MacadressEthernet": string,
	     "Location": string,
	     "NameLotery": string,
	     "NameCasino": string,
	     "DateStart": string,
	     "DateEnd": string,
	     "Status": string [scan, draw],
	     "ClientData": boolean,
	     "Publicity": string  (url separate by ,),
	     "HomePage": string,
	     "ScanPage": string,
	     "ResultPage": string,
	     "GeneralRules": string,
	     "SpecificRules": string,
	     "Secret": string (need to be a regex), //"^[0-9]+$"
	     "SecretLength": number,
	     "UpdatedAt": string (Timestamp),
	     "UpdatedAtParameters": string (Timestamp)
    }
    ```

    > Used by the kiosk for init the application.

### tickets

-   **POST** /tickets: This endpoint is used to create a ticket from the kiosk.

     - Request:

        ```json
        {
           	"ticket_number": string, // mandatory
           	"entry_scan": string (Timestamp) , // optinal, is use for delay insert
           	"client_phone": string  // optinal, depend on the kiosk parameters
        }
        ```

        - Response:
        201
        ```json
        {
        	"message": "Ticket ajouté"
        }
        ```

        409
        ```json
		{
			"error": "Ticket déjà scanné"
		}
		```

		400
		```json
		{
			"error": "Code non valide"
		}
		```

-   **GET** /tickets/:code: This endpoint is used to get a ticket by code with all the informations.

	-   Response:
		200
	```json
		{
	  "id": number,
	  "kiosk_id": number,
	  "id_reward": number | null,
	  "ticket_number": string,
	  "client_phone": string | null,
	  "claim": boolean,
	  "entry_scan": string (Timestamp),
	  "exit_scan": string (Timestamp) | null,
	  "reward_name": string | null,
	  "big_win": boolean | null,
	  "reward_image": string | null,
		}
	```
- **GET** /tickets/claim/:code: This endpoint is used to claim a ticket by code.

	-   Response:
		200
	```json
		{
			"message": "Ticket réclamé"
		}
	```

		409
		```json
		{
			"error": "Ticket déjà réclamé"
		}
		```

		400
		```json
		{
			"error": "Code non valide"
		}
		```
