build: deps sounds
	go build -x -o ctad .

run:
	go run .

deps:
	go mod tidy
	go get .

sounds:
	wget -O assets/sounds/cat-meow-14536.mp3 https://cdn.pixabay.com/download/audio/2022/01/18/audio_4a7219f81f.mp3?filename=cat-meow-14536.mp3
	wget -O assets/sounds/cat-meow-6226.mp3 https://cdn.pixabay.com/download/audio/2021/08/04/audio_c497e4cb01.mp3?filename=cat-meow-6226.mp3
	wget -O assets/sounds/cat-purring-and-meow-5928.mp3 https://cdn.pixabay.com/download/audio/2021/08/03/audio_cbf0f95ef8.mp3?filename=cat-purring-and-meow-5928.mp3
