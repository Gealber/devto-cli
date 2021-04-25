# Devto a cli for dev.to

This is a work in progress at the momment only available:
* Get articles
* Get articles by username
* Get articles by ID
* Get articles with videos
* Create article
* Update article

# Todo

A lot of stuff.

# Build

First make sure you have golang installed

`go build -o devto`

# Commands

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
