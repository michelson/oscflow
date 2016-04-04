
set -e

cd "`dirname '$0'`"
SCRIPTPATH="`pwd`"
cd - > /dev/null

export GOPATH=$SCRIPTPATH
export GOBIN=

deps() {
  echo "Fetching dependencies to $SCRIPTPATH..."
  printf "#           (01/06)\r"
    go get -u -t github.com/hypebeast/go-osc/osc
  printf "##          (02/06)\r"
    go get -u -t golang.org/x/net/websocket
  printf "###         (03/06)\r"
    go get -u -t github.com/miketheprogrammer/go-thrust
  printf "\n"
}

build() {
  go build main.go
}

run() {
  echo $GOPATH
  go build oscflow ; ./oscflow 
}

test() {
  ls ./src | grep -v "\." | sed 's/\///g' | xargs go test -cover
}

format() {
  ls ./src | grep -v "\." | sed 's/\///g' | xargs go fmt
}

$1
