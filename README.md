# Simple Stream
This is a very simple test for streaming raw PCM16 through HTTP

## run server
```
go run main.go
```

## play audio
This does not understand the `Content-Type` header but you could :)
```
ffplay -f s16le -ar 44100 http://localhost:8080/audio
```
