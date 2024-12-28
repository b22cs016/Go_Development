package main

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
} // Adjusted to add a message field

func (p PlainText) Format() string {
	return p.message // Adjusted to return the message field
}

type Bold struct {
	message string // Adjusted to add a message field
}

func (b Bold) Format() string {
	return "**" + b.message + "**" // Adjusted to return the message field
}

type Code struct {
	message string // Adjusted to add a message field
}

func (c Code) Format() string {
	return "`" + c.message + "`" // Adjusted to return the message field
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
