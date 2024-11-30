# Aligator

Aligator is a CLI tool that aggregates RRS feeds, and is built in Go.

Aligator allows users to:

 * Add RSS feeds from across the internet to be collected
 * Store the collected posts in a PostgreSQL database
 * Follow and unfollow RSS feeds that other users have added
 * View summaries of the aggregated posts in the terminal, with a link to the
 full post

This project was a guided project done for [Boot.dev](www.boot.dev).

## Prerequisites

### 1. Install Go 1.23 or later

Aligator requires a Golang installation in order to work. There are two options:

**Option 1**: [webi](https://webinstall.dev/golang/). Just run this command in your
terminal and follow any instructions in the output:

  ```
  curl -sS https://webi.sh/golang | sh
  ```

**Option 2**: [Download and Install](https://go.dev/doc/install)

After installation run `go version` on your command line to make sure everything
is working properly.

### 2. Install PostgreSQL v15+ and Setup Database

Aligator uses PostgreSQL for its database storage. You can download it [here](https://www.postgresql.org/download/).
Choose your OS and follow the instructions to get started.

Once you have completed the installation, run the following and make sure you're
on version 15+:

  ```
  psql --version
  ```

*If you're on Linux run the following after verifying version -- Choose a password you won't forget.*:

  ```
  sudo passwd postgres
  ```

Start a postgres server in the background:

  * Mac: `brew services start postgresql`
  * Linux: `sudo service postgresql start`

Connect to the server. I used `psql` for the project, but feel free to use something like 
[PGAdmin](https://www.pgadmin.org/).

  * Mac: `psql postgres`
  * Linux: `sudo -u postgres psql`

Enter the following:

```
CREATE DATABASE gator;
\c gator
ALTER USER postgres PASSWORD 'postgres';
SELECT version();
```

If everything is working correctly, you should be seeing your version of PostgreSQL.
You can type exit to leave the leave the psql shell and continue to the next step.

## Installation

### 1. Clone Git Repository

Next thing you will want to do is clone the git repository somewhere to your local
machine. Run this command from the project directory you plan to develop in:

  ```
  git clone git@github.com:kairos4213/aligator.git
  ```

### 2. Set Up Config File

Aligator is a multi-user CLI application. There's no server (other than the database),
so it's only intended for local use. 

To get started, you will need to create a config file in your home directory.
Name it ".gatorconfig.json" and put your database connection string into a json
field labeled "db_url" with ssl mode query set to disable. It will look something
like this:

  ```
  {
    "db_url": "protocol://username:password@host:port/gator?sslmode=disable"
  }
  ```

### 3. Run Database Migrations

Aligator uses [Goose](https://github.com/pressly/goose) for it's migration tool.
Install Goose with the following command:

  ```
  go install github.com/pressly/goose/v3/cmd/goose@latest
  ```
Run `goose -version` to make sure it installed correctly.

In the main project directory, there is a `sql/schema` directory. Change into this
directory and run:

  ```
  goose postgres <connection_string> up
  ```

This should set up your database properly, but if for some reason you're having
issues, or just want to run the down migrations to an earlier version of the database
use the down migration command as desired:

  ```
  goose postgres <connection_string> down
  ```

### 3. Install Aligator

To install aligator where you will be able to run it's commands from any directory
on your machine, run this in your terminal:

  ```
  go install github.com/kairos4213/aligator@latest
  ```

## Usage

### register
Adds a new user to the database
  ```
  $ aligator register <username>
  $ aligator register bob
  ```

### login
Sets current user
  ```
  $ aligator login <username> 
  $ aligator login bob
  ```

### users
Lists all users in the database
  ```
  $ aligator users 
  ```

### addfeed
Adds a feed to the database
  ```
  $ aligator addfeed <feedname> <url>
  $ aligator addfeed "Boot.dev Blog" "https://blog.boot.dev/"
  ```

### feeds
Prints all feeds in database
  ```
  $ aligator feeds
  ```

### follow
Allows user to follow a feed
  ```
  $ aligator follow <url>
  $ aligator follow "https://blog.boot.dev/"
  ```

### following
Prints all the feeds a user is following
  ```
  $ aligator following
  ```

### unfollow
Allows user to unfollow a feed
  ```
  $ aligator unfollow <url>
  $ aligator unfollow "https://blog.boot.dev/"
  ```

### agg
Begins to aggregate feeds the user is following and save posts to database. Meant to be ran in background. Time between requests on aggregator is set with a duration string format.
  ```
  $ aligator agg <time_between_requests>
  $ aligator agg 2m
  ```

### browse
Returns the specified number of posts to view from feeds the user follows.
  ```
  $ aligator browse <num_of_posts>
  $ aligator browse 10
  ```

### reset
Allows user to reset the database. Purpose is for testing/development -- and those who like to watch the world burn.
  ```
  $ aligator reset
  ```

