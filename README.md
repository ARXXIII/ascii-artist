# ASCII Artist

A Go-based ASCII art generator from images.  
Supports loading images from both local disk and URLs.

## Features

-   Supports PNG and JPEG formats
-   Enter file path or image URL via console prompt
-   Automatic ASCII art scaling by width (default: 80 characters)
-   Well-structured code split into packages

## Installation

1. Clone the repository:

    ```
    git clone https://github.com/ARXXIII/ascii-artist.git
    cd ascii-artist
    ```

2. Build the project:

    ```
    go build -o ascii-artist main.go
    ```

## Usage

Run the program:

```
ascii-artist
```

You will be prompted to enter a file path or image URL:

```
URL or path: https://golang.org/doc/gopher/frontpage.png
```

The ASCII art will be displayed in your console.

## Example

```
URL or path: ./gopher.png
@##++==...
@##++==...
...
```

## Project Structure

-   `main.go` — entry point, user interaction
-   `ascii/` — package for converting images to ASCII
-   `load/` — package for loading images from disk or URL

## Requirements

-   Go 1.18+
-   Standard library only

##

MIT
