# userutil

## What is it?
A Go library for working with users on the command line.

## API

#### User input
The library's API offers several helpful functions for working with user input.
Prompting for user input can be configured using a `PromptOptions` struct.
Passing an empty `PromptOptions` will use the default options.

For example, you can prompt a user to answer a question:
```go
package main

import (
	"log"

	"github.com/stephen-fox/userutil"
)

func main() {
	result, err := userutil.GetUserInput("How is your day going?", userutil.PromptOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("You said: '" + result + "'")
}
```

If you wanted to hide user input, you can specify so using the `PromptOptions`:
```go
package main

import (
	"log"

	"github.com/stephen-fox/userutil"
)

func main() {
	options := userutil.PromptOptions{
		ShouldHideInput: true,
	}
	result, err := userutil.GetUserInput("Say something nice", options)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("You said: '" + result + "'")
}
```

You can also prompt for a "yes or no" answer:
```go
package main

import (
	"log"

	"github.com/stephen-fox/userutil"
)

func main() {
	answeredYes, err := userutil.GetYesOrNoUserInput("Is today a nice day?", userutil.PromptOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}

	if answeredYes {
		log.Println("You said yes")
	} else {
		log.Println("You said no")
	}
}
```

#### Admin permissions
You can also check if the current user is `root` (or an `Administrator`):
```go
package main

import (
	"log"

	"github.com/stephen-fox/userutil"
)

func main() {
	err := userutil.IsRoot()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("You are root")
}
```
