# personal-finance-manager

## Table of contents
-   [Requirement](#requirement)
-   [How to configure on your local machine](#how-to-configure-on-your-local-machine)
-   [How to run migration](#how-to-run-migration)
-   [How to run on your local machine](#how-to-run-on-your-local-machine)

## Requirement

1. Go 1.18 or above.
2. MySQL.

## How to configure on your local machine

1. Clone this repostory to your local.

    ```bash
    $ git clone https://github.com/aririfani/personal-finance-manager.git
    ```

2. Change working directory to `personal-finance-manager` folder.

    ```bash
    $ cd personal-finance-manager
    ```
    
3. Install dependencies

    ```bash
    $ go get .
    ```

4. Create configuration files.

    ```bash
    $ cp env.toml.example env.toml
    ```

4. Edit configuration values in `env.toml` according to your setting.

## How to run migration

This migration can do these actions:

1. Migration up

    This command will migrate the database to the most recent version available. Migration files can be seen in this folder `migrations/sql/`.

    ```bash
    $ go run main.go migrate:up
    ```

2. Migration down

    This command will undo/rollback database migration.

    ```bash
    $ go run main.go migrate:down
    ```
    
## How to run on your local machine

1. Running the system.
    ```bash
    $ go run main.go
    ```
## Api documentation
You can find [api documentation here](https://rifaniari.stoplight.io/docs/personal-finance-manager)
