help:
	@echo "------------------------------------HELP-----------------------------------"
	@echo "To test the project type make unittest"
	@echo "To remove previous build run make build"
	@echo "To remove the build files from a previous run type make remove-prior-build"
	@echo "---------------------------------------------------------------------------"

remove-prior-build:
	@echo "removing the executable from the previous build"
	@rm -r bin

build-project:
	@echo "building project"
	@cd src/ && go build -o ../bin/webscrapper 

unittest:
	@echo "performing all unittest"
	@cd src/ && go test -v