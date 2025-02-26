.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demoproto && cwgo server -I ../../idl --type RPC --module github.com/MosesHe/gomall/demo/demoproto --service demoproto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demothrift && cwgo server -I ../../idl --type RPC --module github.com/MosesHe/gomall/demo/demothrift --service demothrift --idl ../../idl/echo.thrift

.PHONY: gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth.proto --service frontend -module github.com/MosesHe/gomall/app/frontend -I ../../idl
