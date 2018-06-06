# Profilepics
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Femman27%2Fprofilepics.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Femman27%2Fprofilepics?ref=badge_shield)

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



## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Femman27%2Fprofilepics.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Femman27%2Fprofilepics?ref=badge_large)