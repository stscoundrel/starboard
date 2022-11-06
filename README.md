# Starboard

List stars user has received in their Github repositories. A Golang library.

Also available as an [API](https://github.com/stscoundrel/starboard-gin).

## Install

`go get github.com/stscoundrel/starboard`

## Usage

Troubleman exposes a function for getting Github stars by username. The request goes through public search API, so no authentication is required.

```go
package main

import (
    "fmt"

    "github.com/stscoundrel/starboard/stars"
)

func main() {
    // Get list of all of users repos that have received stars
    stars, err := stars.GetStars("stscoundrel")
    
    fmt.Println(stars[0])
    // {
    // 	Repository:     "stscoundrel/starboard",
    // 	Count:          1,
    // 	Link:           "https://github.com/stscoundrel/starboard",
    // 	StarGazersLink: "https://api.github.com/repos/stscoundrel/starboard/stargazers",
    // },
}
```
