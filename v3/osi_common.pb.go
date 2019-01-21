// Code generated by protoc-gen-go. DO NOT EDIT.
// source: osi_common.proto

package osi3

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//
// \brief A cartesian 3D vector for positions, velocities or accelerations or
// its uncertainties.
//
// The coordinate system is defined as right-handed.
//
// Units are [m] for positions, [m/s] for velocities and [m/s^2] for
// accelerations.
//
type Vector3D struct {
	// The x coordinate.
	//
	// Unit: [m] [m/s] or [m/s^2]
	//
	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	// The y coordinate.
	//
	// Unit: [m] [m/s] or [m/s^2]
	//
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	// The z coordinate.
	//
	// Unit: [m] [m/s] or [m/s^2]
	//
	Z                    float64  `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector3D) Reset()         { *m = Vector3D{} }
func (m *Vector3D) String() string { return proto.CompactTextString(m) }
func (*Vector3D) ProtoMessage()    {}
func (*Vector3D) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb1efe539488b99, []int{0}
}

func (m *Vector3D) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector3D.Unmarshal(m, b)
}
func (m *Vector3D) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector3D.Marshal(b, m, deterministic)
}
func (m *Vector3D) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector3D.Merge(m, src)
}
func (m *Vector3D) XXX_Size() int {
	return xxx_messageInfo_Vector3D.Size(m)
}
func (m *Vector3D) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector3D.DiscardUnknown(m)
}

var xxx_messageInfo_Vector3D proto.InternalMessageInfo

func (m *Vector3D) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector3D) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Vector3D) GetZ() float64 {
	if m != nil {
		return m.Z
	}
	return 0
}

//
// \brief A cartesian 2D vector for positions, velocities or accelerations or
// its uncertainties.
//
// Units are [m] for positions, [m/s] for velocities and [m/s^2] for
// accelerations.
//
type Vector2D struct {
	// The x coordinate.
	//
	// Unit: [m] [m/s] or [m/s^2]
	//
	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	// The y coordinate.
	//
	// Unit: [m] [m/s] or [m/s^2]
	//
	Y                    float64  `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector2D) Reset()         { *m = Vector2D{} }
func (m *Vector2D) String() string { return proto.CompactTextString(m) }
func (*Vector2D) ProtoMessage()    {}
func (*Vector2D) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb1efe539488b99, []int{1}
}

func (m *Vector2D) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector2D.Unmarshal(m, b)
}
func (m *Vector2D) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector2D.Marshal(b, m, deterministic)
}
func (m *Vector2D) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector2D.Merge(m, src)
}
func (m *Vector2D) XXX_Size() int {
	return xxx_messageInfo_Vector2D.Size(m)
}
func (m *Vector2D) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector2D.DiscardUnknown(m)
}

var xxx_messageInfo_Vector2D proto.InternalMessageInfo

func (m *Vector2D) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector2D) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

//
// \brief A timestamp.
//
// Names and types of fields are chosen in accordance to
// google/protobuf/timestamp.proto to allow a possible switch in the future.
// There is no definition of the zero point in time neither it is the Unix
// epoch. A simulation may start at the zero point in time but it is not
// mandatory.
//
type Timestamp struct {
	// The number of seconds since the start of e.g. the simulation / system /
	// vehicle.
	//
	// Unit: [s]
	//
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// The number of nanoseconds since the start of the last second.
	//
	// Range: [0, 999.999.999]
	//
	// Unit: [ns]
	//
	Nanos                uint32   `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb1efe539488b99, []int{2}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() uint32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

//
// \brief The dimension of a 3D box, e.g. the size of a 3D bounding box or its
// uncertainties.
//
// \image html OSI_Dimension3d.jpg
//
// The dimensions are positive. Uncertainties are negative or positive.
//
// Dimension is defined in the specified reference coordinate frame along the
// x-axis (=length), y-axis (=width) and z-axis (=height).
//
type Dimension3D struct {
	// The length of the box.
	//
	// Unit: [m]
	//
	Length float64 `protobuf:"fixed64,1,opt,name=length,proto3" json:"length,omitempty"`
	// The width of the box.
	//
	// Unit: [m]
	//
	Width float64 `protobuf:"fixed64,2,opt,name=width,proto3" json:"width,omitempty"`
	// The height of the box.
	//
	// Unit: [m]
	//
	Height               float64  `protobuf:"fixed64,3,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dimension3D) Reset()         { *m = Dimension3D{} }
func (m *Dimension3D) String() string { return proto.CompactTextString(m) }
func (*Dimension3D) ProtoMessage()    {}
func (*Dimension3D) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb1efe539488b99, []int{3}
}

