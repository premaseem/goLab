# Go lang important notes 

## Go installation on Mac

Insall library 
$ brew update && brew install golang

Setup work space - It’s considered best practice to use $HOME/go location for your workspace, so let’s do that! 
```
mkdir -p $HOME/go/{bin,src,pkg}
```

Setup Environment variabe in .zshrc or .bash
```
export GOPATH=$HOME/go
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
```

## Is go platform independent? 
you go cross-platform same as you would with C, compile to your target OS and CPU, and it’s largely up to you to make things work.
Mostly it’ll just be file paths, i.e. like C:\blah on Windows, rather than /blah on Linux.
By setting GOOS and GOARCH environment variables the compiler will produce executable for the destination platform.

