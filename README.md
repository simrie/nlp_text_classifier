[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/simrie/nlp_text_classifier)


## Build the Project

```
go build
```

## Test the Project's Functions

```
go test -v ./nlp

go test -v ./profile

go test -v ./utils
```


## Start Mongo

Navigate to the Mongo directory bin and start Mongo.

```
mongod
```

## Start the API Server

```
./nlp_text_classifier
```

## cUrl commands

The following assume that the API server is running locally and has successfully connected to MongoDB.


### List the names of all databases

```
curl "http://127.0.0.1:8080/databases"
```


### List all "profile" collection documents in a database

```
curl "http://127.0.0.1:8080/profiles/db/{database name}"
```

### Return a profile object from a database by its id string

```
curl "http://127.0.0.1:8080/db/{database_name}/id/{id}"
```

### Send a rawDoc object and receive (but do not store) a Profile object

```
curl -X POST --data '{\"key\":\"BabyChickenDoc\",\"text\":\"Chicks are baby chickens that are cute little chicken babies\",\"tag\":\"cuteness\"}' -H 'Content-Type: application/json' -H 'Accept: application/json'  http://127.0.0.1:8080/profile
```

The profile returned contains the minified stems and original word counts:

```json
{
	"name": "BabyChickenDoc",
	"tag": "cuteness",
	"blocks": [{
		"mini_stem": "CHICK",
		"sources": [{
			"word": "Chicks",
			"seen": 1
		}],
		"weight": 1,
		"count": 1
	}, {
		"mini_stem": "AR",
		"sources": [{
			"word": "are",
			"seen": 2
		}],
		"weight": 1,
		"count": 2
	}, {
		"mini_stem": "BABI",
		"sources": [{
			"word": "baby",
			"seen": 1
		}, {
			"word": "babies",
			"seen": 1
		}],
		"weight": 1,
		"count": 2
	}, {
		"mini_stem": "CHICKN",
		"sources": [{
			"word": "chickens",
			"seen": 1
		}, {
			"word": "chicken",
			"seen": 1
		}],
		"weight": 1,
		"count": 2
	}, {
		"mini_stem": "CUTE",
		"sources": [{
			"word": "cute",
			"seen": 1
		}],
		"weight": 1,
		"count": 1
	}, {
		"mini_stem": "LITTL",
		"sources": [{
			"word": "little",
			"seen": 1
		}],
		"weight": 1,
		"count": 1
	}]
}
```
