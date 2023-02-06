# ImgToASCII

## Build instructions

```bash
mkdir bin
go build -o bin/img2ascii
```

## How to use it

```
Convert png or jpeg to ascii
usage : img2ascii [-h] {filename} [-o|-d|-a]
-h display this text
-o {filename}       set the output file
-d {int} {int}      set the size of the ascii art
-d {int} {int}      set the size of the ascii art
-a                  output all the stage of conversion in current dir
```