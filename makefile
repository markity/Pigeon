kitex:
	kitex -type protobuf -module pigeon -I ./ idl/service/im-gateway.proto
	kitex -type protobuf -module pigeon -I ./ idl/service/im-auth-route.proto