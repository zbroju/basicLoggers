# basicLoggers
Basic loggers for commandline applications.

# Examples

    package main

    import (
        "github.com/zbroju/basicLoggers"
        "os"
    )
    
    func main() {
        number := len(os.Args)
        printMessage, errorMessage := basicLoggers.GetLoggers("exampleApp")
        
        // If no arguments given exit gently with the error message
        if number == 1 {
            errorMessage.Fatalln("At least one argument needed.")
        }
    
        printMessage.Printf("Performing action on arguments")
    }


Gives the following results:

    $ example argument
    exampleApp: Performing action on arguments

    $ example
    exampleApp: At least one argument needed.

# License
GNU General Public License

# Author
Marcin 'Zbroju' Zbroiński