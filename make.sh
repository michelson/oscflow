
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
  printf "####        (04/06)\r"
    go get github.com/bmizerany/pat
  printf "\n"
}

dev() {
  cd interface ; middleman s
}

build() {
  go build oscflow
  #cd interface ; middleman build
}

run() {
  echo $GOPATH
  go run src/oscflow/main.go 
}

test() {
  ls ./src | grep -v "\." | sed 's/\///g' | xargs go test -v #-cover
}

format() {
  ls ./src | grep -v "\." | sed 's/\///g' | xargs go fmt
}

$1
