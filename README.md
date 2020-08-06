## Ride fare estimator exercise

#### Tech notes
* Hexagonal architecture:
    * The root of the internal package is the domain layer. Ride is the aggregate root here
    * Application layer is formed by the packages refering to some action or service, like `creator` and `reader`
    * Any other package refers to the infrastructure implementation of the domain, like the `storage` one
    * IO is outside the internal package
* Avoiding an event dispatcher as this is useless for the purpose of the test
* Dependency injection with [google wire](https://github.com/google/wire)
* [Command bus](https://github.com/chiguirez/cromberbus). [Chiguirez](https://github.com/chiguirez) is a github organization that me and a few friends have, where we make a few Go snippets for developing microservices.
* [Snout](https://github.com/chiguirez/snout)
    
#### Prerequisites
* [GOLang](https://golang.org/) v1.14
* [Make](https://ftp.gnu.org/old-gnu/Manuals/make-3.79.1/html_chapter/make_2.html)

the following shows the commands

    > make help

    usage: make <command>
    
    commands:
        install              - get the modules
        test                 - run all tests
        run                  - execute
        
to install the modules

    > make install

#### Running

to run the program the input and output filepath need to be defined.
do this by defining both `CSV_INPUT_FILEPATH` and `CSV_OUTPUT_FILEPATH` environment variables or having a 
`fare-estimator.yaml` file in the root of the project. A `fare-estimator.yaml.dist` example file is provided 

then let

    > make run

#### Testing

to execute the tests

    > make test

#### Author

* [Jose Ortiz](https://github.com/hosseio)