func (m *Dimension3D) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dimension3D.Unmarshal(m, b)
}
func (m *Dimension3D) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dimension3D.Marshal(b, m, deterministic)
}
func (m *Dimension3D) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dimension3D.Merge(m, src)
}
func (m *Dimension3D) XXX_Size() int {
	return xxx_messageInfo_Dimension3D.Size(m)
}
func (m *Dimension3D) XXX_DiscardUnknown() {
	xxx_messageInfo_Dimension3D.DiscardUnknown(m)
}

var xxx_messageInfo_Dimension3D proto.InternalMessageInfo

func (m *Dimension3D) GetLength() float64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Dimension3D) GetWidth() float64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Dimension3D) GetHeight() float64 {
	if m != nil {
		return m.Height
	}
	return 0
}

//
// \brief A 3D orientation, orientation rate or orientation acceleration (i.e.
// derivatives) or its uncertainties denoted in euler angles.
//
// Units are [rad] for orientation [rad/s] for rates and [rad/s^2] for
// accelerations
//
// The preferred angular range is (-pi, pi]. The coordinate system is defined as
// right-handed.
// For the sense of each rotation, the right-hand rule applies.
//
// The rotations are to be performed \b yaw \b first (around the z-axis),
// \b pitch \b second (around the new y-axis) and \b roll \b third (around the
// new x-axis) to follow the definition according to [1] (Tait-Bryan / Euler
// convention z-y'-x'').
//
// Roll/Pitch are 0 if the objects xy-plane is parallel to its parent's
// xy-plane. Yaw is 0 if the object's local x-axis is parallel to its parent's
// x-axis.
//
// <tt>Rotation_yaw_pitch_roll = Rotation_roll*Rotation_pitch*Rotation_yaw</tt>
//
// <tt>vector_global_coord_system :=
// Inverse_Rotation_yaw_pitch_roll(</tt><tt>Orientation3d</tt><tt>)*(vector_local_coord_system)
// + local_origin::position</tt>
//
// \attention This definition changed in OSI version 3.0.0. Previous OSI
// versions  (V2.xx) had an other definition.
//
// \par References:
// - [1] DIN ISO 8855:2013-11
//
type Orientation3D struct {
	// The roll angle/rate/acceleration.
	//
	// Unit: [rad] [rad/s] or [rad/s^2]
	//
	Roll float64 `protobuf:"fixed64,1,opt,name=roll,proto3" json:"roll,omitempty"`
	// The pitch angle/rate/acceleration.
	//
	// Unit: [rad] [rad/s] or [rad/s^2]
	//
	Pitch float64 `protobuf:"fixed64,2,opt,name=pitch,proto3" json:"pitch,omitempty"`
	// The yaw angle/rate/acceleration.
	//
	// Unit: [rad] [rad/s] or [rad/s^2]
	//
	Yaw                  float64  `protobuf:"fixed64,3,opt,name=yaw,proto3" json:"yaw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Orientation3D) Reset()         { *m = Orientation3D{} }
func (m *Orientation3D) String() string { return proto.CompactTextString(m) }
func (*Orientation3D) ProtoMessage()    {}
func (*Orientation3D) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb1efe539488b99, []int{4}
}

func (m *Orientation3D) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Orientation3D.Unmarshal(m, b)
}
func (m *Orientation3D) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Orientation3D.Marshal(b, m, deterministic)
}
func (m *Orientation3D) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Orientation3D.Merge(m, src)
}
func (m *Orientation3D) XXX_Size() int {
	return xxx_messageInfo_Orientation3D.Size(m)
}
func (m *Orientation3D) XXX_DiscardUnknown() {
	xxx_messageInfo_Orientation3D.DiscardUnknown(m)
}

var xxx_messageInfo_Orientation3D proto.InternalMessageInfo

func (m *Orientation3D) GetRoll() float64 {
	if m != nil {
		return m.Roll
	}
	return 0
}

func (m *Orientation3D) GetPitch() float64 {
	if m != nil {
		return m.Pitch
	}
	return 0
}

func (m *Orientation3D) GetYaw() float64 {
	if m != nil {
		return m.Yaw
	}
	return 0
}

//
// \brief A common identifier (ID), represented as an integer.
//
// Has to be unique among all simulated items at any given time. For ground
// truth, the identifier of an item (object, lane, sign, etc.) must remain
// stable over its lifetime. \c Identifier values may be only be reused if the
// available address space is exhausted and the specific values have not been in
// use for several timesteps. Sensor specific tracking IDs have no restrictions
// and should behave according to the sensor specifications.
//
// The value MAX(uint64) = 2^(64) -1 =
// 0b1111111111111111111111111111111111111111111111111111111111111111 is
// reserved and indicates an invalid ID or error.
//
type Identifier struct {
	// The identifier's value.
	//
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb1efe539488b99, []int{5}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

//
// \brief Specifies the mounting position of a sensor.
//
// Details are specified in each instance where \c MountingPosition is used.
//
type MountingPosition struct {
	// Offset position relative to the specified reference coordinate system.
	//
	Position *Vector3D `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	// Orientation offset relative to the specified reference coordinate system.
	//
	// <tt>Origin_sensor :=
	// Rotation_yaw_pitch_roll(#orientation)*(Origin_reference_coordinate_system
	// - #position)</tt>
	//
	Orientation          *Orientation3D `protobuf:"bytes,2,opt,name=orientation,proto3" json:"orientation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MountingPosition) Reset()         { *m = MountingPosition{} }
func (m *MountingPosition) String() string { return proto.CompactTextString(m) }
func (*MountingPosition) ProtoMessage()    {}
func (*MountingPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb1efe539488b99, []int{6}
}

func (m *MountingPosition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MountingPosition.Unmarshal(m, b)
}
func (m *MountingPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MountingPosition.Marshal(b, m, deterministic)
}
func (m *MountingPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MountingPosition.Merge(m, src)
}
func (m *MountingPosition) XXX_Size() int {
	return xxx_messageInfo_MountingPosition.Size(m)
}
func (m *MountingPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_MountingPosition.DiscardUnknown(m)
}

var xxx_messageInfo_MountingPosition proto.InternalMessageInfo

func (m *MountingPosition) GetPosition() *Vector3D {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *MountingPosition) GetOrientation() *Orientation3D {
	if m != nil {
		return m.Orientation
	}
	return nil
}

//
// \brief A spherical representation for a point or vector in 3D space.
//
// Used e.g., for low level representations of radar detections.
//
// Azimuth and elevation are defined as the rotations that would have to be
// applied to the local frame (e.g sensor frame definition in
// \c SensorDetectionHeader) to make its x-axis point towards the referenced
// point or to align it with the referenced vector. The rotations are to be
// performed \b azimuth \b first (around the z-axis) and \b elevation \b second
// (around the new y-axis) to follow the definition of \c Orientation3d. For the
// sense of each rotation, the right-hand rule applies.
//
// <tt>vector_cartesian :=
// Rotation(#elevation)*Rotation(#azimuth)*Unit_vector_x*#distance</tt>
//
type Spherical3D struct {
	// The radial distance.
	//
	// Unit: [m]
	//
	Distance float64 `protobuf:"fixed64,1,opt,name=distance,proto3" json:"distance,omitempty"`
	// The azimuth (horizontal) angle.
	//
	// Unit: [rad]
	//
	Azimuth float64 `protobuf:"fixed64,2,opt,name=azimuth,proto3" json:"azimuth,omitempty"`
	// The elevation (vertical) angle.
	//
	// Unit: [rad]
	//
	Elevation            float64  `protobuf:"fixed64,3,opt,name=elevation,proto3" json:"elevation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Spherical3D) Reset()         { *m = Spherical3D{} }
func (m *Spherical3D) String() string { return proto.CompactTextString(m) }
func (*Spherical3D) ProtoMessage()    {}
func (*Spherical3D) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb1efe539488b99, []int{7}
}

