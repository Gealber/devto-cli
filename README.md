# Devto a cli for dev.to

This is a work in progress so don't a expect a full support for [Dev API(beta)](https://docs.forem.com/api/).

## Articles
* Get articles
* Get articles by username
* Get articles with queries
* Get articles by ID
* Get articles with videos
* Get articles of authenticated user
* Get all articles of authenticated user
* Get published articles of authenticated user
* Get unpublished articles of authenticated user
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

## Podcasts
* Retrieve podcast availables.

## Reading Lists
* Retrieve reading lists availables.

## Webhooks. Need to be authenticated
* Retrieve webhooks they have previously registered.
* Retrieve webhooks by id
* Create webhooks
* Delete webhook

## Profile Images
* Retrieve retrieve a user or organization profile image information by its corresponding username

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
* Get articles of authenticated user
    * `./devto articles me`
* Get published articles of authenticated user
    * `./devto articles me -p`
* Get unpublished articles of authenticated user
    * `./devto articles me -up`
* Get all articles of authenticated user
    * `./devto articles me -all`
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

## Podcasts
* Retrieve podcast availables.
    * `./devto podcasts`

## Reading Lists
* Retrieve podcast availables.
    * `./devto reading_lists`

## Webhooks. Need to be authenticated
* Retrieve webhooks they have previously registered.
    * `./devto webhooks`
* Retrieve webhooks by id
    * `./devto webhooks <id>`
* Create webhooks
    * `./devto webhooks create`
* Delete webhook
    * `./devto webhooks delete <id>`

## Profile Images
* Retrieve retrieve a user or organization profile image information by its corresponding username
    * `./devto profile_images <username>`

# TODO
Finish all the endpoints. From now on is to refactor and improve user experience.

1. In  some cases the returned data could be cleaned up instead of throughing the raw json to stdout.
2. In the case of a single article we could provide some kind of `read mode`. Something similar to the command
less on Linux. Not preaty sure if spend time on this one 
3. REFACTOR, REFACTOR AND REFACTOR.
