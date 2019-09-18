# pickture
visual picture selection utility

this helps you toss photos into "good" and "bad" folders

aka shots you wanna keep and shots you don't wanna keep

## Use case

I usually take about a billion photos on my camera, dump them to my laptop, and then never sort through which ones are good or bad.

Instead of:

1. Putting all of your photos in the same folder and making sure there are no filename collisions
1. Opening up a photo in fullscreen
1. Deciding whether to keep it or not
1. Finding it in the folder again, and tossing it in the trash or in a "saved" folder
1. Closing the photo browsing window
1. Repeat

You can now:
1. Toss whatever jumble of folders you have together
1. Slap open `pickture`
1. Smash some keys to toss photos in folders ("keep," "delete," "photos with dogs in them," whatever category you want)
1. Look at your nicely sorted photos!

## Installation

```bash
$ go get -u github.com/CyrusRoshan/pickture
$ cd $GOPATH/src/github.com/CyrusRoshan/pickture/
$ make
```

Now you can run `pickture` anywhere!