func (m *Spherical3D) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spherical3D.Unmarshal(m, b)
}
func (m *Spherical3D) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spherical3D.Marshal(b, m, deterministic)
}
func (m *Spherical3D) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spherical3D.Merge(m, src)
}
func (m *Spherical3D) XXX_Size() int {
	return xxx_messageInfo_Spherical3D.Size(m)
}
func (m *Spherical3D) XXX_DiscardUnknown() {
	xxx_messageInfo_Spherical3D.DiscardUnknown(m)
}

var xxx_messageInfo_Spherical3D proto.InternalMessageInfo

func (m *Spherical3D) GetDistance() float64 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Spherical3D) GetAzimuth() float64 {
	if m != nil {
		return m.Azimuth
	}
	return 0
}

func (m *Spherical3D) GetElevation() float64 {
	if m != nil {
		return m.Elevation
	}
	return 0
}

//
// \brief The base attributes of a stationary object or entity.
//
// This includes the \c StationaryObject , \c TrafficSign ,
// \c TrafficLight , \c RoadMarking messages.
//
// \image html OSI_BaseStationary.jpg
//
// All coordinates and orientations from ground truth objects are relative to
// the global ground truth frame (see image). All coordinates and orientations
// from detected objects are relative to the host vehicle frame (see:
// \c MovingObject::Vehicle vehicle reference point).
//
type BaseStationary struct {
	// The 3D dimensions of the stationary object (bounding box), e.g. a
	// landmark.
	//
	Dimension *Dimension3D `protobuf:"bytes,1,opt,name=dimension,proto3" json:"dimension,omitempty"`
	// The reference point for position and orientation, i.e. the center (x,y,z)
	// of the bounding box.
	//
	Position *Vector3D `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	// The relative orientation of the stationary object w.r.t. its parent
	// frame, noted in the parent frame. The orientation becomes global/absolute
	// if the parent frame is inertial (all parent frames up to ground truth).
	//
	// <tt>Origin_base_stationary_entity :=
	// Rotation_yaw_pitch_roll(#orientation)*(Origin_parent_coordinate_system -
	// #position)</tt>
	//
	// \note There may be some constraints how to align the orientation w.r.t.
	// to some stationary object's or entity's definition.
	//
	Orientation *Orientation3D `protobuf:"bytes,3,opt,name=orientation,proto3" json:"orientation,omitempty"`
	// Usage as ground truth:
	// The two dimensional (flat) contour of the object. This is an extension of
	// the concept of a bounding box as defined by \c Dimension3d. The contour
	// is the projection of the object's outline onto the z-plane in the object
	// frame (independent of its current position and orientation). The height
	// is the same as the height of the bounding box.
	//
	// Usage as sensor data:
	// The polygon describes the visible part of the object's contour.
	//
	// General definitions:
	// The polygon is defined in the local object frame: x pointing forward and
	// y to the left.
	// The origin is the center of the bounding box.
	// As ground truth, the polygon is closed by connecting the last with the
	// first point. Therefore these two points must be different. The polygon
	// must consist of at least three points.
	// As sensor data, however, the polygon is open.
	// The polygon is defined counter-clockwise.
	//
	BasePolygon          []*Vector2D `protobuf:"bytes,4,rep,name=base_polygon,json=basePolygon,proto3" json:"base_polygon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BaseStationary) Reset()         { *m = BaseStationary{} }
func (m *BaseStationary) String() string { return proto.CompactTextString(m) }
func (*BaseStationary) ProtoMessage()    {}
func (*BaseStationary) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb1efe539488b99, []int{8}
}

