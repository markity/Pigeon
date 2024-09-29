kitex:
	kitex -type protobuf -module pigeon -I ./idl idl/service/im-gateway.proto
	kitex -type protobuf -module pigeon -I ./idl idl/service/im-auth-route.proto
	kitex -type protobuf -module pigeon -I ./idl idl/service/im-relay.proto