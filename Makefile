EXERCISMS := $(shell find . -name ".exercism")
EXERCISES := $(EXERCISMS:./%/.exercism=%)
MAKEFILES := $(foreach ex,${EXERCISES},${ex}/Makefile)


init: ${MAKEFILES}
.PHONY: init

%/Makefile:
	@echo "EXERCISE := ${@D}" > $@
	@echo "" >> $@
	@echo "include ../Makefile" >> $@


# Exercise Commands

test:
	go test
.PHONY: test

dev:
	find . -name "*.go" | entr -cdr bash -c "timeout 3 go test"
.PHONY: dev
