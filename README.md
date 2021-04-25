# Devto a cli for dev.to

This is a work in progress:

## Articles
* Get articles
* Get articles by username
* Get articles with queries
* Get articles by ID
* Get articles with videos
* Create article
* Update article

## Comments
* Retrieve comments on article and podcast
* Retrieve comment with its children comments

# Todo

A lot of stuff.

# Build

First make sure you have golang installed

`go build -o devto`

# Commands

## Articles
* Store api_key
    * `./devto auth <api_key>`
* Retrieve articles
    * `./devto articles`
* Retrieve articles with queries
    * `./devto articles -q`
* Retrieve articles by username
    * `./devto articles gealber`
* Retrieve articles by id
    * `./devto articles <id>`
* Retrieve articles with videos
    * `./devto articles videos`
* Create an article
    * `./devto articles create`
* Update a given article
    * `./devto articles update <id>`

## Comments

* Retrieve comments on article and podcast
    * `./devto comments -a_id <id>`
    * `./devto comments -p_id <id>`
* Retrieve comments on article and podcast
    * `./devto comments -id <id>`

# TODO

1. In  some cases the returned data could be cleaned up.
2. In the case of a single article we could provide some kind of `read mode`. Something similar to the command
less on Linux.
3. REFACTOR, REFACTOR AND REFACTOR.
