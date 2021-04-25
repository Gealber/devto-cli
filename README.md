# Devto a cli for dev.to

This is a work in progress so don't a expect a full support for [Dev API(beta)](https://docs.forem.com/api/).

## Articles
* Get articles
* Get articles by username
* Get articles with queries
* Get articles by ID
* Get articles with videos
* Create article
* Update article

## Comments
* Retrieve comments on an article or a podcast.
* Retrieve comments with its children comments.

## Tags
* Retrieve all the availables tags. An api_key is required, so make sure you provide one with `auth` command.
* Retrieve tags that I follow. An api_key is required, so make sure you provide one with `auth` command.

## Followers
* Retrieve my followers. An api_key is required, so make sure you provide one with `auth` command.

## Listings
* Retrieve listing availables.
* Create a listing.
* Update a listing.
* Reading a given listing by its id.

## Organizations
* Retrieve organization by username.
* Retrieve users on an organization. 
* Retrieve listing on an organization. 
* Retrieve articles belonging to an organization. 

# Build

`go build -o devto`

> **NOTE**: `First make sure you have golang installed`

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
* Retrieve articles by username with queries
    * `./devto articles gealber -q`
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
* Retrieve comments given an id.
    * `./devto comments -id <id>`

## Tags
* Retrieve all tags. An api_key is required, so make sure you provide one with `auth` command.
    * `./devto tags`
* Retrieve tags that I follows. An api_key is required, so make sure you provide one with `auth` command.
    * `./devto tags follows`

## Followers
* Retrieve my followers. An api_key is required, so make sure you provide one with `auth` command.
    * `./devto followers`

## Listings
* Retrieve listing availables.
    * `./devto listings`
* Create a listing.
    * `./devto listings create`
* Update a listing.
    * `./devto listings update <id>`
* Reading a given listing by its id.
    * `./devto listings retrieve <id>`

## Organizations
* Retrieve organization by username.
    * `./devto organizations <username>`
* Retrieve users on an organization. 
    * `./devto organizations <username> -u`
* Retrieve listing on an organization. 
    * `./devto organizations <username> -l`
* Retrieve articles belonging to an organization. 
    * `./devto organizations <username> -a`

# TODO
* Organizations
* Podcast Episodes
* Reading List
* Users
* Videos
* Webhooks
* Profile Images

1. In  some cases the returned data could be cleaned up instead of throughing the raw json to stdout.
2. In the case of a single article we could provide some kind of `read mode`. Something similar to the command
less on Linux. Not preaty sure if spend time on this one 
3. REFACTOR, REFACTOR AND REFACTOR.
