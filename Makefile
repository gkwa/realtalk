run: realtalk
	./realtalk

realtalk: ./dist/realtalk_darwin_amd64_v1/realtalk
	cp $< $@

./dist/realtalk_darwin_amd64_v1/realtalk: main.go
	gofumpt -w $<
	goreleaser build --single-target --snapshot --clean

clean:
	rm -f realtalk
	rm -rf dist
