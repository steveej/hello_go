.PHONY: clean bitcount.test cpu.profile

all: bitcount.test cpu.profile clean

clean:
	rm -rf cpu.profile bitcount.test

bitcount.test:
	go test -cpuprofile cpu.profile

cpu.profile: bitcount.test
	go tool pprof --list CountBits cpu.profile
