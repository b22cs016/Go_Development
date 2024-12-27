package main

type User struct {
	Name       string
	Membership Membership
}

func (user User) SendMessage(message string, messageLength int ) (string, bool) {
	if user.Membership.MessageCharLimit < messageLength {
		return "", false
	}
	return message, true
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}
