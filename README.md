# manage-pictures
A simple program written in Golang that allows you to rename all pictures (.jpg/.jpeg) using the EXIF data (datetime,placename (API OpenStreetSource), gps position) if available.

## Usage
```
# Install package
go get -u github.com/xor-gate/goexif2/exif
go get -u github.com/xor-gate/goexif2/tiff

# Build program
go build rename-pictures.go

# Run program
./rename-pictures <directory_path>
```