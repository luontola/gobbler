
Gobbler
======

Gobbler is a build tool for the [Go programming language](http://golang.org/). It allows building and testing multi-package Go projects with one command, without need for manually written Makefiles.

Source code is available at <http://github.com/orfjackal/gobbler>

For discussion, use the [golang-nuts mailing list](http://groups.google.com/group/golang-nuts), at least until Gobbler has so many users that it requires its own mailing list. You may also contact Gobbler's developer, [Esko Luontola](http://github.com/orfjackal), by email.


Project Goals
-------------

- **Simple to use** - Build the project and run its tests with one command-line command.

- **Convention over configuration** - When files are put into source and test directories (/src/main/go and /src/test/go), then all packages and tests are detected automatically. The project configuration file should only need to say whether the project is a library or an executable, and what is the name of the executable.

- **Smart dependency handling** - Allow packages (including test code) to depend on other packages in the same project, without requiring the build process to "make install" the dependencies to $GOROOT. Build only those packages which have changed since last build (or whose dependencies have changed). In the long-term, also dependencies to external libraries will be resolved automatically, similar to Maven et al.


License
-------

Copyright © 2010 Esko Luontola <<http://www.orfjackal.net>>  
This software is released under the Apache License 2.0.  
The license text is at <http://www.apache.org/licenses/LICENSE-2.0>
