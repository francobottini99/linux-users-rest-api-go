# REST API in Go

The program creates two web services using the Gin framework: a user service and a processing service. These services run on different ports and have different routes and associated controllers. The program also configures log files and redirects the standard output and error output to these files.

### Authors:
- **Bottini, Franco Nicolas**

### How to compile?

You can compile and run the program using the Makefile as follows:

```bash
$ git clone https://github.com/francobottini99/APIRESTGO-2023.git
$ cd APIRESTGO-2023
$ sudo make install
```

This generates the binary and runs it as a service using `systemd`, as well as configuring the `nginx` program to redirect requests to the appropriate web services.

> [!NOTE]
> To compile the project, you need to have `Go`, `systemd`, and `nginx` installed on the system.

To verify that the program ran successfully, you can use the command:

```bash
$ systemctl status lab6
```

Finally, if you want to uninstall the program and remove all generated files, you can use the command:

```bash
$ sudo make uninstall
```

## Services

The web services run on ports `8555` and `85556`, and they can be accessed via the domains `dashboard.com` and `sensors.com` respectively. To use the endpoints exposed by the services, a `JWT` token must be sent in the "Authorization" header of the request. The token can be obtained by logging into the user service.

### User Service

The user service runs on port `8555` and can be accessed through the domain `dashboard.com`. It exposes three endpoints that allow creating new users, logging in, and getting a list of existing users. The endpoints are: `api/users/createuser`, `api/users/login`, and `api/users/listall`.

#### **api/users/login**

This endpoint allows validating an existing user in the system. To do so, send a *POST* request with the following format:

```json
{
    "username": "user",
    "password": "pass"
}
```

If the credentials match a user in the system, the response will include a `JWT` token that allows using other endpoints on the server.

```json
{
    "message": "Login successful !",
    "token": "token jwt"
}
```

#### **api/users/createuser**

This endpoint allows creating a new user in the system. To do so, send a *POST* request with the following format:

```json
{
    "username": "user",
    "password": "pass"
}
```

If the creation is successful, the response will contain the information of the newly created user.

```json
{
    "id": 1000,
    "username": "user",
    "create_at": "2023-06-08 10:03:15"
}
```

Users created via this endpoint are enabled to access the server via `ssh`.

#### **api/users/listall**

This endpoint allows retrieving a list of all registered users in the system. To do so, send a *GET* request.

The response will include a list with the information of all registered users.

```json
[
    {
        "id": 1000,
        "username": "user",
        "create_at": "2023-06-08 10:03:15"
    },
    {
        "id": 1001,
        "username": "user1",
        "create_at": "2023-06-08 10:03:15"
    }
]
```

### Processing Service

The processing service runs on port `8556` and can be accessed via the domain `sensors.com`. It exposes two endpoints that allow submitting sensor information and retrieving sensor data. The endpoints are: `api/sensors/submit` and `api/sensors/summary`.

#### **api/sensors/submit**

This endpoint allows submitting sensor information. To do so, send a *POST* request with the following format:

```json
{
    "processing": 0.235,
    "free_memory": 546,
    "swap": 32532
}
```

If the submission is successful, the response will include a confirmation message.

```json
{
    "message": "Processing submitted !"
}
```

#### **api/sensors/summary**

This endpoint allows retrieving sensor information. To do so, send a *GET* request.

The response will include a list with sensor information.

```json
[
    {
        "id": 1,
        "processing": 0.235,
        "free_memory": 546,
        "swap": 32532
    },
    {
        "id": 2,
        "processing": 0.2443,
        "free_memory": 3432,
        "swap": 3532
    }
]
```

> [!NOTE]
> No authentication token is required to use this endpoint.
