run-local:
	grnc-yaml-bind && go build main.go && ./main -c config
	