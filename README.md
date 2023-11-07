# rest-api

# employee
GET /employees -- list of employees -- 200, 404, 500
GET /employees/:id -- employee by id -- 200, 404, 500
POST /employees -- create employee -- 201, 4xx, Header Location: url
PUT /employees/:id -- fully update employee -- 204, 404, 400, 500
PATCH /employees/:id -- partially update employee -- 204, 404, 400, 500
DELETE /employees/:id -- delete employee by id -- 204, 404, 400