# Stock Feed build
# GoLang

# Author:
# Nick Pankow
# Dec 01, 2014

# Project Paths
#GOPATH = $(GOPATH)
SOURCE = github.com
USER = nickpankow

# Project Files
TARGET = stockfeed
BUILDFILES =  feed.go
BUILDEXT = .exe

# Go Paths
INSTALLPATH = $(GOPATH)\bin\
OUTPUT = $(INSTALLPATH)\$(TARGET)$(BUILDEXT)

# Build
# $(OUTPUT) : $(BUILDFILES)
all:
	go install $(SOURCE)\$(USER)\$(TARGET)
