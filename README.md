## REST specification

All methods require an `Authorization: Bearer <JWT>` header to authenticate with [firebase](https://firebase.google.com/docs/auth/web/start).

This idtoken contains the user info embedded.

### users methods 

| Method | HTTP request | q-params | b-params | description | response |
|--------|--------------|:--------:|:-------:|-------------|:--------:|
| **list** | GET **`/users`** | {id, email, nick, maxresults, page} | **-** | Get user(s) by the given qparams | List of {id, **fullname(*)**, nick, bio, followers, followings} (\* only if the user matched is the current user)|
| **list recommended** | GET  **`/users/recommended`** | **-** | **-** | Get recomended user(s) for the given user | List of {id, nick, bio, followers, followings} |
| **create** | POST **`/users`**  | **-** | {email, fullname, nick, zone, bio} | Creates a user | **-** |
| **update** | PUT **`/users`**  | **-** | {nick, bio} | Updates information user | **-** |
| **delete** | DELETE **`/users`** | **-** | **-** | Delete account (of user issuing the query) | **-** |
| **follow** | POST **`/users/follow/{id}`** | **-** | **-** | Follow a user | **-** |
| **unfollow** | DELETE **`/users/follow/{id}`** | **-** | **-** | Unfollow a user | **-** |

### post methods 

| Method | HTTP request | q-params | b-params | description | response |
|--------|--------------|:-------:|:-------:|-------------|:--------:|
| **list** | GET **`/posts`** | {hashtags, nick, text, maxresults, page} | **-**  | List visible posts matched by the filter | List of {pid, nick, timestamp, text, mediaURIs, likes} |
| **list feed** | GET **`/posts/feed`** | {maxresults, page} | **-**  | List the posts of the followed users (feed) | List of {pid, nick, timestamp, text, mediaURIs, likes} |
| **list recommended** | GET **`/posts/recommended`** | {maxresults, page} | **-**  | List recommended posts for the given user, matched by the filter | List of {pid, nick, timestamp, text, mediaURIs, likes} |
| **create** | POST **`/posts`** | **-** | {hashtags, text, mediaURLs, ispublic} | Creates a post | **-** |
| **update** | PUT **`/posts/{id}`** | **-** | {text, hashtags} | Updates own post | **-** |
| **delete** | DELETE **`/posts/{id}`** | **-** | **-** | Deletes own post | **-** |
| **like** | POST **`/posts/like/{id}`** | **-** | **-** | Like a post | **-** |
| **unlike** | DELETE **`/posts/like/{id}`** | **-** | **-** | Unlike a post | **-** |
| **list favs** | GET **`/posts/fav`** | {maxresults, page} | **-** | Get fav posts | List of {pid, nick, timestamp, text, mediaURIs, likes} |
| **fav** | POST **`/posts/fav/{id}`** | **-** | **-** | Mark a post as favorite | **-** |
| **unfav** | DELETE **`/posts/fav/{id}`** | **-** | **-** | Unfav a post | **-** |

### admin methods 
| Method | HTTP request | q-params | b-params | description | response |
|--------|--------------|:--------:|:-------:|-------------|:--------:|
| **create admin** | PUT **`/admin/users/{id}`**  | **-** | **-** | Designate an existing user as admin | **-** |
| **delete** | DELETE **`/admin/users/{id}`** | **-** | **-** | Delete a given user account | **-** |
| **delete** | DELETE **`/admin/posts/{id}`** | **-** | **-** | Deletes any post | **-** |

