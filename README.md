# Golang-API

This is a REST API build using Golang.

# User Data Model

type User struct {

	Name		string				
	Dob			string				
	Address		string			
	Description	string				
	CreatedAt	time.Time			

}

# Routing
Routing is done using gorilla/mux package
# Database
MongoDB Atlas(NoSQL) is used for CRUD operations
