## HOUDINI v1.0.0
Houdini is a simple REST API that retrieves information about repositories and its commits from GitHub.
The API is built with Go, Docker,Redis,and Postgres. The API retrieves repository and its commits from GitHub based on the settings provided and 
stores them in a Postgres database. 
The settings {repo owner , repo name, perPage count and since date } are persisted in the redis cache. 
The API also has a cron job that runs at interval to check for commits and update the data in the database.

### Structure 
```Go,
houdini/
│
├── cmd/
│   └── main.go
| 
│
├── internal/
│   ├── config/
│   │   └── config.go
│   └── github/
│       └── commits.go
│       └── github_interactor.go
│       └── helper.go
│       └── repo.go
    └── models/
         └── commit.go
         └── repo.go
         └── author.go
    
    └── repository/
    
            └── cache/
              └── ICache.go
          └── ICommit_.go
          └── IRepo_.go
          └── IStore.go
        
    └── server/
│
├── pkg/
│   └── github/
│       └── get_repo.go
│       └── github_client.go
│       └── http.go
│       └── list_commits.go
│       └── model.go
└   └── cron/ 
│       └── cron.go
├── storage/
│   └── postgres/
│       └── commit_store.go
│       └── repo_store.go
        └── postgres.go
        
│   └── redis/
│       └── redis_client.go
│
├── .env
├── README.md
└── go.mod
```

### env example
<p> Find the  <a style="color: coral; font-size: 20px;"> .env.example </a> file in the root directory to create a .env file </p>

## Installation
- <p style="color: red; font-weight: bold;"> To install the project, you need docker running on ur machine </p>
- <p> Clone the project from the repository <a href="https://github.com/Dilly3/houdini">Github</a> </p>
- <p> Run the command `go mod tidy` </p>
- <p> Create a .env file in the root directory of the project , use the env example above </p> 
- <p style="color: red; font-weight: bold;"> *** Add the GitHub token before you run the app </p>
- <p> How to create gitHub token <a href="https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-personal-access-token-classic">Token</a></p>
- <p> After you have created your token, insert it in the env file</p>
- <p> Now you are ready to run the app</p>

## Run App
<p> Initially , to build the App, run the command <a style="color: coral; font-size: 20px;"> make up-build </a> to build and run the app </p>
<p> Subsequently, to run the app, run <a style="color: coral; font-size: 20px;"> make up </a> to spin up the app. This runs the app without a fresh 
build </p>

## Test  
<p> To run the test, run the command <a style="color: coral; font-size: 24px;"> make test</a></p>

#### API Features
```Go,
- Update Settings
- Retrieve repositories by language
- Retrive top authors by commit count
- Retrieve commits by repository name
- Retrieve repository by name
```
<p style="color: yellow; font-weight: bold;"> *** You can read more on the API documentation in the docs folder in the root directory </p>
### Author
<p> Name: Anikammadu Michael  </p>
