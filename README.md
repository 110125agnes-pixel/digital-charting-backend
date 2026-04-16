# Digital Charting Backend

## Getting Started

### Installing

1. Check first go version by running this on your terminal:
    ```
    go version
    ```
2. if Go version is not `1.26.2`, uninstall current go, then download & install: [go 1.26.2](https://go.dev/dl/go1.26.2.windows-amd64.msi)

3. Clone this repository:
    ```
    git clone https://github.com/110125agnes-pixel/digital-charting-backend.git
    ```
4. Cd to project root directory:
    ```
    cd digital-charting-backend
    ```
5. Installing the packages:
    ```
    go mod tidy
    ```


### Executing program

1. Copy & paste `.env.sample` and rename the copied file to `.env`

2. Running the server:
    ```
    air
    ```
    server url: `http://localhost:8000/`

### How to Contribute

1. Pick an issue in [Issues](https://github.com/110125agnes-pixel/digital-charting-backend/issues) tab and assign yourself.

2. Make sure you are in `main` branch in your project

3. Run this command and make sure you have not done any changes in the main branch:
    ```
    git pull origin main
    ```

3. Create a new branch:
    ```
    git branch -M issue#/your-issue-title
    ```

4. Once you are done working on your issue commit your changes:
    - Commit message format: `type: message`
    - types:
      - `feat` - Commits that add, adjust or remove a feature to/of/from the API or UI
      - `fix` - Commits that fix an API or UI bug of a preceded feat commit
      - `chore` - Commits that represent tasks like initial commit, modifying .gitignore, ...
    
5. Push your commits:
    ```
    git push origin issue#/your-issue-title
    ```

