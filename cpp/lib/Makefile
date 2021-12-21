CXX = g++
CXXFLAGS = -std=c++11 -c -Wall

ifeq ($(shell uname), Darwin)
	LDLIBS += 
endif

all: build clean

build:
	rm -f libEngine.a
	$(CXX) *.cpp $(CXXFLAGS)
	ar rvs libEngine.a *.o

clean:
	rm *.o