build:
		protoc -I. --go_out=plugins=micro:. \
	proto/vessel/vessel.proto
		docker build -t ryanyogan/shippy-vessel-service .
		docker push ryanyogan/shippy-vessel-service

run:
		docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 \
		ryanyogan/shippy-vessel-service