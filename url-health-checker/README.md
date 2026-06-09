Build a program that checks whether websites are online.

Use:
HTTP requests
Goroutines
Channels
Timeouts
Basic concurrency patterns

go run main.go https://google.com https://openai.com https://github.com

https://google.com -> 200 OK
https://github.com -> 200 OK
https://fake-site.xyz -> failed
