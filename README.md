# arcgonaut

![License](http://img.shields.io/:license-mit-blue.svg)]

Software for creating arc diagrams such as:

![Airports](airports.png)

arcgonaut is written in Go. It uses [Laurent Le Goff](https://github.com/llgcode)'s [draw2d](https://github.com/llgcode/draw2d) for the imaging tools and [Lucas Beyer](https://github.com/lucasb-eyer)'s [go-colorful](https://github.com/lucasb-eyer/go-colorful) for generating colors. The rest is a lot geometry/rescaling.

In this example the list in the middle corresponds to country names, the arc between names corresponds to the amount the first name is "sending" to the second name.

## Setup

You can install Go [here](https://golang.org/doc/install).

```sh
go get https://github.com/MaxHalford/arcgonaut
cd $GOPATH/src/github.com/MaxHalford/arcgonaut
go get
go build
```

## Usage

For the moment the tool is usable in the command-line.

- Naviguate towards the directory where the binary is.
- ``./arcgonaut -f=data/example.arcgo -c1=#ffc3e1``

![Example](example.png)

A PNG will be generated into the same directory. The program takes as input a file that has a [specific format](example.arcgo). The file extension doesn't matter. Tile has to be a list of lines where every line takes the shape "Steve>Alice>10" (Steve sends 10 to Alice).

## To do

- Add different file extensions
- Find a way to handle enormous data files.

## Contact

If you have any suggestions please do not hesitate to send me a mail: <maxhalford25@gmail.com> or to open an issue on GitHub.

