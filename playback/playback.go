package playback

import (
	"io/fs"
	"log"
	"math/rand"
	"os"
	"time"
  "path"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/jmbit/ctad/assets"
)

func PlayRandomSound(dir string) {
	var files []fs.DirEntry
	var err error
  var streamer beep.StreamSeekCloser
  var format beep.Format
  done := make(chan bool)
	if dir == "" {
		files, err = assets.SoundFS.ReadDir("sounds")
	} else {
		files, err = os.ReadDir(dir)
	}
	if err != nil {
		log.Println("Error getting sound files:", err)
		return
	}
	fileCount := len(files)
	randNum := rand.Intn(fileCount - 1)
	if dir == "" {
		file, err := assets.SoundFS.Open(path.Join("sounds",files[randNum].Name()))
		if err != nil {
			log.Println("Error opening file ", files[randNum].Name(), err)
      return
		}
    streamer, format, err = mp3.Decode(file)
    if err != nil {
      log.Println("Could not parse file ", files[randNum].Name(), err)
      return
    }

	} else {
		file, err := os.Open(path.Join("dir",files[randNum].Name()))
		if err != nil {
			log.Println("Error opening file ", files[randNum].Name(), err)
      return
		}
    streamer, format, err = mp3.Decode(file)
    if err != nil {
      log.Println("Could not parse file ", files[randNum].Name(), err)
      return
    }
	}
  defer streamer.Close()
  speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
  speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
