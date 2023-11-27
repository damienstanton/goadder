default: build go

build:
	@cargo build

go:
	@uniffi-bindgen-go src/math.udl