# api-students
API to manage 'Golang do Zero' course students

Routes:
- GET /students - List all students
- POST /students - Creat student
- GET /students/:id -  Get infos from a specific student
- PUT /students/:id - Update student
- DELETE /student/:id - Delete student


Struct Student:
- Name (string)
- CPF (int)
- Email (string)
- Age (int)
- Active (bool)