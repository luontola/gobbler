include $(GOROOT)/src/Make.$(GOARCH)

TARG=gobbler
GOFILES=\
	hello.go\

EXECUTABLE=gob
CLEANFILES=\
	main\

main: package
	$(GC) -I_obj main.go
	$(LD) -L_obj -o $@ main.$O

install: main
	cp main $(QUOTED_GOBIN)/$(EXECUTABLE)

include $(GOROOT)/src/Make.pkg

