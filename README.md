# ctad is a joke program
ctad (or cta daemon) is a small joke program that can be run in a background (i.e. daemon) that randomly makes cat noises (or other configured noises)
Because the speaker is reset for every sound file, they can be in completely different sample rates, BUT can't overlap.

## Usage
```
$ ctad --help
Program that lets your computer meow (or play any other noise, really), in a random interval
  the delay set is in Seconds. 
  Sound files have to be mp3 and all in a single directory, 3 cat sounds are already embedded

Usage:
  ctad [flags]

Flags:
      --config string   config file (default is $HOME/.ctad.yaml)
  -h, --help            help for ctad
  -m, --max-delay int   Maximum delay between noises (default 500)
  -i, --min-delay int   Maximum delay between noises (default 30)
  -s, --sounds string   directory to use for sound files
```

## Build
To build this program, you need go 1.23.1 or later, wget and make. If you have these prerequisites, just type `make`
