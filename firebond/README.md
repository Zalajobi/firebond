### PostgreSQL Installation And Database Creation

If your project requires a PostgreSQL database, follow these steps to install it on macOS:

1. Install Homebrew (if not already installed):
    - Open Terminal.
    - Run the following command:
      ```
      /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
      ```

2. Install PostgreSQL using Homebrew:
    - Open Terminal.
    - Run the following command:
      ```
      brew install postgresql
      ```

3. Start and enable PostgreSQL service:
    - Open Terminal.
    - Run the following command to start the PostgreSQL service:
      ```
      brew services start postgresql
      ```

4. Verify the installation:
    - Open Terminal.
    - Run the following command to connect to the PostgreSQL server:
      ```
      psql -U postgres
      ```
      Note: By default, the username is "postgres".

    - If the connection is successful, you should see the PostgreSQL command prompt:
      ```
      psql (version)
      Type "help" for help.
      postgres=#
      ```

    - To exit the PostgreSQL command prompt, type:
      ```
      \q
      ```
      
5. Create Table
    - Open Terminal
    - Run the command to create the table
    ```
    createdb firebond
   ```

### Setup Infura Credentials
1.  Go to the Infura website: https://infura.io.
2. Click on the "Get Started for Free" button or "Get Started" in the top-right corner.
3. Sign up for an Infura account by providing your email address and creating a password. You can also signup via your logged in gmail/email account
4. After signing up, you'll be redirected to the Infura dashboard. Click on the "CREATE NEW API KEY" button.
    - On the network, select "Web3 API"
    - Name input button, enter project name
5. Once your project is created, you'll be redirected to the project settings page. Click on endpoint and make sure the ethereum dropdown is set to the mainnet - This is essential
6. Copy Your API KEY on this page

### Environment Variables
1.  Open Base Project Folder and add the following values and set environment variable
   - DEV_INFURIA_API_KEY='Your Ethereum API key from "Setup Infura Credentials"
   - DEV_PORT=:9001
   - BASE_URL=/api/v1/firebond/pricing
   - DEV_DATABASE_NAME=firebond
   - 