func (m *BaseStationary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseStationary.Unmarshal(m, b)
}
func (m *BaseStationary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseStationary.Marshal(b, m, deterministic)
}
func (m *BaseStationary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseStationary.Merge(m, src)
}
func (m *BaseStationary) XXX_Size() int {
	return xxx_messageInfo_BaseStationary.Size(m)
}
func (m *BaseStationary) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseStationary.DiscardUnknown(m)
}

var xxx_messageInfo_BaseStationary proto.InternalMessageInfo

func (m *BaseStationary) GetDimension() *Dimension3D {
	if m != nil {
		return m.Dimension
	}
	return nil
}

func (m *BaseStationary) GetPosition() *Vector3D {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *BaseStationary) GetOrientation() *Orientation3D {
	if m != nil {
		return m.Orientation
	}
	return nil
}

func (m *BaseStationary) GetBasePolygon() []*Vector2D {
	if m != nil {
		return m.BasePolygon
	}
	return nil
}

//
// \brief The base attributes of an object that is moving.
//
// This includes the \c MovingObject messages.
//
// All coordinates and orientations from ground truth objects are relative to
// the global ground truth frame. All coordinates and orientations
// from detected objects are relative to the host vehicle frame
// (see: \c MovingObject::Vehicle vehicle reference point).
//
type BaseMoving struct {
	// The 3D dimension of the moving object (its bounding box).
	//
	Dimension *Dimension3D `protobuf:"bytes,1,opt,name=dimension,proto3" json:"dimension,omitempty"`
	// The reference point for position and orientation: the center (x,y,z) of
	// the bounding box.
	//
	Position *Vector3D `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	// The relative orientation of the moving object w.r.t. its parent frame,
	// noted in the parent frame. The orientation becomes global/absolute if
	// the parent frame is inertial (all parent frames up to ground truth).
	//
	// <tt>Origin_base_moving_entity :=
	// Rotation_yaw_pitch_roll(#orientation)*(Origin_parent_coordinate_system -
	// #position)</tt>
	//
	// \note There may be some constraints how to align the orientation w.r.t.
	// to some stationary object's or entity's definition.
	//
	Orientation *Orientation3D `protobuf:"bytes,3,opt,name=orientation,proto3" json:"orientation,omitempty"`
	// The relative velocity of the moving object w.r.t. the parent frame,
	// noted in the parent frame. The velocity becomes global/absolute if
	// the parent frame does is inertial (all parent frames up to ground truth).
	//
	// <tt>#position (t) := #position (t-dt)+ #velocity *dt</tt>
	//
	Velocity *Vector3D `protobuf:"bytes,4,opt,name=velocity,proto3" json:"velocity,omitempty"`
	// The relative acceleration of the moving object w.r.t. its parent frame,
	// noted in the parent frame. The acceleration becomes global/absolute if
	// the parent frame is inertial (all parent frames up to ground truth).
	//
	// <tt> #position (t) := #position (t-dt)+ #velocity *dt+ #acceleration
	// /2*dt^2</tt>
	//
	// <tt> #velocity (t) := #velocity (t-dt)+ #acceleration *dt</tt>
	//
	Acceleration *Vector3D `protobuf:"bytes,5,opt,name=acceleration,proto3" json:"acceleration,omitempty"`
	// The relative orientation rate of the moving object w.r.t. its parent
	// frame and parent orientation rate in the center point of the bounding box
	// (origin of the bounding box frame), noted in the parent frame.
	// The orientation becomes global/absolute if the parent frame is inertial
	// (all parent frames up to ground truth).
	//
	// <tt>Rotation_yaw_pitch_roll(#orientation (t)) :=
	// Rotation_yaw_pitch_roll(#orientation_rate
	// *dt)*Rotation_yaw_pitch_roll(#orientation (t-dt))</tt>
	//
	// \note <tt>#orientation (t)</tt> is \b not equal <tt>#orientation
	// (t-dt)+#orientation_rate *dt</tt>
	//
	OrientationRate *Orientation3D `protobuf:"bytes,6,opt,name=orientation_rate,json=orientationRate,proto3" json:"orientation_rate,omitempty"`
	// Usage as ground truth:
	// The two dimensional (flat) contour of the object. This is an extension of
	// the concept of a bounding box as defined by \c Dimension3d. The contour
	// is the projection of the object's outline onto the z-plane in the object
	// frame (independent of its current position and orientation). The height
	// is the same as the height of the bounding box.
	//
	// Usage as sensor data:
	// The polygon describes the visible part of the object's contour.
	//
	// General definitions:
	// The polygon is defined in the local object frame: x pointing forward and
	// y to the left. The origin is the center of the bounding box.
	// As ground truth, the polygon is closed by connecting the last with the
	// first point. Therefore these two points must be different. The polygon
	// must consist of at least three points. As sensor data, however, the
	// polygon is open.
	// The polygon is defined counter-clockwise.
	//
	BasePolygon          []*Vector2D `protobuf:"bytes,7,rep,name=base_polygon,json=basePolygon,proto3" json:"base_polygon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BaseMoving) Reset()         { *m = BaseMoving{} }
func (m *BaseMoving) String() string { return proto.CompactTextString(m) }
func (*BaseMoving) ProtoMessage()    {}
func (*BaseMoving) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb1efe539488b99, []int{9}
}

