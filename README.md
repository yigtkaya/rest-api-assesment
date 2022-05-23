# rest-api-assesment

Task: Create a user and group management REST API service. Service should provide a way to list, add, modify and remove users and groups. Each user can belong to at most one group.

coded with latest GO release, using gin web-framework and mongodb for database. Swagger docs are written with OPEN API 2.0 and annotations within the code is from go-swagger package.
Containerized with Docker. 

 to get project run:
    
    $ git clone https://github.com/yigtkaya/rest-api-assesment.git

 go to project file:
   
    cd /to/file/path
 
 for test go to controller folder where endpoint test functions are:

    $ cd controller 

 Running from project: 

     $ go run main.go

Running with Dockerfile:
  
     $ docker build -t rest-api .
     $ docker run -d -p 9090:9090 -e "MONGODB_URI=mongodb+srv://IAESTE:123234@cluster0.ouanz.mongodb.net/?retryWrites=true&w=majority" rest-api  


Running directly with docker-compose:

    $ docker compose up


After app start listening on 9090 port we can reach it with these commands and urls

/CreateUser
link: http://localhost:9090/v1/user/createUser

curl command:

	$ curl -X POST -H "Content-Type: application/json" -d '{"id" : "2","email":"john.terry@gmail.com","password": "121212","name" : "john            terry","membership":{"id": "1","group_name": "Fire"}}' "http://localhost:9090/v1/user/createUser"


/GetUser
link:
http://localhost:9090/v1/user/getUser/{id}

curl command:

	$ curl -X GET -H "Content-Type: application/json" "http://localhost:9090/v1/user/getUser/2"


/GetAllUsers
link:
http://localhost:9090/v1/user/getAllUsers

curl command:

	$ curl -X GET -H "Content-Type: application/json" "http://localhost:9090/v1/user/getAllUsers"


/UpdateUser
link:
http://localhost:9090/v1/user/UpdateUser

curl command:

	$ curl -X PATCH -H "Content-Type: application/json" -d '{"id" : "2","email":"Eren.Kaya@gmail.com","password": "155155","name" : "Eren Kaya","membership":{"id": "2","group_name": "FireBender"}}' "http://localhost:9090/v1/user/updateUser"

/DeleteUser
link:
http://localhost:9090/v1/user/deleteUser/{id}

curl command:

	$ curl -X DELETE -H "Content-Type: application/json" "http://localhost:9090/v1/user/deleteUser/2"


-----------------------------------------------------------

/CreateGroup
link:
http://localhost:9090/v1/group/createGroup

curl command: 
		
	$ curl -X POST -H "Content-Type: application/json" -d '{"id": "1","group_name": "Fire"}' "http://localhost:9090/v1/group/createGroup"


/GetGroup
link:
http://localhost:9090/v1/group/getUser/{id}

curl command:

	$ curl -X GET -H "Content-Type: application/json" "http://localhost:9090/v1/group/GetGroup/2"


/getAllGroups
link:
http://localhost:9090/v1/group/getAllGroups

curl command:
	
	$ curl -X GET -H "Content-Type: application/json" "http://localhost:9090/v1/group/getAllGroups"


/UpdateUser
link:
http://localhost:9090/v1/group/UpdateGroup

curl command:
	
	$ curl -X PATCH -H "Content-Type: application/json" -d '"id": "1","group_name": "FireBender"}' "http://localhost:9090/v1/group/updateGroup"

/DeleteUser
link:
http://localhost:9090/v1/group/deleteGroup/{id}

curl command:
	
	$ curl -X DELETE -H "Content-Type: application/json" "http://localhost:9090/v1/group/deleteGroup/1"



After running, stop container with docker:

get name from output
	
	$ docker container ls 

then:
	
	$ docker stop <container_name>

with compose:

	$ docker compose stop
