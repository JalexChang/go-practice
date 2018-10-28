NUM = 101
OUTPUT_FILE = out
CODE_PATH := $(shell echo $(PWD)/**/$(NUM)-** )

all:
	@ echo "Please use 'make run NUM=<number>'" ;

run: $(CODE_PATH)
	@ echo "Find the code file in $(CODE_PATH)" ;
	@ go build -o $(OUTPUT_FILE) $(CODE_PATH) ;
	@ echo "Execution Result:\n---------\n" ;  ./$(OUTPUT_FILE) ; echo "\n---------\n" ; 
	@ rm $(OUTPUT_FILE) ;
