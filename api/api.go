package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
	// "errors"
	"strconv"

)


//Declare a global array of Credentials
//See credentials.go

/*YOUR CODE HERE*/
var credentials []Credentials



func RegisterRoutes(router *mux.Router) error {

	/*

	Fill out the appropriate get methods for each of the requests, based on the nature of the request.

	Think about whether you're reading, writing, or updating for each request


	*/

	router.HandleFunc("/api/getCookie", getCookie).Methods(http.MethodGet)
	router.HandleFunc("/api/getQuery", getQuery).Methods(http.MethodGet)
	router.HandleFunc("/api/getJSON", getJSON).Methods(http.MethodGet)

	router.HandleFunc("/api/signup", signup).Methods(http.MethodPost)
	router.HandleFunc("/api/getIndex", getIndex).Methods(http.MethodGet)
	router.HandleFunc("/api/getpw", getPassword).Methods(http.MethodGet)
	router.HandleFunc("/api/updatepw", updatePassword).Methods(http.MethodPut)
	router.HandleFunc("/api/deleteuser", deleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/api/printUsers", printUsers).Methods(http.MethodGet)

	return nil
}

func printUsers(response http.ResponseWriter, request *http.Request) {
	for _, cred := range credentials {
		fmt.Fprintf(response, "User: " + cred.Username + " Pass" + cred.Password + "\n")
	}
}

func findUser(credential Credentials) int {
	for i, cred := range credentials {
		if cred.Username == credential.Username {
			return i
		}
	}
	return -1
}

func getCookie(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "access_token" cookie's value and write it to the response

		If there is no such cookie, write an empty string to the response
	*/

	/*YOUR CODE HERE*/
	cookie, err := request.Cookie("access_token")
	if err != nil {
		fmt.Fprintf(response, "")
		return
	}
	accessToken := cookie.Value
	fmt.Fprintf(response, accessToken)
}

func getQuery(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "userID" query paramter and write it to the response
		If there is no such query parameter, write an empty string to the response
	*/

	/*YOUR CODE HERE*/
	userID := request.URL.Query().Get("userID")
	fmt.Fprintf(response, userID)

}

func getJSON(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then, write the username and password to the response, separated by a newline.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(&credential)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	username := credential.Username
	password := credential.Password
	if username == "" || password == "" {
		http.Error(response, "missing username or password", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(response, username + "\n" + password)
}

func signup(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then store it ("append" it) to the global array of Credentials.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(&credential)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if credential.Username == "" {
		http.Error(response, "missing username", http.StatusBadRequest)
		return
	}
	if findUser(credential) != -1 {
		http.Error(response, "username already used", http.StatusBadRequest)
		return
	}
	credentials = append(credentials, credential)
	response.WriteHeader(201)

}

func getIndex(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}


		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Return the array index of the Credentials object in the global Credentials array

		The index will be of type integer, but we can only write strings to the response. What library and function was used to get around this?

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(&credential)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	index := findUser(credential)
	if index == -1 {
		http.Error(response, "credential does not exist", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(response, strconv.Itoa(index))

}

func getPassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}


		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Write the password of the specific user to the response

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(&credential)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	index := findUser(credential)
	if index == -1 {
		http.Error(response, "credential does not exist", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(response, credentials[index].Password)

}



func updatePassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials.

		The password in the JSON file is the new password they want to replace the old password with.

		You don't need to return anything in this.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(&credential)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	index := findUser(credential)
	if index == -1 {
		http.Error(response, "credential does not exist", http.StatusBadRequest)
		return
	}
	credentials[index].Password = credential.Password

}

func deleteUser(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials.

		Remove this user from the array. Preserve the original order. You may want to create a helper function.

		This wasn't covered in lecture, so you may want to read the following:
			- https://gobyexample.com/slices
			- https://www.delftstack.com/howto/go/how-to-delete-an-element-from-a-slice-in-golang/

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	credential := Credentials{}
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(&credential)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	index := findUser(credential)
	if index == -1 {
		http.Error(response, "credential does not exist", http.StatusBadRequest)
		return
	}
	if credentials[index].Password != credential.Password {
		http.Error(response, "password is incorrect", http.StatusBadRequest)
		return
	}
	credentials = append(credentials[0:index], credentials[index + 1:]...)

}
