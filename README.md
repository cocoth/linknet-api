# LinkNet API

This is the LinkNet API, a powerful tool for managing users, tasks, etc...

## Installation

To install the LinkNet API, clone the repository and navigate to the project directory:

```sh
git clone https://github.com/cocoth/linknet-api.git
cd linknet-api
```

## Usage

To use the LinkNet API, follow these steps:

1. **Build the project**:

    ```sh
    cd src 
    go build
    ```

2. **Run the API server for the first time**:
    ```sh
    ../build/linknet-api initdb
    ```
    this **initdb** its a must when starting this program for the first time, you'll need to fill **admin user, admin email, admin password** 

3. **Access the API**:
    Open your browser or use a tool like `curl` or `Postman` to interact with the API at `http://localhost:3000`. Default port is 3000

4. **If you already** `initdb` **you can use this method for running the program**:
    ```sh
    cd linknet-api
    ./build/linknet-api serve
    ```
    ## **Warning!!:**
    `initdb` supposed to be run only at the first time, if you run it again you will lose data inside the databse. Soo.. makesure after you run the first program, use `serve` instead of `initdb`.


## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the **MIT** License.