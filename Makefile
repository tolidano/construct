.PHONY: list
.PHONY: test
.PHONY: cov
.PHONY: install

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs

test:
	go test

cov:
	go test --coverprofile /tmp/cov.out
	go tool cover --cover /tmp/cov.out --html /tmp/cov.html
	google-chrome /tmp/cov.html

install:
	go install ./construct
