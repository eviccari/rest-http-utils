# REST HTTP Utils

Helper package to provide some features that handle common operations when devs must implements REST APIs with [golang](https://go.dev) using [net/http](https://pkg.go.dev/net/http#pkg-index) package.

## Dependencies
* Golang 1.19
* gopkg.in/go-playground/validator.v9 v9.31.0 (**direct**)


## Main Features

* Provides basic HTTP Error structures to map API errors uniformly 
* Functions that writes JSON messages into http.ResponseWriter object
* Functions that maps structures into JSON String and map[string]interface{} 
* Validate function to improve entity validations

<br></br>
# How to...

## **httperrors package**


```golang
// Sample of HTTP Errors utilization 
package main

import (
	"errors"
	"fmt"

	"github.com/eviccari/rest-http-utils/httperrors"
)

func main() {
	err := errors.New("something bad here")
	he := httperrors.NewBadRequestError(err)

	message := fmt.Sprintf("Error Message: %s \nStatus Code: %d", he.Message, he.StatusCode())
	fmt.Println(message)
}
```

The output looks like the follow:

```bash
Error Message: something bad here 
Status Code: 400
```

Actually httperrors supports the follow error types:

- Internal Server Error
- Bad Request Error
- Forbidden Error
- Unauthorized Error
- Not Found Error
- Unprocessable Entity Error





