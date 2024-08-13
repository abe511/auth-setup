
## a simple auth setup in go and typescript


### Clone this repo

`git clone git@github.com:abe511/auth-setup.git`

`cd auth-setup`

### Start the backend and the db

`docker compose up`


### Run the client
`cd client`

`npm install`

`npm run build`

`npm run preview`


### Open a browser:
go to `http://localhost:4173/`

---

## Usage:

click `Login`

use one of these users' credentials:


email: user1@server.net, password: Pa$$word1


email: user2@server.net, password: Pa$$word2


email: user3@server.net, password: Pa$$word3


 ---

### Uninstall

#### remove containers:
`docker rm auth-server postgres`

#### remove images:
`docker rmi auth-server`

`docker rmi postgres`