protoc --go_out=. calculator/calculatorpb/calculator.proto
protoc --go-grpc_out=. calculator/calculatorpb/calculator.proto
protoc --go_out=. greet/greetpb/greet.proto
protoc --go-grpc_out=. greet/greetpb/greet.proto