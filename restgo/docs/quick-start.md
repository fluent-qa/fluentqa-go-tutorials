# PrestD

## API 

1.***GET***
_health｜Health check endpoint
/databases ｜ List all databases
/schemas ｜ List all schemas
/tables ｜List all tables
/show/{DATABASE}/{SCHEMA}/{TABLE} ｜Lists table structure - all fields contained in the table
/{DATABASE}/{SCHEMA} ｜Lists table tables - find by schema
/{DATABASE}/{SCHEMA}/{TABLE} ｜List all rows, find by database, schema and table
/{DATABASE}/{SCHEMA}/{VIEW} ｜List all rows, find by database, schema and view

2. ***POST***
/{DATABASE}/{SCHEMA}/{TABLE}
Body:
```json
{
    "FIELD1": "string value",
    "FIELD2": 1234567890
}
```
3. PUT/PATCH: query string for filter
```shell
/{DATABASE}/{SCHEMA}/{TABLE}?{FIELD NAME}={VALUE}
{
    "FIELD1": "string value",
    "FIELD2": 1234567890,
    "ARRAYFIELD": ["value 1","value 2"]
}
```
4. DELETE
```shell
/{DATABASE}/{SCHEMA}/{TABLE}?{FIELD NAME}={VALUE}
```

## API Visiit

## Custom Query

- folder: awesome_folder/example_of_powerful.read.sql
- sql:
```sql
SELECT * FROM table WHERE name = "{{.field1}}" OR name = "{{.field2}}"
```
- api:
```shell
GET /_QUERIES/awesome_folder/example_of_powerful?field1=foo&field2=bar
```
- configuration: 

```toml
[queries]
location = /path/to/queries/
```

- examples:

```shell
queries/
└── foo
    └── some_get.read.sql
    └── some_create.write.sql
    └── some_update.update.sql
    └── some_delete.delete.sql
└── bar
    └── some_get.read.sql
    └── some_create.write.sql
    └── some_update.update.sql
    └── some_delete.delete.sql

URLs to foo folder:

GET    /_QUERIES/foo/some_get?field1=bar
POST   /_QUERIES/foo/some_create?field1=bar
PUT    /_QUERIES/foo/some_update?field1=bar
PATCH  /_QUERIES/foo/some_update?field1=bar
DELETE /_QUERIES/foo/some_delete?field1=bar


URLs to bar folder:

GET    /_QUERIES/bar/some_get?field1=foo
POST   /_QUERIES/bar/some_create?field1=foo
PUT    /_QUERIES/bar/some_update?field1=foo
PATCH  /_QUERIES/bar/some_update?field1=foo
DELETE /_QUERIES/bar/some_delete?field1=foo
```

