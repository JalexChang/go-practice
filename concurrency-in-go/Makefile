NUM = 101
OUTPUT_FILE = out
CODE_PATH := $(shell find . -name "$(NUM)-*.go" )

all:
	@ echo "Please use 'make run NUM=<number>'" ;

run:
ifdef CODE_PATH
	@ echo "Find the code file in $(CODE_PATH)" ;
	@ go build -o $(OUTPUT_FILE) $(CODE_PATH) ;
	@ echo "Execution Result:\n---------\n" ;  ./$(OUTPUT_FILE) ; echo "\n---------\n" ; 
	@ rm $(OUTPUT_FILE) ;
else
	@ echo "Can not find the code num: $(NUM)" ;
endif
