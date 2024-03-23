templateFile=$GOPATH/src/github.com/go-easygen/easygen/test/commandlineFlag
[ -s $templateFile.tmpl ] || templateFile=/usr/share/gocode/src/github.com/go-easygen/easygen/test/commandlineFlag
[ -s $templateFile.tmpl ] || templateFile=/usr/share/doc/easygen/examples/commandlineFlag
[ -s $templateFile.tmpl ] || templateFile=/usr/local/share/easygen/examples/commandlineFlag
[ -s $templateFile.tmpl ] || {
  echo No template file found
  exit 1
}

easygen $templateFile diskcheck_cli | gofmt > diskcheck_cliDef.go
