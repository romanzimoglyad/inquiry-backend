// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api.proto

package inquiry

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InquiryClient is the client API for Inquiry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InquiryClient interface {
	// Ping
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// Create lesson
	CreateLesson(ctx context.Context, in *LessonRequest, opts ...grpc.CallOption) (*CreateLessonResponse, error)
	// Update lesson
	UpdateLesson(ctx context.Context, in *LessonRequest, opts ...grpc.CallOption) (*Empty, error)
	// List Lessons
	ListLessons(ctx context.Context, in *ListLessonsRequest, opts ...grpc.CallOption) (*ListLessonsResponse, error)
	// Get Lesson
	GetLesson(ctx context.Context, in *GetLessonRequest, opts ...grpc.CallOption) (*GetLessonResponse, error)
	// Delete Lesson
	DeleteLesson(ctx context.Context, in *GetLessonRequest, opts ...grpc.CallOption) (*Empty, error)
	// List Subjects
	ListSubjects(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListSubjectsResponse, error)
	ListDictionary(ctx context.Context, in *ListDictionaryRequest, opts ...grpc.CallOption) (*ListDictionaryResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type inquiryClient struct {
	cc grpc.ClientConnInterface
}

func NewInquiryClient(cc grpc.ClientConnInterface) InquiryClient {
	return &inquiryClient{cc}
}

func (c *inquiryClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/inquiry.Inquiry/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inquiryClient) CreateLesson(ctx context.Context, in *LessonRequest, opts ...grpc.CallOption) (*CreateLessonResponse, error) {
	out := new(CreateLessonResponse)
	err := c.cc.Invoke(ctx, "/inquiry.Inquiry/CreateLesson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inquiryClient) UpdateLesson(ctx context.Context, in *LessonRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/inquiry.Inquiry/UpdateLesson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inquiryClient) ListLessons(ctx context.Context, in *ListLessonsRequest, opts ...grpc.CallOption) (*ListLessonsResponse, error) {
	out := new(ListLessonsResponse)
	err := c.cc.Invoke(ctx, "/inquiry.Inquiry/ListLessons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inquiryClient) GetLesson(ctx context.Context, in *GetLessonRequest, opts ...grpc.CallOption) (*GetLessonResponse, error) {
	out := new(GetLessonResponse)
	err := c.cc.Invoke(ctx, "/inquiry.Inquiry/GetLesson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inquiryClient) DeleteLesson(ctx context.Context, in *GetLessonRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/inquiry.Inquiry/DeleteLesson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inquiryClient) ListSubjects(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListSubjectsResponse, error) {
	out := new(ListSubjectsResponse)
	err := c.cc.Invoke(ctx, "/inquiry.Inquiry/ListSubjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inquiryClient) ListDictionary(ctx context.Context, in *ListDictionaryRequest, opts ...grpc.CallOption) (*ListDictionaryResponse, error) {
	out := new(ListDictionaryResponse)
	err := c.cc.Invoke(ctx, "/inquiry.Inquiry/ListDictionary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inquiryClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/inquiry.Inquiry/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InquiryServer is the server API for Inquiry service.
// All implementations must embed UnimplementedInquiryServer
// for forward compatibility
type InquiryServer interface {
	// Ping
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// Create lesson
	CreateLesson(context.Context, *LessonRequest) (*CreateLessonResponse, error)
	// Update lesson
	UpdateLesson(context.Context, *LessonRequest) (*Empty, error)
	// List Lessons
	ListLessons(context.Context, *ListLessonsRequest) (*ListLessonsResponse, error)
	// Get Lesson
	GetLesson(context.Context, *GetLessonRequest) (*GetLessonResponse, error)
	// Delete Lesson
	DeleteLesson(context.Context, *GetLessonRequest) (*Empty, error)
	// List Subjects
	ListSubjects(context.Context, *Empty) (*ListSubjectsResponse, error)
	ListDictionary(context.Context, *ListDictionaryRequest) (*ListDictionaryResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedInquiryServer()
}

// UnimplementedInquiryServer must be embedded to have forward compatible implementations.
type UnimplementedInquiryServer struct {
}

func (UnimplementedInquiryServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedInquiryServer) CreateLesson(context.Context, *LessonRequest) (*CreateLessonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLesson not implemented")
}
func (UnimplementedInquiryServer) UpdateLesson(context.Context, *LessonRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLesson not implemented")
}
func (UnimplementedInquiryServer) ListLessons(context.Context, *ListLessonsRequest) (*ListLessonsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLessons not implemented")
}
func (UnimplementedInquiryServer) GetLesson(context.Context, *GetLessonRequest) (*GetLessonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLesson not implemented")
}
func (UnimplementedInquiryServer) DeleteLesson(context.Context, *GetLessonRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLesson not implemented")
}
func (UnimplementedInquiryServer) ListSubjects(context.Context, *Empty) (*ListSubjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubjects not implemented")
}
func (UnimplementedInquiryServer) ListDictionary(context.Context, *ListDictionaryRequest) (*ListDictionaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDictionary not implemented")
}
func (UnimplementedInquiryServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedInquiryServer) mustEmbedUnimplementedInquiryServer() {}

// UnsafeInquiryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InquiryServer will
// result in compilation errors.
type UnsafeInquiryServer interface {
	mustEmbedUnimplementedInquiryServer()
}

func RegisterInquiryServer(s grpc.ServiceRegistrar, srv InquiryServer) {
	s.RegisterService(&Inquiry_ServiceDesc, srv)
}

func _Inquiry_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InquiryServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inquiry.Inquiry/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InquiryServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inquiry_CreateLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LessonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InquiryServer).CreateLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inquiry.Inquiry/CreateLesson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InquiryServer).CreateLesson(ctx, req.(*LessonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inquiry_UpdateLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LessonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InquiryServer).UpdateLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inquiry.Inquiry/UpdateLesson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InquiryServer).UpdateLesson(ctx, req.(*LessonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inquiry_ListLessons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLessonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InquiryServer).ListLessons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inquiry.Inquiry/ListLessons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InquiryServer).ListLessons(ctx, req.(*ListLessonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inquiry_GetLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLessonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InquiryServer).GetLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inquiry.Inquiry/GetLesson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InquiryServer).GetLesson(ctx, req.(*GetLessonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inquiry_DeleteLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLessonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InquiryServer).DeleteLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inquiry.Inquiry/DeleteLesson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InquiryServer).DeleteLesson(ctx, req.(*GetLessonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inquiry_ListSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InquiryServer).ListSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inquiry.Inquiry/ListSubjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InquiryServer).ListSubjects(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inquiry_ListDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InquiryServer).ListDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inquiry.Inquiry/ListDictionary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InquiryServer).ListDictionary(ctx, req.(*ListDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inquiry_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InquiryServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inquiry.Inquiry/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InquiryServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Inquiry_ServiceDesc is the grpc.ServiceDesc for Inquiry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inquiry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inquiry.Inquiry",
	HandlerType: (*InquiryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Inquiry_Ping_Handler,
		},
		{
			MethodName: "CreateLesson",
			Handler:    _Inquiry_CreateLesson_Handler,
		},
		{
			MethodName: "UpdateLesson",
			Handler:    _Inquiry_UpdateLesson_Handler,
		},
		{
			MethodName: "ListLessons",
			Handler:    _Inquiry_ListLessons_Handler,
		},
		{
			MethodName: "GetLesson",
			Handler:    _Inquiry_GetLesson_Handler,
		},
		{
			MethodName: "DeleteLesson",
			Handler:    _Inquiry_DeleteLesson_Handler,
		},
		{
			MethodName: "ListSubjects",
			Handler:    _Inquiry_ListSubjects_Handler,
		},
		{
			MethodName: "ListDictionary",
			Handler:    _Inquiry_ListDictionary_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Inquiry_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
