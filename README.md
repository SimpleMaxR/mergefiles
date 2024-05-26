# 📁 MergeFiles  

## **Features**  
Merge Files: Combine the contents of multiple files into one.  
File Extension Filter: Specify the type of files to merge based on their extension.  
Recursive Option: Optionally include files from subdirectories.  
Sorted Output: Files are merged in alphabetical order.  

## **Usage** 
-input: (Required) The directory containing the files to be merged.  
-output: The name of the output file. Default is output.txt.  
-ext: The file extension to filter by. Default is .txt.  
-recursive: Include files from subdirectories. Default is false.  

## **Example**  
go run main.go -input ./myfiles -output merged.txt -ext .log -recursive  
This command will merge all .log files in the ./myfiles directory and its subdirectories into a single file named merged.txt.  

Happy merging! 🚀✨  

If you have any questions or need further assistance, don't hesitate to reach out. 😊
