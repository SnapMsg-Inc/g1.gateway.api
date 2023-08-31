## REST specification

All methods require an `Authorization: {idtoken}` header to authenticate.
This idtoken contains the user info embedded.

### users resource

| Method | HTTP request | q-params | b-params | description | response |
|--------|--------------|:--------:|:-------:|-------------|:--------:|
| **list** | GET **`/users`** | {id, email, nick, page} | **-** | Get user(s) by the given qparams | List of {id, nick, bio, followers, followings} (if no qparams provided, empty list is returned) |
| **list recommended** | GET  **`/users/recommended`** | **-** | **-** | Get recomended user(s) for the given user | List of {id, nick, bio} |
| **create** | POST **`/users`**  | **-** | {email, nick, zone, bio} | Creates a user | **-** |
| **update** | PUT **`/users`**  | **-** | {nick, bio} | Updates information of given user | **-** |
| **delete** | DELETE **`/users`** | **-** | **-** | Delete account (of user issuing the query) | **-** |
| **follow** | POST **`/users/follow/{otherid}`** | **-** | **-** | Follows a user | **-** |
| **unfollow** | DELETE**`/users/follow/{otherid}`** | **-** | **-** | Unfollows a user | **-** |

### posts resource

| Method | HTTP request | q-params | b-params | description | response |
|--------|--------------|:-------:|:-------:|-------------|:--------:|
| **list** | GET **`/posts`** | {hashtags, nick, text, page} | **-**  | List visible posts matched by the filter | List of {nick, timestamp, text, mediaURIs} |
| **list recommended** | GET **`/posts/recommended`** | {page} | **-**  | List recommended posts for the given user, matched by the filter | List of {nick, timestamp, text, mediaURIs} |
| **create** | POST **`/posts`** | **-** | {id, hashtags, text, mediaURLs, ispublic} | Creates a post | **-** |
| **update** | PUT **`/posts/{id}`** | **-** | {text, hashtags} | Updates a post | **-** |
| **delete** | DELETE **`/posts/{id}`** | **-** | **-** | Deletes a post | **-** |

### administrator methods 
| Method | HTTP request | q-params | b-params | description | response |
|--------|--------------|:--------:|:-------:|-------------|:--------:|
| **create admin** | PUT **`/admin/users/{id}`**  | **-** | **-** | Designate an existing user as admin | **-** |
| **delete** | DELETE **`/admin/users/{id}`** | **-** | **-** | Delete a given user account | **-** |


