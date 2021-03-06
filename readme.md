# Fetch Wordpress userInteractionCount

This is just a simple script I wrote to fetch the "userInteractionCount" from the plugin pages. 
It is made in a way that you can pipe a file with plugins to this script, or interactively use it when ran.

**pipe input**

![Pipe example](https://raw.githubusercontent.com/wiardvanrij/Fetch-Wordpress-userInteractionCount/master/img/pipe-example.gif)

**Direct input**

![Live example](https://raw.githubusercontent.com/wiardvanrij/Fetch-Wordpress-userInteractionCount/master/img/live-example.gif)

## tl;dr
Just check the releases for your flavor and go.

## Workings

It requires the **slug** of the plugin. See test.txt for examples.

## Usage

For example:  `cat test.txt | ./main-darwin-amd64`

The other option is to start the program, in which you just type or paste your slugs.

## Builds

I recommend just using `go run main.go` or build it yourself. 
Though you can find the releases here: https://github.com/wiardvanrij/Fetch-Wordpress-userInteractionCount/releases


## Support

None. :D But you can try to contact me on Twitter or via an issue here.