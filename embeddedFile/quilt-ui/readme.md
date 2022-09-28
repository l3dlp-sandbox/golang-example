# Quilt UI

Quilt UI is an admin UI which can be used by product/marketing to configure content without using engineering effort.

## 1. Dependencies
- Make sure you save and configure `.env.dev`, `.env.prod` and `.env.static-dev` in the quilt-ui directory from another developer.
- `.env.dev` is used to load the environment variables when you are using `2.1 Local development with static files served from Quilt`
- `.env.static-dev` is used to load the environment variables when you are using `2.2  Local Development`
- `.env.prod` is used to load the environment variables when you are building the static files for deployment `3. How Quilt UI is served`

## 2. Setup

### 2.0 Setup etc/hosts
1. edit `/etc/hosts` file with admin rights
2. Add the following host paths to the file
```
127.0.0.1 	ssg-quilt.circles.local
127.0.0.1 	stw-quilt.circles.local
127.0.0.1 	sau-quilt.circles.local
127.0.0.1 	sid-quilt.circles.local
127.0.0.1 	sjp-quilt.circles.local

127.0.0.1 	qsg-quilt.circles.local
127.0.0.1 	qtw-quilt.circles.local
127.0.0.1 	qau-quilt.circles.local
127.0.0.1 	qid-quilt.circles.local
127.0.0.1 	qjp-quilt.circles.local

127.0.0.1 	psg-quilt.circles.local
127.0.0.1 	ptw-quilt.circles.local
127.0.0.1 	pau-quilt.circles.local
127.0.0.1 	pidgc-quilt.circles.local
127.0.0.1 	pjp-quilt.circles.local
```

### 2.1 Local development with static files served from Quilt
Use this if you need to work on both frontend and backend logic.
1. Make sure you have obtained the .env files in the `1. Dependencies` section
2. Make sure you have installed [esc](https://github.com/mjibson/esc) `go get -u github.com/mjibson/esc`
3. Make sure you are in Quilt root directory
4. `make build-quilt-ui-dev && make run`
5. visit `http://ssg-quilt.circles.local:9991/web/`

### 2.2 Local Development
Use this if you only need to work on frontend logic.

This uses webpack dev server(which watches for file changes) for productivity. Ie developer doesnt need to keep building the Quilt app each time there's a change ie the `2.1` aproach described below.
#### Start Json-server
json-server is a simple json server. 
We use it in lieu of making api requests to the local quilt server so that we do not need to rebuild and restart the quilt server everytime we make a change to the frontend code.
1. Install [json-server](https://www.npmjs.com/package/json-server)
    `npm install -g json-server`
2. Start the server
    `json-server --watch mock-quilt.json`
3. Update `mock-quilt.json` file if you need to mock a specific quilt page

#### Start quilt-ui dev server
1. Make sure you have obtained the .env files in the `1. Dependencies` section
2. `npm run start`
3. To get past the authentication page. Create a cookie with name `isLoggedIn` and value `true`
4. Navigate to the pagelist to select a page to start development `http://ssg-quilt.circles.local:8000/web/pagelist`

## 3. Deploy Quilt UI
Quilt UI's static files(index.html and .js files) are served through Quilt's api server.
To deploy your code, follow the steps below when you are ready to check in your frontend code changes.
The make script will first build the UI code and put the output in `apps/web`
It will then run [esc](https://github.com/mjibson/esc) a go file embedder that will enable quilt to serve Quilt UI via the Quilt server

1. Make sure you have obtained the .env files in the `1. Dependencies` section
2. Make sure you have installed [esc](https://github.com/mjibson/esc) `go get -u github.com/mjibson/esc`
3. Make sure you are in Quilt's root directory
4. Run `make build-quilt-ui`
5. Check in the `apps/web/web.go` file into the git repo