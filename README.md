# go-server
A simple HTTP Server written in Go Language

###Setup Go (for Ubuntu users):
1. Download the archive from https://golang.org/dl/
2. Extract it to /usr/local/ <br>
```tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz```
3. Add /usr/local/go/bin to the PATH environment variable by adding 
this to the last line of your /etc/profile or $HOME/.profile:<br>
```export PATH=$PATH:/usr/local/go/bin```
4. Create a directory to contain your workspace, 
$HOME/golang , and set the GOPATH environment variable to point to that location in similar manner as point 3. <br>
```export GOPATH=$HOME/golang```

###Get the source code:
Simply do a ```go get github.com/aashaka/go-server``` in the directory $HOME/golang you just made.

###Run the server:
Simply do a ```go install``` and then ```$GOPATH/bin/go-server``` in 
the ```$HOME/golang/src/github.com/aashaka/go-server``` directory.

The server runs at port 1200.

###TODOs: 
1. Add code to serve different pages. Currently the server takes only GET methods and serves only index.html.
