@echo off
start "api1" go run server_main.go --server_address :8001 &
start "api2" go run server_main.go --server_address :8002 &
start "api3" go run server_main.go --server_address :8003
pause