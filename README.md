# golang-gorillamux-RESTAPI


This is Simple simgle page golang REST API Application which contain 4 sample APIs.

Command to Run the Application
------------------------------------
1) Download the Latest golang binary release from offcial golang website of google https://golang.org/dl/ based on your Platform requirements
2) Follow the install instructions from https://golang.org/doc/install
3) Check out the code from here
4) place it any location in your work machine
5) command to build the go runnable code and start the application --> "go build && http-api"

Command 5 Explanation
1) It contain total 2 commmand 1 is "go build" which is used to build the executable file based on the platform specific
2) http-api is the command to run the executable file


1st API
--------
http://localhost:8080/testjson --> GET end point Return json response "{"Status":"Application is Running"}" which indicate Application Running Status

2st API
--------
http://localhost:8080/teststring --> GET end point Return String response "Application is Running" which indicate Application Running Status

3rd API
--------
http://localhost:8080/addstring/sample text -->  POST end point to add sample text to in memory list and return the complete list of strings as a response i json format

4th API
--------
http://localhost:8080/adduser --> POST end point to add a json object to a in memory list and return the success message as "User Added Successfully"

post request body :

{
	"name":"goloang",
	"age":11,
	"useraddress" :{
		"area":"Google",
		"city":"Google",
		"pincode":"00000"
	}
}

5th API
--------
http://localhost:8080/getallusers --> GET end point to get all User details in json object