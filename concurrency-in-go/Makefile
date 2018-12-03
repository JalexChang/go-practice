NUM = 101
CODE_PATH := $(shell find . -name "$(NUM)-*.go" )

all:
	@ echo "Please use 'make run NUM=<number>'" ;

run:
ifdef CODE_PATH
	@ echo "Find the code file in $(CODE_PATH)" ;
	@ go run $(CODE_PATH) ;
else
	@ echo "Can not find the code num: $(NUM)" ;
endif

test:
ifdef CODE_PATH
	@ echo "Find the code file in $(CODE_PATH)" ;
	@ go test -v $(CODE_PATH) ;
else
	@ echo "Can not find the code num: $(NUM)" ;
endif

bench:
ifdef CODE_PATH
	@ echo "Find the code file in $(CODE_PATH)" ;
	@ go test -v -bench=. $(CODE_PATH) ;
else
	@ echo "Can not find the code num: $(NUM)" ;
endif
