curl -X POST -d '{"title":"开发"}' localhost:8086/projects
curl localhost:8086/projects
curl -X PUT -d '{"title": "写bug"}' localhost:8086/projects/1
curl -X POST -d '{"title":"游戏"}' localhost:8086/projects
curl -X DELETE localhost:8086/projects/2

curl -X POST -d '{"title":"personal: golang crud package"}' localhost:8086/todos
curl localhost:8086/todos
curl -X PUT -d '{"detail":"write a package to do CRUD automatically"}' localhost:8086/todos/1

curl -X POST -d '{"ID": 1}' localhost:8086/projects/1/todos
curl localhost:8086/projects/1/todos
curl -X POST -d '{"title": "work: debug FooBar"}' localhost:8086/projects/1/todos
curl -X PUT -d '{"done": true}' localhost:8086/todos/2
curl -X DELETE localhost:8086/projects/1/todos/2
curl 'localhost:8086/todos?filter_by=done&filter_value=1'
curl -X DELETE localhost:8086/todos/2
curl  localhost:8086/todos
curl  'localhost:8086/projects?preload=Todos'