func (m *BaseMoving) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseMoving.Unmarshal(m, b)
}
func (m *BaseMoving) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseMoving.Marshal(b, m, deterministic)
}
func (m *BaseMoving) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseMoving.Merge(m, src)
}
func (m *BaseMoving) XXX_Size() int {
	return xxx_messageInfo_BaseMoving.Size(m)
}
func (m *BaseMoving) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseMoving.DiscardUnknown(m)
}

var xxx_messageInfo_BaseMoving proto.InternalMessageInfo

func (m *BaseMoving) GetDimension() *Dimension3D {
	if m != nil {
		return m.Dimension
	}
	return nil
}

func (m *BaseMoving) GetPosition() *Vector3D {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *BaseMoving) GetOrientation() *Orientation3D {
	if m != nil {
		return m.Orientation
	}
	return nil
}

func (m *BaseMoving) GetVelocity() *Vector3D {
	if m != nil {
		return m.Velocity
	}
	return nil
}

func (m *BaseMoving) GetAcceleration() *Vector3D {
	if m != nil {
		return m.Acceleration
	}
	return nil
}

func (m *BaseMoving) GetOrientationRate() *Orientation3D {
	if m != nil {
		return m.OrientationRate
	}
	return nil
}

func (m *BaseMoving) GetBasePolygon() []*Vector2D {
	if m != nil {
		return m.BasePolygon
	}
	return nil
}

