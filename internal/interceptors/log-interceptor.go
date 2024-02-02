package interceptors

import (
	"context"
	"strings"

	"github.com/liperm/ff-manager-server/pkg/logger"
	"google.golang.org/grpc"
)

func LoggerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	methodName := getMethodName(info.FullMethod)

	logger.Logger.RequestToMethod(methodName, req)

	response, err := handler(ctx, req)
	if err != nil {
		logger.Logger.ErrorFromMethod(methodName, err)
	} else {
		logger.Logger.ResponseFromMethod(methodName, response)
	}

	return response, err
}

func getMethodName(fullMethod string) string {
	splitMethod := strings.Split(fullMethod, "/")
	methodName := splitMethod[len(splitMethod)-1]
	return methodName
}
