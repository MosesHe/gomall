export ROOT_MOD = github.com/MosesHe/gomall
.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demoproto && cwgo server -I ../../idl --type RPC --module ${ROOT_MOD}/demo/demoproto --service demoproto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demothrift && cwgo server -I ../../idl --type RPC --module ${ROOT_MOD}/demo/demothrift --service demothrift --idl ../../idl/echo.thrift

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl

.PHONY: gen-user
gen-user:
	@cd app/user && cwgo server --type RPC --service user --module ${ROOT_MOD}/app/user --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto
	@cd rpc_gen && cwgo client --type RPC --service user --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/user.proto 