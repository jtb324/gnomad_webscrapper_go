help:
	@echo "---------------------------HELP------------------------"
	@echo "To test the project type make unittest"
	@echo "To clean files from a previous run type make clean"
	@echo "To install all necessary dependencies type make install"
	@echo "-------------------------------------------------------"

build_project:
	@echo "building project"
	@cd src && go build -o ../build/webscrapper

unittest:
	@echo "performing all unittest"
	@ cd src && go test -v