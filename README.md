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

# Routing is done using gorilla/mux package
# Database used : MongoDB Atlas(NoSQL)
