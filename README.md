[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/simrie/nlp_text_classifier)


## Build the Project

go build

## Test the Project's Functions

go test -v ./nlp

## Start Mongo

Navigate to the Mongo directory bin and start Mongo.

> mongod

## Start the API Server

> ./nlp_text_classifier

## cUrl commands

The following assume that the API server is running locally and has successfully connected to MongoDB.


### List the names of all databases

curl "http://127.0.0.1:12345/databases"


### List all "profile" collection documents in a database

curl "http://127.0.0.1:12345/profiles/db/{database name}"

### Return a profile object from a database by its id string

curl "http://127.0.0.1:12345/db/{database_name}/id/{id}"

### Send a rawDoc object and receive (but do not store) a Profile object


curl -X POST --data '{\"key\":\"BoopsieDoc\",\"text\":\"Baby Chickens cute chickens that a little Babies\",\"tag\":\"cuteness\"}' -H 'Content-Type: application/json' -H 'Accept: application/json'  http://127.0.0.1:12345/profile


