kitex:
	kitex -type protobuf -module pigeon -I ./idl idl/service/im-gateway.proto
	kitex -type protobuf -module pigeon -I ./idl idl/service/im-auth-route.proto
	kitex -type protobuf -module pigeon -I ./idl idl/service/im-relay.proto
	kitex -type protobuf -module pigeon -I ./idl idl/service/im-relation.proto
	kitex -type protobuf -module pigeon -I ./idl idl/service/im-chat-evloop.proto

remake:
	etcdctl del "" --prefix
	etcdctl put /chatevloop_config/node_info/1 '{"version":1,"nodes":[{"node_name":"node1","ipport":"127.0.0.1:8001"}]}'
	etcdctl put /chatevloop_config/version 1

	redis-cli flushdb

	mysql -uroot -e "drop database if exists im_relation"
	mysql -uroot -e "drop database if exists im_chatevloop"
	mysql -uroot -e "drop database if exists im_auth_route"
	mysql -uroot -e "create database im_relation"
	mysql -uroot -e "create database im_chatevloop"
	mysql -uroot -e "create database im_auth_route"

	./cmd/init_mysql_user/init_mysql_user

	