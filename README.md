# Blogging Software

Following the [Build Your Own Blogging Software](https://codingchallenges.fyi/challenges/challenge-blog) challenge.

## How to run

1. Spin up a PostgreSQL server and create a database.
1. Create a table `posts` with the following command:
   ```
   CREATE TABLE posts (
       id SERIAL PRIMARY KEY,
       title VARCHAR(50) NOT NULL,
       summary VARCHAR(100) NOT NULL,
       body TEXT NOT NULL,
       created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
       updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
   );
   ```
1. Save the database URL in an environment variable `DATABASE_URL` (something like "postgres://user:pass@localhost:5432/dbname").
1. Run `go run main.go` and navigate to [localhost:8090](http://localhost:8090) to use the service.
