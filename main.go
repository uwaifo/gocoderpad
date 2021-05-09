// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, World!")
	}
}

/*
Your previous Markdown content is preserved below:

# Business Problem

We aggregate data from many sources in order to determine and process user permissions and state.  For this exercise, a user is represented as:

```go
type User struct {
    ID          string                 `json:"id"`
    Name        string                 `json:"name"`
    Email       string                 `json:"email"`
    Permissions map[string]bool        `json:"permissions"`
    Preferences map[string]interface{} `json:"preferences"`
    Documents   []string               `json:"documents"`
}
```

The user data comes from the following sources:

- a `user` service, where the user `ID`, `Name`, and `Email` are fetched based on the user's login credentials (email address)
- a `permissions` service, where permissions for a user are fetched by their user `ID`
- a `preferences` service, where user preferences are fetched by their user `ID`
- a `documents` service, where user document uris are fetched by their creator (user) `ID`

The permissions, preferences, and documents services have no dependencies on each other and may be run concurrently.  However, each of these calls require a user ID, which can only be obtained from the user service.

Ideally all calls to each service will complete with no issues, but occasionally, calls will fail due to timeouts, network issues, or service interruptions.  When that happens, we would still like to return user data, albeit in a degraded state.  When we know we are going to return a degraded user struct, we use this error mechanism:

```go
type UserStateErr struct {
    State   int
    stateMX sync.Mutex
}

const (
    NoPermissions = 1 << 0
    NoPreferences = 1 << 1
    NoDocuments   = 1 << 2
)

// Error satisfies the error interface
func (e *UserStateErr) Error() string {
    sb := &strings.Builder{}
    sb.WriteString("the following properties are missing:")

    if e.State&NoPermissions == NoPermissions {
        sb.WriteString(" [permissions] ")
    }
    if e.State&NoPreferences == NoPreferences {
        sb.WriteString(" [preferences] ")
    }
    if e.State&NoDocuments == NoDocuments {
        sb.WriteString(" [documents] ")
    }

    return sb.String()
}

// Combine aggregates two errors
func (e *UserStateErr) Combine(other *UserStateErr) {
    e.stateMX.Lock()
    e.State |= other.State
    e.stateMX.Unlock()
}
```

## Create a service that

- Satisfies the following interface:

  ```go
  type UserRetriever interface {
    GetUserByEmailAddress(email string) (*User, error)
  }
  ```
  const GetUserByEmailAddress = async ( user: User)=> {
  let {id,name,email,permision, prefrences ,documents} = user
  let userServiceUrl = 'http://www.mocky.io/v2/5c6ef716340000f139892fa3" <--- Don't woryr
  let

  const opts = {
    headers: {
      "Request-ID": "requestid123456"
    }
  }

  //try(

  //Get use data
  const userData = await axios.get(userServiceUrl, params({ // fine
  header: opts

  }))

  userData.Permissions = {}
  userData.Prefer.. = {}
  userData.Documents = []

  userError = UserStateErr()

  const additionalData = [
    axius.get(permisisons).then((result) => {
      userData.Permission = result
    }.catch((err) => {
      userError.combine(UserStateError(NoPermissions))
    }),
    ... prefs,
    ... documents,
  ]

  await Promise.all(additionalData)

  return userData, userError


  //Get permision
  const userPerm = await axios.get(permisionServiceUrl


  if(!persmisionData || !documentData || ! prefrencesData) {
    return {
    userProfile: userData

    }

  }
  )catch(err) {

  //Handle error here
  }


  }



  logger.Info("asdfasdfasdf")
  logger.With({
    "Request-ID": "asdfasdfasdf",
    "UserId": userData.Id,
  }).Info("adfasdfasdfasDF")


### Constraints

- If a `User` struct is returned, it **must** contain a nonzero `ID`, `Name`, and `Email`.
- If the user's permissions, preferences, or documents cannot be fetched, indicate to the caller that the user is in a degraded state
- The following header **must** accompany each outgoing request: `Request-ID`
  - for this exercise, you can make up your own request id
- Use the structured logger of your choice to log errors

### Test Data

Use the following live endpoints for http requests:

- Mock URLs
  - user service: `http://www.mocky.io/v2/5c6ef716340000f139892fa3`
  - permission service: `http://www.mocky.io/v2/5c6ef7ba340000063a892fa6`
  - preferences service: `http://www.mocky.io/v2/5c6ef98c3400005600892fae`
  - documents service: `http://www.mocky.io/v2/5c8837ae3200000f693bd6e0`
*/
