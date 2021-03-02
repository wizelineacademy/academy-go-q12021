##
# Wizeline's Go Workshop
#
# @file
# @version 0.1

PID      = /tmp/awesome-golang-project.pid
GO_FILES = $(wildcard *.go)
APP      = "./app"

serve: restart
	@fswatch -or . --event=Updated | xargs -n1 -I{} make restart || make kill

changed:
	@echo "Changed $?"

kill:
	@kill `cat $(PID)` || true

before:
	@echo "actually do nothing"

$(APP): $(GO_FILES)
	@echo "$(APP) phase"
	@echo $@
	@go build -o $@ $?

restart: kill before $(APP)
	@echo "Changed $?"
	@$(APP) & echo $$! > $(PID)

.PHONY: serve restart kill before # let's go to reserve rules names


# end
