# Image Resizer and Maskable Icon Generator

## Overview
A command-line tool for resizing images and creating maskable icons.

## Features
- Resize images to specified dimensions
- Create maskable icons for web and mobile apps
- Supports PNG and JPEG formats

## Installation
```bash
git clone https://github.com/yourusername/image-resizer.git
cd image-resizer
go mod tidy
```

## Usage

### Resize Image
```bash
go run cmd/image-resizer/main.go resize -input input.jpg -output output.jpg -width 300 -height 200
```

### Create Maskable Icon
```bash
go run cmd/image-resizer/main.go maskable -input original_icon.png -output maskable_icon.png [-size 512]
```

## Build
```bash
go build -o image-resizer cmd/image-resizer/main.go
```

## License
[Your License Here]