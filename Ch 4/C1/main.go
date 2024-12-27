package main

type Membership struct {
	Type string
	MessageCharLimit int 
}
type User struct {
	Name string
	Membership   
}

func newUser(name string, membershipType string) User {
	u := User{Name: name}
	u.Membership.Type = membershipType
	if membershipType == "premium" {
		u.MessageCharLimit = 1000
	} else {
		u.MessageCharLimit = 100
	}
	return u
}