
#INSTALL AND RUN APP 

### setup golang

  install golang

  export gopath, example

  `export GOPATH=~/go`

  `go get github.com/hypebeast/go-osc/osc`

  `go run main.go`

  `go build main.go`

### Run application

tab1: `cd interface ; middleman build`

tab2: `npm start` // `make run` (todo: ws to interface)


### Build static assets

`webpack`

#MODEL

![model](./model.png)

# COMPILATION ARM

https://github.com/adafruit/ARM-toolchain-vagrant

vagrant up

vagrant ssh

  + http://dave.cheney.net/2012/09/25/installing-go-on-the-raspberry-pi

  + https://github.com/sourcegraph/go-webkit2

  + puredta
    https://puredata.info/docs/developer/BuildingPdExtendedForRaspberryPiRaspbianWheezyArmhf

  - sudo apt-get install libwebkit2gtk-3.0-dev

  - osx
  - instalar webkit1 (agregar a /usr/local/formulas/webkit1.rb)
  - https://github.com/pedromartins/homebrew-webkit/blob/master/webkitgtk.rb

#TODO

+ send messages from interface to pd
+ model message schema
+ build with "electron-rebuild": "^0.2.2",
+ crosscompile

### refs: 

https://medium.com/@Agro/developing-desktop-applications-with-electron-and-react-40d117d97564#.hbtyqn71a