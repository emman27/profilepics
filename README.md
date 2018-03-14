# Profilepics
Profile pictures just like how Telegram does them by default, done in Go. (plus a couple more colors, cos they look better that way)

### Installing and using
HTTP is served by default over port 8080 on the path /
It accepts a single query parameter `name`.
```
dep ensure
go run main.go
```

### Credits
Much credits go to [ajstarks' svgo](https://github.com/ajstarks/svgo), for the initial implementation (now removed)

