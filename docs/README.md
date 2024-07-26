
<h1 style="color:dodgerblue; font-size:36px;">Api Documentation</h1>

<h2 style="color:dodgerblue;"> Overview</h2>

<h4 style="color:white;">This API provides endpoints to retrieve repository information and commits stored in the database. It also provides 
endpoints to set the 
credentials of the github repository to monitor and retrieve repositories by language, stars, and name.
Github API is used to retrieve the repository information and commits.</h3>

<h3 style="color:dodgerblue;"> Endpoints </h3>

#### 1) Update Settings.

<p style="color:green; font-style: normal;text-decoration: underline; font-weight: bold;"> Patch localhost:8086/v1/settings </p>

update the settings for the cron jobs. Set the repo to monitor and time since for commits.
After the settings are updated, a database reset is perform to delete records in the database newer than the since date.

**Request Body:**

```Go,
type SettingsPayload struct {
Owner string `json:"owner" validate:"required"`
Repo  string `json:"repo" validate:"required"`
Since string `json:"since" validate:"required"`
}
```
```json 
{
    "owner" : "repo owner",
    "repo" : "repo name",
    "since" : "2022-01-01",
    "per_page": 40
}
```


**Response:**
- `200 OK`: repo credentials set successfully.
<p style="color:green; font-style:normal; font-weight: bold">Sample Response </p>
  
```json
  {
  "message": "settings updated successfully",
  "status": 200
}
  ```

####  2) Retrieve repo by name.

<p style="color:green; font-style:normal; font-weight: bold; text-decoration: underline;"> GET localhost:8086/v1/repo/{name}


<p> Retrieve repo by name gets the repo by the repo name </p>

**Request Parameters:**

- `name` (path parameter, required): The Name of the repo.

**Response:**

- `200 OK`: repo retrieved successfully.
<p style="color:green; font-style:normal; font-weight: bold"> Sample Response </p>
  
```json
  {
     "message": "repo retrieved successfully",
    "data": {
        "id": 497859013,
        "name": "shop-arena",
        "created_at": "2022-05-30T08:42:30Z",
        "updated_at": "2022-05-30T08:54:23Z",
        "html_url": "https://github.com/Dilly3/shop-arena",
        "description": "API(Golang) multi vendor platform ",
        "language": "Go",
        "forks": 0,
        "stargazers_count": 0,
        "open_issues": 0
    },
    "status": 200
  } 
  ```
- `404 Not Found`: repo not found.
<p style="color:green; font-style:normal; font-weight: bold"> Sample Response </p>
  
```json
  {
    "message": "record not found",
    "status": 404
  }
  ```

#### 3) Retrieve commits by repo name.
<p style="color:green; font-style:normal; font-weight: bold;text-decoration: underline;">GET localhost:8086/v1/commits/{name}/{limit} </p>

<p>Retrieve commits by repo name get all the commits associated to a repo</p>

**Request Parameters:**

- `name` (path parameter, required): The name of the repo.
- `limit` (path parameter, required): The limit of the commits to retrieve.

**Response:**

- `200 OK`: repo retrieved successfully.
<p style="color:green; font-style:normal; font-weight: bold">Sample Response </p>
  
```json
  {
     "message": "commits retrieved successfully",
    "data": {
      "ID": "41791b4b3e8cb9678271b34309a024ba8870682c",
      "repo_name": "houdini",
      "Message": "Merge pull request #6 from Dilly3/repository-functions\n\nStorage functions",
      "AuthorName": "D'TechNiShan",
      "AuthorEmail": "",
      "Date": "2024-07-14T20:02:14Z",
      "URL": ""

    },
    "status": 200
  } 
  ```
#### 4) Retrieve repo by language.
<p style="color:green; font-style:normal; font-weight: bold;text-decoration: underline;">GET localhost:8086/v1/repos/{language}/{limit} </p>



**Request Parameters:**

- `language` (path parameter, required): The programming language used for the repo.
- `limit` (path parameter, required): The limit of the repos to retrieve.

**Response:**

- `200 OK`: repos retrieved successfully.
<p style="color:green; font-style:normal; font-weight: bold"> Sample Response </p>
  
```json
  {
  "message": "repo retrieved successfully",
  "data": [
    {
      "id": 828183354,
      "name": "houdini",
      "created_at": "2024-07-13T11:14:02Z",
      "updated_at": "2024-07-14T12:13:04Z",
      "html_url": "https://github.com/Dilly3/houdini",
      "description": "Houdini is a repository for github apis",
      "language": "Go",
      "forks": 0,
      "stargazers_count": 0,
      "open_issues": 0
    }
  ],
  "status": 200
  } 
  ```
#### 5) Get the top N commits authors by commit counts from the database.
 <p style="color:green; font-style:normal; font-weight: bold; text-decoration: underline;"> GET localhost:8086/v1/authors/top/{repo_name}/{limit} </p>



**Request Parameters:**

- `repo_name` (path parameter, required): The repo name you want to query.
- `limit` (path parameter, Optional): The limit of the top authors to retrieve. if not provided, the default is 10.

**Response:**
- `200 OK`: repos retrieved successfully.
<p style="color:green; font-style:normal; font-weight: bold"> Sample Response </p>
  
```json
{
  "message": "top authors by commits retrieved",
  "data": [
    {
      "author": "name",
      "commit_count": 3
    },
    {
      "author": "name",
      "commit_count": 2
    },
    {
      "author": "",
      "commit_count": 1
    }
  ],
  "status": 200
}
  ```