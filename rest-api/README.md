Build a simple backend API for notes.

Use:
net/http
JSON request/response
Routing
CRUD design
Basic API testing with Postman or curl

Endpoints:

GET    /notes
POST   /notes
GET    /notes/{id}
PUT    /notes/{id}
DELETE /notes/{id}

Use an in-memory map first:

map[int]Note

Save notes to a JSON file or SQLite.