func init() {
	proto.RegisterType((*Vector3D)(nil), "osi3.Vector3d")
	proto.RegisterType((*Vector2D)(nil), "osi3.Vector2d")
	proto.RegisterType((*Timestamp)(nil), "osi3.Timestamp")
	proto.RegisterType((*Dimension3D)(nil), "osi3.Dimension3d")
	proto.RegisterType((*Orientation3D)(nil), "osi3.Orientation3d")
	proto.RegisterType((*Identifier)(nil), "osi3.Identifier")
	proto.RegisterType((*MountingPosition)(nil), "osi3.MountingPosition")
	proto.RegisterType((*Spherical3D)(nil), "osi3.Spherical3d")
	proto.RegisterType((*BaseStationary)(nil), "osi3.BaseStationary")
	proto.RegisterType((*BaseMoving)(nil), "osi3.BaseMoving")
}

func init() { proto.RegisterFile("osi_common.proto", fileDescriptor_3fb1efe539488b99) }

var fileDescriptor_3fb1efe539488b99 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe5, 0xc4, 0x4d, 0x93, 0x71, 0x5a, 0xc2, 0x82, 0x90, 0x85, 0x38, 0x54, 0x7b, 0x40,
	0xa8, 0x87, 0x20, 0x6c, 0x38, 0x21, 0x71, 0xa8, 0x38, 0x80, 0x50, 0x45, 0xe5, 0x20, 0xae, 0xd1,
	0xd6, 0x1e, 0xec, 0x95, 0xec, 0x5d, 0xcb, 0xbb, 0x49, 0xeb, 0xbc, 0x28, 0x0f, 0xc2, 0x0b, 0xa0,
	0xf5, 0xae, 0xf3, 0x01, 0xe5, 0xeb, 0xc6, 0x6d, 0x7e, 0xde, 0xff, 0xcc, 0x7f, 0x3e, 0x24, 0xc3,
	0x4c, 0x2a, 0xbe, 0x4c, 0x65, 0x55, 0x49, 0x31, 0xaf, 0x1b, 0xa9, 0x25, 0xf1, 0xa5, 0xe2, 0x31,
	0x7d, 0x09, 0xe3, 0xcf, 0x98, 0x6a, 0xd9, 0xc4, 0x19, 0x99, 0x82, 0x77, 0x1b, 0x7a, 0x67, 0xde,
	0x33, 0x2f, 0xf1, 0x6e, 0x0d, 0xb5, 0xe1, 0xc0, 0x52, 0x6b, 0x68, 0x13, 0x0e, 0x2d, 0x6d, 0xe8,
	0xd3, 0x3e, 0x2b, 0xfa, 0x6d, 0x16, 0x7d, 0x0d, 0x93, 0x4f, 0xbc, 0x42, 0xa5, 0x59, 0x55, 0x93,
	0x10, 0x8e, 0x15, 0xa6, 0x52, 0x64, 0xaa, 0x93, 0x0f, 0x93, 0x1e, 0xc9, 0x43, 0x38, 0x12, 0x4c,
	0x48, 0xd5, 0x25, 0x9e, 0x24, 0x16, 0xe8, 0x02, 0x82, 0xb7, 0xbc, 0x42, 0xa1, 0xb8, 0x14, 0x71,
	0x46, 0x1e, 0xc1, 0xa8, 0x44, 0x91, 0xeb, 0xc2, 0x99, 0x39, 0x32, 0xc9, 0x37, 0x3c, 0xd3, 0x85,
	0x73, 0xb5, 0x60, 0xd4, 0x05, 0xf2, 0xbc, 0xd0, 0xae, 0x69, 0x47, 0xf4, 0x03, 0x9c, 0x7c, 0x6c,
	0x38, 0x0a, 0xcd, 0xb4, 0x2d, 0x4b, 0xc0, 0x6f, 0x64, 0x59, 0xba, 0xa2, 0x5d, 0x6c, 0x4a, 0xd6,
	0x5c, 0xa7, 0xdb, 0x92, 0x1d, 0x90, 0x19, 0x0c, 0x5b, 0x76, 0xe3, 0xea, 0x99, 0x90, 0x52, 0x80,
	0xf7, 0x19, 0x0a, 0xcd, 0xbf, 0x70, 0x6c, 0x4c, 0xd6, 0x9a, 0x95, 0x2b, 0xec, 0x4a, 0xf9, 0x89,
	0x05, 0xba, 0x82, 0xd9, 0xa5, 0x5c, 0x09, 0xcd, 0x45, 0x7e, 0x25, 0x15, 0x37, 0xae, 0xe4, 0x1c,
	0xc6, 0xb5, 0x8b, 0x3b, 0x71, 0x10, 0x9d, 0xce, 0xcd, 0x35, 0xe6, 0xfd, 0x29, 0x92, 0xed, 0x3b,
	0x79, 0x05, 0x81, 0xdc, 0x35, 0xdc, 0x75, 0x14, 0x44, 0x0f, 0xac, 0xfc, 0x60, 0x92, 0x64, 0x5f,
	0x47, 0x19, 0x04, 0x8b, 0xba, 0xc0, 0x86, 0xa7, 0xac, 0x8c, 0x33, 0xf2, 0x18, 0xc6, 0x19, 0x57,
	0x9a, 0x89, 0x14, 0xdd, 0xa4, 0x5b, 0x36, 0x77, 0x61, 0x1b, 0x5e, 0xad, 0xb6, 0x2b, 0xec, 0x91,
	0x3c, 0x81, 0x09, 0x96, 0xb8, 0xb6, 0xce, 0x76, 0xee, 0xdd, 0x07, 0xfa, 0xd5, 0x83, 0xd3, 0x0b,
	0xa6, 0x70, 0x61, 0x2d, 0x59, 0xd3, 0x92, 0xe7, 0x30, 0xc9, 0xfa, 0x93, 0xb9, 0xc9, 0xee, 0xdb,
	0x56, 0xf7, 0x2e, 0x99, 0xec, 0x34, 0x07, 0x9b, 0x18, 0xfc, 0xdb, 0x26, 0x86, 0x7f, 0xb7, 0x09,
	0xf2, 0x02, 0xa6, 0xd7, 0x4c, 0xe1, 0xb2, 0x96, 0x65, 0x9b, 0x4b, 0x11, 0xfa, 0x67, 0xc3, 0x1f,
	0x6d, 0xa2, 0x2c, 0x09, 0x8c, 0xe6, 0xca, 0x4a, 0xe8, 0xb7, 0x01, 0x80, 0x99, 0xec, 0x52, 0xae,
	0xb9, 0xc8, 0xff, 0xcb, 0xa9, 0xce, 0x61, 0xbc, 0xc6, 0x52, 0xa6, 0x5c, 0xb7, 0xa1, 0x7f, 0xb7,
	0x45, 0xff, 0x4e, 0x22, 0x98, 0xb2, 0x34, 0xc5, 0x12, 0x1b, 0xeb, 0x71, 0x74, 0xa7, 0xfe, 0x40,
	0x43, 0xde, 0xc0, 0x6c, 0xcf, 0x6e, 0xd9, 0x30, 0x8d, 0xe1, 0xe8, 0xd7, 0xbd, 0xdd, 0xdb, 0x13,
	0x27, 0x4c, 0xe3, 0x4f, 0x5b, 0x3f, 0xfe, 0xe3, 0xd6, 0x2f, 0x06, 0xef, 0xbc, 0xeb, 0x51, 0xf7,
	0x6f, 0x8a, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x8d, 0x0a, 0xdb, 0xaf, 0x04, 0x00, 0x00,
}