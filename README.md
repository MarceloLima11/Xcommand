# Xcommand - Simple Directory Manager

<img src="img/Xcommand.jpeg" align="right"
     alt="Size Limit logo by Anton Lovchikov" width="150" height="178" style="border-radius:10px">

Xcommand is a simple tool for managing directories, designed to make it easy to view and manipulate files in a local file system. With Xcommand, you can:

* List files in a **specific** directory.
* Display **details** of files, such as size and modification date.
* **Delete** files you no longer need.

## Installation
To use Xcommand, follow these steps:

1. Download the repository or clone it to your computer:

        git clone https://github.com/MarceloLima11/Xcommand.git



2. Navigate to the project folder:

        cd Xcommand

3. Run the program:

        go run xcommand.go

## How to Use
### Listing Files in a Directory
To list files in a specific directory, run the following command:

    go run xcommand.go af /path/to/directory

Replace /path/to/directory with the absolute path of the directory you wish to list or search in your current path.


### Displaying File Details
To display details of a specific file, such as size and modification date, run the following command:

    go run xcommand.go dt /path/to/file

Replace /path/to/file with the absolute path of the file for which you want to obtain details or search in your current path.

### Deleting Files
To delete a file, use the following command:

    go run xcommand.go del /path/to/file

Replace /path/to/file with the absolute path of the file you want to delete or too search in current path.


### MarceloLima11: ty for read :)
