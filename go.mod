module go-facebook-bot

go 1.18

// go 1.17

require (
	github.com/valyala/fasthttp v1.31.0
	golang-messenger-chatbot/api v0.0.0-00010101000000-000000000000
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
)

replace golang-messenger-chatbot/api => ./api
