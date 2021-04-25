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

## Tags
* Retrieve all tags. An api_key is required, so make sure you provide one with `auth` command.
* Retrieve tags that user I follows. An api_key is required, so make sure you provide one with `auth` command.

## Followers
* Retrieve my followers. An api_key is required, so make sure you provide one with `auth` command.


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

## Tags
* Retrieve all tags. An api_key is required, so make sure you provide one with `auth` command.
    * `./devto tags`
* Retrieve tags that user I follows. An api_key is required, so make sure you provide one with `auth` command.
    * `./devto tags follows`

## Followers
* Retrieve my followers. An api_key is required, so make sure you provide one with `auth` command.
    * `./devto followers`

# TODO
* Listings
* Organizations
* Podcast Episodes
* Reading List
* Users
* Videos
* Webhooks
* Profile Images

1. In  some cases the returned data could be cleaned up.
2. In the case of a single article we could provide some kind of `read mode`. Something similar to the command
less on Linux.
3. REFACTOR, REFACTOR AND REFACTOR.
