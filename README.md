# Explore-go-migrate

Learning how to use [golang-migrate/migrate](https://github.com/golang-migrate/migrate).

Goal of this project is to create a small golang project that using the golang-migrate package for do a basic mysql database migration.

## Install [golang-migrate/migrate](https://github.com/golang-migrate/migrate) cli-tool

You should have the golang-migrate cli-tool for generating the migration files.

### build from source

clone the golang-mirate project from [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)


or run this on your terminal

```sh
GO111MODULE=on go get github.com/golang-migrate/migrate
```

After that, you should have the migrate project under your $GOPATH/src/github.com/golang-migrate/migrate.

Next, `cd` to the migrate repo. them run `make` when you are inside the golang-migrate/migrate repo directory.

```sh
cd path/to/gomigrate/dir
make
```

There should be a executable file name "migrate" generated inside the directory. 
To run the cli-tool just excute the excute the executable.

```sh
./migrate
```

## Migration (source) files

The golang-migrate/golang package will need 2 files (up and down) for every migration step. 
Those files should be in this format.
```
{version}_{title}.up.{extension}
{version}_{title}.down.{extension}
```

Example from the github repo.
```
1_initialize_schema.down.sql
1_initialize_schema.up.sql
2_add_table.down.sql
2_add_table.up.sql
...
```
You can read more about Filename format [here](https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md)

In order to have our migration excute correctly, 
we must have the ordering number in the name running correctly.
You can create them manully or use the cli-tool to help us generate them.

### Use the migrate cli-tool to generate the migration files

#### basic usage

Let's say, you want to create a new migration version for the mysql server.
You can use command below.
This migration we want to create a new table for user table.

```sh
migrate create -ext sql -dir ./migration -seq create_users_table
```

`create` is the command we use.

use `-ext` to choose your file extension sql for this case.

use `-dir` to choose the file generation target directory.

use `-seq` option to generate sequential up/down migrations with N digits.

`create_users_table` is our title for the migration

Run `migrate -help` for more about command options

In case if this is your first migration ever. This will create migration files instal ./migration directory.
Inside the directory you will file 2 file below.

```
000001_create_users_table.up.sql
000001_create_users_table.down.sql
```



