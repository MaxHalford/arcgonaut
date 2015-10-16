# arcgonaut

Software for creating arc diagrams such as:

![Example](example.png)

arcgonaut is written in Go. It uses [Laurent Le Goff](https://github.com/llgcode)'s [draw2d](https://github.com/llgcode/draw2d) for the imaging tools and [Lucas Beyer](https://github.com/lucasb-eyer)'s [go-colorful](https://github.com/lucasb-eyer/go-colorful) for generating colors. The rest is a lot geometry/rescaling.

## Installation

### For users

[Download the binary](http://maxhalford.com/data/arcgonaut)

### For developpers

```sh
go get https://github.com/MaxHalford/arcgonaut
```

## Usage

For the moment the tool is usable in the command-line.

- Naviguate towards the directory where the binary is.
- ``./arcgonaut -f="example.arcgo"

A PNG will be generated into the same directory. The program takes as input a file that has a [specific format](example.arcgo). The file extension doesn't matter. Tile has to be a list of lines where every line takes the shape "Steve>Alice>10" (Steve sends 10 to Alice).

## Improvements

If you have any suggestions please do not hesitate to send me a mail: <maxhalford25@gmail.com> or to open an issue on GitHub. My goal is to make this tool robust and scalable, if you encounter any bugs please tell me.

## License

See the [license file](LICENSE).