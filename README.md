# FloorService API
## requirements
- go version 1.17
- mysql version 8.0
### external dependencies:
- github.com/gin-gonic/gin v1.7.7
- github.com/go-sql-driver/mysql v1.6.0
- github.com/ilyakaznacheev/cleanenv v1.2.6
- github.com/onsi/gomega v1.18.1
- go.uber.org/zap v1.20.0
## Initialization using docker-compose
~~~bash
docker-compose up
~~~
server will be listening to `localhost:8000`

a swagger-ui will be available at `localhost:8080` 

## Manual Initialization
### create database user:
~~~bash
mysql -uroot -p < ./scripts/user.sql
~~~

### create database schema:
~~~bash
mysql -uroot -p ./scripts/schema.sql
~~~

### insert sample data
~~~bash
mysql -uroot -p ./scripts/sample.sql
~~~

### fetch dependencies
~~~bash
go get -t ./...
~~~

## running server
### run server:
configuration are read from env variables. check [.env.sample](.env.sample) for full list of available variables.
~~~bash
make serve
~~~
or
~~~bash
make server
source .env.sample
./floor-service
~~~

### query server:
- **get providers:**
~~~bash
curl --location --request POST 'http://localhost:8000/get_providers' \
  --header 'Content-Type: application/json'
  --data-raw '{"material":"wood", "address":{"lat":-26.66119,"long":40.95858}, "area":100, "phone_number":"1-800-2"}'
~~~
response:
~~~json
{
  "code":200,
  "message":"list of providers",
  "data":[
    {"name":"provider7","experience":["wood"],"address":{"lat":-26.66116,"long":40.95858},"operating_radius":10,"rating":4.8},
    {"name":"provider4","experience":["wood","carpet"],"address":{"lat":-26.66117,"long":40.95858},"operating_radius":10,"rating":4.7},
    {"name":"provider3","experience":["wood"],"address":{"lat":-26.66116,"long":40.95858},"operating_radius":10,"rating":4.5},
    {"name":"provider5","experience":["wood"],"address":{"lat":-26.66115,"long":40.95858},"operating_radius":10,"rating":4.5},
    {"name":"provider6","experience":["wood","tile"],"address":{"lat":-26.66118,"long":40.95858},"operating_radius":2,"rating":4.1},
    {"name":"provider1","experience":["wood","carpet","tile"],"address":{"lat":-26.66119,"long":40.95858},"operating_radius":10,"rating":3.5}
  ]
}
~~~

request format:
~~~json
[
    {
      "material":"enum['wood', 'carpet', 'tile']",
      "address": {
        "lat": "decimal",
        "long": "decimal"
      },
      "area": "decimal",
      "phone_number": "string"
    }
]
~~~
response format:
~~~json
{
  "code":"integer",
  "message":"string",
  "data":[
    {
      "name":"string",
      "experience":"enum['wood', 'carpet', 'tile']",
      "address": {"lat":  "decimal", "long": "decimal"},
      "operating_radius": "decimal",
      "rating": "decimal"
    }
  ]
}
~~~
check [OpenAPI Specifications](api/openapi.yml) for complete api documentation.

## run tests:
~~~bash
make test
~~~
