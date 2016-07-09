# Simple HTTP Server #

A very simple file server to code static web content against instead of having to setup a more complicated webserver.

### Installation ###

Go is required: [https://golang.org/](https://golang.org/)

    go get -u bitbucket.org/andrevdwal/simplehttpserver

To build:

    go build simplehttpserver.go

### Usage ###
~~~
Usage: simplehttpserver [options...]

Options:
  -d  Root directory to serve (default "_public").
  -p  Port to listen on (default 5050)
~~~

Assuming the application is run with defaults, it will:

* Create a `_public` directory if needed
* Serve the content on `http://localhost:5050`
* `index.html` will be treated as the default page

Below is an assumed directory structure with the links next to them:
~~~
_public/index.html        http://localhost:5050 or http://localhost:5050/index.html
_public/about.html        http://localhost:5050/about.html
_public/blog/index.html   http://localhost:5050/blog or http://localhost:5050/blog/index.html
~~~