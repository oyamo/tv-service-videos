 protoc -I. --go_out=plugins=grpc:. proto/videos/videos.proto
 mv github.com/oyamoh-brian/tv-service-videos/videos.pb.go proto/videos
 rm -rf github.com