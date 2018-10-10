package prottp

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/tomasen/realip"

	"github.com/golang/protobuf/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/mssola/user_agent"
	"github.com/theplant/appkit/kerrs"
	"github.com/theplant/appkit/server"
	"google.golang.org/grpc"
)

type ContextKey int

const (
	ContextIP     ContextKey = iota
	ContextDevice ContextKey = iota
)

type Service interface {
	Description() grpc.ServiceDesc
}

type HTTPStatusError interface {
	HTTPStatusCode() int
}

type ErrorResponse interface {
	Message() proto.Message
}

type ErrorWithStatus interface {
	HTTPStatusError
	ErrorResponse
	error
}

type respError struct {
	statusCode int
	body       proto.Message
}

func (re *respError) Message() proto.Message {
	return re.body
}

func (re *respError) HTTPStatusCode() int {
	if re.statusCode == 0 {
		return http.StatusUnprocessableEntity
	}
	return re.statusCode
}

func (re *respError) Error() string {
	return "prottp error"
}

func NewError(statusCode int, body proto.Message) ErrorWithStatus {
	return &respError{statusCode: statusCode, body: body}
}

func Handle(mux *http.ServeMux, service Service, interceptor grpc.UnaryServerInterceptor, mws ...server.Middleware) {
	sn := service.Description().ServiceName
	fmt.Println("/" + sn)
	hd := Wrap(service, interceptor)
	if len(mws) > 0 {
		hd = server.Compose(mws...)(hd)
	}
	mux.Handle("/"+sn+"/", http.StripPrefix("/"+sn, hd))
}

func Wrap(service Service, interceptor grpc.UnaryServerInterceptor) http.Handler {
	mux := http.NewServeMux()

	d := service.Description()

	for _, desc := range d.Methods {
		fmt.Println("/" + d.ServiceName + "/" + desc.MethodName)
		mux.Handle("/"+desc.MethodName, wrapMethod(service, desc, interceptor))
	}

	return mux
}

var marshaler = jsonpb.Marshaler{
	EnumsAsInts:  false,
	EmitDefaults: true,
	Indent:       "\t",
	OrigName:     true,
}

var unmarshaler = jsonpb.Unmarshaler{
	AllowUnknownFields: false,
}

const jsonContentType = "application/json"
const xprottpContentType = "application/x.prottp"

func isMimeTypeJSON(contentType string) bool {
	return strings.Index(strings.ToLower(contentType), jsonContentType) >= 0
}

func isContentTypeJSON(r *http.Request) bool {
	return isMimeTypeJSON(r.Header.Get("Content-Type"))
}

func shouldReturnJSON(r *http.Request) bool {
	acceptString := strings.ToLower(r.Header.Get("Accept"))
	if len(acceptString) == 0 {
		return isContentTypeJSON(r)
	}
	jsonIndex := strings.Index(acceptString, jsonContentType)
	xprottpIndex := strings.Index(acceptString, xprottpContentType)

	if jsonIndex < 0 && xprottpIndex < 0 {
		return isContentTypeJSON(r)
	}

	if jsonIndex < 0 {
		jsonIndex = 9999
	}
	if xprottpIndex < 0 {
		xprottpIndex = 10000
	}

	if jsonIndex < xprottpIndex {
		return true
	}

	return false
}

func wrapMethod(service interface{}, m grpc.MethodDesc, interceptor grpc.UnaryServerInterceptor) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dec := func(i interface{}) (err error) {
			unmarshaler.Unmarshal(r.Body, i.(proto.Message))
			return
		}

		newInterceptor := grpc_middleware.ChainUnaryServer(
			prottpInterceptor(r),
			interceptor,
		)
		resp, err := m.Handler(service, r.Context(), dec, newInterceptor)

		statusCode := 0
		if err != nil {
			handled := false
			if statusErr, ok := err.(HTTPStatusError); ok {
				handled = true
				statusCode = statusErr.HTTPStatusCode()
			}
			if msgErr, ok := err.(ErrorResponse); ok {
				WriteMessage(statusCode, msgErr.Message(), w, r)
				handled = true
			}

			if !handled {
				panic(err)
			}
			return
		}
		WriteMessage(statusCode, resp.(proto.Message), w, r)
	})
}

// WriteMessage is exported to be used for middleware to return proto message
func WriteMessage(statusCode int, msg proto.Message, w http.ResponseWriter, r *http.Request) {
	var err error
	var isJSON = isMimeTypeJSON(w.Header().Get("Content-Type"))
	if w.Header().Get("Content-Type") == "" {
		isJSON = shouldReturnJSON(r)
		contentType := xprottpContentType
		if isJSON {
			contentType = jsonContentType
		}
		w.Header().Set("Content-Type", fmt.Sprintf("%s;type=%s", contentType, proto.MessageName(msg)))
	}

	// start write body
	var b []byte

	if isJSON {
		buf := bytes.NewBuffer(nil)
		err = marshaler.Marshal(buf, msg)
		if err != nil {
			panic(kerrs.Wrapv(err, "marshal message to json error", "response", msg))
		}
		b = buf.Bytes()
	} else {
		b, err = proto.Marshal(msg)
		if err != nil {
			panic(kerrs.Wrapv(err, "marshal message to proto error", "response", msg))
		}
	}

	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	w.WriteHeader(statusCode)
	w.Write(b)
}

func prottpInterceptor(request *http.Request) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ip := realip.FromRequest(request)
		ua := user_agent.New(request.Header.Get("User-Agent"))
		newContext := context.WithValue(ctx, ContextIP, ip)
		name, _ := ua.Browser()
		newContext = context.WithValue(newContext, ContextDevice, name)
		return handler(newContext, req)
	}
}
