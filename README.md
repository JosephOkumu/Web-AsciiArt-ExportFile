# Web-AsciiArt-ExportFile

## Overview
This program is a Golang web application that allows users to generate and export/download ASCII art. This application provides functionality to export ASCII art in a chosen format via HTTP. It includes appropriate headers for file transfer. 
## Documentation
This section illustrates how to make use of this program.

### Installation
To run this program, download and install the latest version of Go from [here](https://go.dev/doc/install)

### Usage
1. Clone this repository on to the terminal by using the following command:
```bash
git clone https://github.com/JosephOkumu/Web-AsciiArt-ExportFile.git
```
2. Navigate into the ascii-art-web directory by using the command:
```bash
cd web-asciiart-exportfile
```
3. Run the program by typing this command on the terminal:
```bash
go run .
```
4. Navigate to localhost port 8080 to view the web version of our project by entering this url in your browser:
```bash
http://localhost/8080
```
A form will be displayed, and on the form, enter your input text in the text field, choose banner format and click "Convert to Ascii Art".
Example:
When you type "hello" and select "standard" as the banner format, then upon clicking "Convert to ASCII Art," the following output will be displayed on a new page.

```bash
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
```
5. Download the output by choosing your preferred file format, then click "Download".

## Features
- This program can write the ascii-art graphical representation to a new page.
- This program allows you to select a banner file format to display the ascii art.
- The web pages of this program are styled and are responsive.
- This program has our own extra customized banner file. 
- This program provides csv, txt and html file download options.

## Implementation details: algorithm
### Endpoints:
1. Root ("/)
2. Ascii-art ("/ascii-art)

### Algorithm:
1. Set up HTTP server using Go's 'net/http' package.
2. Define route handlers for both endpoints:
3. In the root handler serve the main page with a form for input text.
4. In the ascii-art handler:
    - Parse data from the request.
    - Read the appropriate ASCII art file based on user input
    - Process file contents to generate ASCII art.
    - Generate HTML response with the ASCII art.
    - Write HTML to response writer.

### Data Flow:
User Input -> Server -> File Read -> ASCII Art Generation -> HTML Response -> Web Browser


## Contributions
Pull requests are welcome where users of this program are allowed to contribute to this project in terms of adding features, or fixing bugs.

For major changes, please open an issue first to discuss what you would like to change.
## Authors
[JosephOKumu](https://github.com/JosephOkumu)


## Licence
[MIT](https://choosealicense.com/licenses/mit/)
## Credits
[Zone01Kisumu](https://www.zone01kisumu.ke/)
