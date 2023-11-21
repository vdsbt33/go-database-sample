# Go Database Sample

This is a minimalistic sample on connecting to a Postgres database using Go programming language.

## Requirements

- Postgres database

## How to run

- Install dependencies
   - Run `go get database/sql github.com/lib/pq`
- Create table
   - Create the sample table
   `CREATE TABLE "Person" (
     "id" uuid NOT NULL PRIMARY KEY,
     "name" varchar(255) NOT NULL,
     "email" varchar(255) NOT NULL,
     "createdAt" timestamptz NOT NULL,
     "updatedAt" timestamptz
   );`
   - Insert a couple of rows in the database
   `INSERT INTO "Person" (
     "id",
     "name",
     "email",
     "createdAt",
     "updatedAt"
   ) VALUES (
     '3ba86131-36cf-4052-b312-11ddbaa8fe9f',
     'One',
     'one@email.com',
     now(),
     null
   );
 
   INSERT INTO "Person" (
     "id",
     "name",
     "email",
     "createdAt",
     "updatedAt"
   ) VALUES (
     '6dca7cd3-87a9-497d-8cbc-4250872c65a0',
     'Two',
     'two@email.com',
     now(),
     null
   );`
- Change the connection string in the `src/main.go` file. Point it to your database.
- Run the code with `go run src/main.go`. You should see the rows you inserted in your shell.
