package passwords

import (
	"log";
	"fmt";
	"regexp";
	"github.com/vzglad-smerti/password_hash";
)

func Password_hash( pass string ) string {
	//I can't break up with PHP xd
	hash, err := password.Hash(pass)
	if err != nil {
		panic(err)
	}

	return hash
}

func Password_verify( hash string, pass string ) bool {
	verify, err := password.Verify(hash, pass)
	if err != nil {
		panic(err)
	}

	return verify
}

func getPwd() string {
    // Prompt the user to enter a password
    fmt.Println("Enter a password")
    // Variable to store the users input
    var pwd string
    // Read the users input
    _, err := fmt.Scan(&pwd)
    if err != nil {
        log.Println(err)
    }
    // Return the users input as a byte slice which will save us
    // from having to do this conversion later on
    return pwd
}

func Verify_password_strength (password string) bool {
	//first I gotta make sure there are no spaces in the password and shit like that
	return true
}

func Validate_email( email string ) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return bool(re.MatchString(email))
}

/*One of (_, a-zA-Z0-9) up to 15 times*/
func Validate_username( username string ) bool {
	re := regexp.MustCompile(`(?m)^[_,a-zA-Z0-9]{1,15}$`)
	return re.MatchString(username)
}

func Validate_password(password string) bool {
	re := regexp.MustCompile(`(?m)^[_,a-zA-Z0-9]{5,15}$`) //this is bad but I CBA doing anything more
	return re.MatchString(password)
}
