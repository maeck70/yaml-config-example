# yaml-config-example

## YAML Configuration Loader Example

This project demonstrates how to use the `yamlconfig` package ([invalid URL removed]) to load and access configuration data from a YAML file.

**Features:**

- Leverages the `yamlconfig` package for convenient YAML configuration loading.
- Defines custom structs (`myConfig_t`, `RabbitMQ_t`) to represent configuration data.
- Provides a `main` function to load the configuration and access its values.

**Prerequisites:**

- Go installed on your system.
- The `yamlconfig` package installed. Use `go get github.com/maeck70/yaml-config` to install it.

**Usage:**

1. **Create a YAML Configuration File:**

   - Create a file named `example.config.yaml` (or any desired name) with the following structure:

   ```yaml
   Name: "My Application"
   City: "New York"
   State: "NY"
   Id: 1234
   Options:
     - "Option 1"
     - "Option 2"
   Rabbitmq:
     main:
       host: "localhost"
       port: 5672
       user: "guest"
       password: "guest"
       vhost: "/"
     baskets:
       host: "localhost"
       port: 5672
       user: "baskets_user"
       password: "baskets_password"
       vhost: "/baskets"
   ```

2. **Run the Program:**

   - From the command line, navigate to the directory containing this code (`main.go`) and the `example.config.yaml` file.
   - Run the program using `go run main.go`.

**Output:**

   ```
   new conf: &{Name:My Application City:New York State:NY Id:1234 Options:[Option 1 Option 2] Rabbitmq:map[string]main.RabbitMQ_t{main:{host:localhost port:5672 user:guest password:guest vhost:/} baskets:{host:localhost port:5672 user:baskets_user password:baskets_password vhost:/baskets}}}
   Name: My Application
   Rabbitmq main vhost: /
   Rabbitmq baskets vhost: /baskets
   ```

   The output displays the loaded configuration data, demonstrating successful reading of values from the YAML file.

**Explanation:**

- The code defines two custom structs:
   - `myConfig_t`: This struct represents the overall configuration, including `Name`, `City`, `State`, `Id`, `Options`, and `Rabbitmq` (a map of named `RabbitMQ_t` structs).
   - `RabbitMQ_t`: This struct represents RabbitMQ connection details, with fields for `host`, `port`, `user`, `password`, and `vhost`.

- The `main` function uses the `yamlconfig.LoadConfig` function to load the configuration from `example.config.yaml` and store it in a variable of type `interface{}`.
- The loaded configuration is then type-asserted to the `myConfig_t` struct to access its fields.
- The program then prints various values from the configuration to demonstrate its functionality.

**Customization:**

- You can modify the structure of the YAML file and the corresponding structs to match your specific configuration requirements.
- The `yamlconfig` package may offer additional options for customizing behavior; refer to its documentation for further details.

**Additional Notes:**

- While not explicitly shown in this example, error handling is essential in real-world applications. Consider using appropriate error handling mechanisms when working with the `yamlconfig` package.

This README provides a clear explanation of the code's functionality, demonstrates usage with an example configuration file, and offers guidance for customization and error handling.