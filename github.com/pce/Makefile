gen: 
	# protoc -I=proto/googleapis -I=proto/ -I=proto/pce/mixer --go-grpc_out=src/proto/pcemixer --go_out=src/proto/pcemixer proto/pce/mixer/pcemixer.proto
	protoc -I=proto/googleapis -I=proto/ -I=proto/pce/service --go-grpc_out=require_unimplemented_servers=false=src/proto/pceadder --go_out=src/proto/pceadder proto/pce/service/pceadder.proto 

# clean:
# 	rm pb/*.go


