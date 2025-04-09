# Backend for Melo Music Application
This repository contains the backend source code for the music streaming service, Melo. It is built with Golang using the Echo framework and PostgreSQL through Supabase.

# Running it locally
Make sure you have git installed. You can install git using https://git-scm.com/downloads.

1. Clone the repo onto your machine
```bash
git clone https://github.com/cis444-team-1/backend
```
2. Open the project on your IDE
2. Install GoLang using https://go.dev/doc/install
3. Create a new file in the root directory and paste the variables from interal communication
```
DATABASE_URL=123
JWT_ACCESS_SECRET=123
JWT_REFRESH_SECRET=123
FRONTEND_ORIGIN=http://localhost:5173
S3_BUCKET_NAME=123
AWS_REGION=123
AWS_ACCESS_KEY=123
AWS_SECRET_ACCESS_KEY=123
AWS_CLOUDFRONT_DOMAIN=123
SUPABASE_URL=123
SUPABASE_ANON_KEY=123
```
4. Run the project in the terminal
```
go run ./cmd
```
> [!NOTE]
> You may have the re-run the project a few times if it does not work at first. Sometimes it throws an error with the TCP connection to the database. This may be because the database is put on "sleep" when it's not being actively used. So just re-run it until it works.
5. Open a browser and test the api using the url "localhost:8080"
6. You should see some text saying message not found, this means the project is working successfully.
