// Package GLTFObjectModelProperty provides methods for working with GLTFObjectModelProperty object instances.
package GLTFObjectModelProperty

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
GLTFObjectModelProperty defines a mapping between a property in the glTF object model and a NodePath in the Godot scene tree. This can be used to animate properties in a glTF file using the [code]KHR_animation_pointer[/code] extension, or to access them through an engine-agnostic script such as a behavior graph as defined by the [code]KHR_interactivity[/code] extension.
The glTF property is identified by JSON pointer(s) stored in [member json_pointers], while the Godot property it maps to is defined by [member node_paths]. In most cases [member json_pointers] and [member node_paths] will each only have one item, but in some cases a single glTF JSON pointer will map to multiple Godot properties, or a single Godot property will be mapped to multiple glTF JSON pointers, or it might be a many-to-many relationship.
[Expression] objects can be used to define conversions between the data, such as when glTF defines an angle in radians and Godot uses degrees. The [member object_model_type] property defines the type of data stored in the glTF file as defined by the object model, see [enum GLTFObjectModelType] for possible values.
*/
type Instance [1]gdclass.GLTFObjectModelProperty

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFObjectModelProperty() Instance
}

/*
Appends a [NodePath] to [member node_paths]. This can be used by [GLTFDocumentExtension] classes to define how a glTF object model property maps to a Godot property, or multiple Godot properties. Prefer using [method append_path_to_property] for simple cases. Be sure to also call [method set_types] once (the order does not matter).
*/
func (self Instance) AppendNodePath(node_path string) { //gd:GLTFObjectModelProperty.append_node_path
	Advanced(self).AppendNodePath(Path.ToNode(String.New(node_path)))
}

/*
High-level wrapper over [method append_node_path] that handles the most common cases. It constructs a new [NodePath] using [param node_path] as a base and appends [param prop_name] to the subpath. Be sure to also call [method set_types] once (the order does not matter).
*/
func (self Instance) AppendPathToProperty(node_path string, prop_name string) { //gd:GLTFObjectModelProperty.append_path_to_property
	Advanced(self).AppendPathToProperty(Path.ToNode(String.New(node_path)), String.Name(String.New(prop_name)))
}

/*
The GLTF accessor type associated with this property's [member object_model_type]. See [member GLTFAccessor.accessor_type] for possible values, and see [enum GLTFObjectModelType] for how the object model type maps to accessor types.
*/
func (self Instance) GetAccessorType() gdclass.GLTFAccessorGLTFAccessorType { //gd:GLTFObjectModelProperty.get_accessor_type
	return gdclass.GLTFAccessorGLTFAccessorType(Advanced(self).GetAccessorType())
}

/*
Returns [code]true[/code] if [member node_paths] is not empty. This is used during import to determine if a [GLTFObjectModelProperty] can handle converting a glTF object model property to a Godot property.
*/
func (self Instance) HasNodePaths() bool { //gd:GLTFObjectModelProperty.has_node_paths
	return bool(Advanced(self).HasNodePaths())
}

/*
Returns [code]true[/code] if [member json_pointers] is not empty. This is used during export to determine if a [GLTFObjectModelProperty] can handle converting a Godot property to a glTF object model property.
*/
func (self Instance) HasJsonPointers() bool { //gd:GLTFObjectModelProperty.has_json_pointers
	return bool(Advanced(self).HasJsonPointers())
}

/*
Sets the [member variant_type] and [member object_model_type] properties. This is a convenience method to set both properties at once, since they are almost always known at the same time. This method should be called once. Calling it again with the same values will have no effect.
*/
func (self Instance) SetTypes(variant_type variant.Type, obj_model_type gdclass.GLTFObjectModelPropertyGLTFObjectModelType) { //gd:GLTFObjectModelProperty.set_types
	Advanced(self).SetTypes(variant_type, obj_model_type)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFObjectModelProperty

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFObjectModelProperty"))
	casted := Instance{*(*gdclass.GLTFObjectModelProperty)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) GltfToGodotExpression() [1]gdclass.Expression {
	return [1]gdclass.Expression(class(self).GetGltfToGodotExpression())
}

func (self Instance) SetGltfToGodotExpression(value [1]gdclass.Expression) {
	class(self).SetGltfToGodotExpression(value)
}

func (self Instance) GodotToGltfExpression() [1]gdclass.Expression {
	return [1]gdclass.Expression(class(self).GetGodotToGltfExpression())
}

func (self Instance) SetGodotToGltfExpression(value [1]gdclass.Expression) {
	class(self).SetGodotToGltfExpression(value)
}

func (self Instance) NodePaths() []string {
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetNodePaths())))
}

func (self Instance) SetNodePaths(value []string) {
	class(self).SetNodePaths(gd.ArrayFromSlice[Array.Contains[Path.ToNode]](value))
}

func (self Instance) ObjectModelType() gdclass.GLTFObjectModelPropertyGLTFObjectModelType {
	return gdclass.GLTFObjectModelPropertyGLTFObjectModelType(class(self).GetObjectModelType())
}

func (self Instance) SetObjectModelType(value gdclass.GLTFObjectModelPropertyGLTFObjectModelType) {
	class(self).SetObjectModelType(value)
}

func (self Instance) JsonPointers() [][]string {
	return [][]string(gd.ArrayAs[[][]string](gd.InternalArray(class(self).GetJsonPointers())))
}

func (self Instance) SetJsonPointers(value [][]string) {
	class(self).SetJsonPointers(gd.ArrayFromSlice[Array.Contains[Packed.Strings]](value))
}

func (self Instance) VariantType() variant.Type {
	return variant.Type(class(self).GetVariantType())
}

func (self Instance) SetVariantType(value variant.Type) {
	class(self).SetVariantType(value)
}

/*
Appends a [NodePath] to [member node_paths]. This can be used by [GLTFDocumentExtension] classes to define how a glTF object model property maps to a Godot property, or multiple Godot properties. Prefer using [method append_path_to_property] for simple cases. Be sure to also call [method set_types] once (the order does not matter).
*/
//go:nosplit
func (self class) AppendNodePath(node_path Path.ToNode) { //gd:GLTFObjectModelProperty.append_node_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(node_path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_append_node_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
High-level wrapper over [method append_node_path] that handles the most common cases. It constructs a new [NodePath] using [param node_path] as a base and appends [param prop_name] to the subpath. Be sure to also call [method set_types] once (the order does not matter).
*/
//go:nosplit
func (self class) AppendPathToProperty(node_path Path.ToNode, prop_name String.Name) { //gd:GLTFObjectModelProperty.append_path_to_property
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(node_path)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(prop_name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_append_path_to_property, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
The GLTF accessor type associated with this property's [member object_model_type]. See [member GLTFAccessor.accessor_type] for possible values, and see [enum GLTFObjectModelType] for how the object model type maps to accessor types.
*/
//go:nosplit
func (self class) GetAccessorType() gdclass.GLTFAccessorGLTFAccessorType { //gd:GLTFObjectModelProperty.get_accessor_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GLTFAccessorGLTFAccessorType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_get_accessor_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetGltfToGodotExpression() [1]gdclass.Expression { //gd:GLTFObjectModelProperty.get_gltf_to_godot_expression
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_get_gltf_to_godot_expression, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Expression{gd.PointerWithOwnershipTransferredToGo[gdclass.Expression](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGltfToGodotExpression(gltf_to_godot_expr [1]gdclass.Expression) { //gd:GLTFObjectModelProperty.set_gltf_to_godot_expression
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gltf_to_godot_expr[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_set_gltf_to_godot_expression, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGodotToGltfExpression() [1]gdclass.Expression { //gd:GLTFObjectModelProperty.get_godot_to_gltf_expression
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_get_godot_to_gltf_expression, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Expression{gd.PointerWithOwnershipTransferredToGo[gdclass.Expression](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGodotToGltfExpression(godot_to_gltf_expr [1]gdclass.Expression) { //gd:GLTFObjectModelProperty.set_godot_to_gltf_expression
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(godot_to_gltf_expr[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_set_godot_to_gltf_expression, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNodePaths() Array.Contains[Path.ToNode] { //gd:GLTFObjectModelProperty.get_node_paths
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_get_node_paths, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Path.ToNode]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [member node_paths] is not empty. This is used during import to determine if a [GLTFObjectModelProperty] can handle converting a glTF object model property to a Godot property.
*/
//go:nosplit
func (self class) HasNodePaths() bool { //gd:GLTFObjectModelProperty.has_node_paths
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_has_node_paths, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNodePaths(node_paths Array.Contains[Path.ToNode]) { //gd:GLTFObjectModelProperty.set_node_paths
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(node_paths)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_set_node_paths, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetObjectModelType() gdclass.GLTFObjectModelPropertyGLTFObjectModelType { //gd:GLTFObjectModelProperty.get_object_model_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GLTFObjectModelPropertyGLTFObjectModelType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_get_object_model_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetObjectModelType(atype gdclass.GLTFObjectModelPropertyGLTFObjectModelType) { //gd:GLTFObjectModelProperty.set_object_model_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_set_object_model_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJsonPointers() Array.Contains[Packed.Strings] { //gd:GLTFObjectModelProperty.get_json_pointers
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_get_json_pointers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Packed.Strings]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [member json_pointers] is not empty. This is used during export to determine if a [GLTFObjectModelProperty] can handle converting a Godot property to a glTF object model property.
*/
//go:nosplit
func (self class) HasJsonPointers() bool { //gd:GLTFObjectModelProperty.has_json_pointers
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_has_json_pointers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJsonPointers(json_pointers Array.Contains[Packed.Strings]) { //gd:GLTFObjectModelProperty.set_json_pointers
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(json_pointers)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_set_json_pointers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVariantType() variant.Type { //gd:GLTFObjectModelProperty.get_variant_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[variant.Type](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_get_variant_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVariantType(variant_type variant.Type) { //gd:GLTFObjectModelProperty.set_variant_type
	var frame = callframe.New()
	callframe.Arg(frame, variant_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_set_variant_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [member variant_type] and [member object_model_type] properties. This is a convenience method to set both properties at once, since they are almost always known at the same time. This method should be called once. Calling it again with the same values will have no effect.
*/
//go:nosplit
func (self class) SetTypes(variant_type variant.Type, obj_model_type gdclass.GLTFObjectModelPropertyGLTFObjectModelType) { //gd:GLTFObjectModelProperty.set_types
	var frame = callframe.New()
	callframe.Arg(frame, variant_type)
	callframe.Arg(frame, obj_model_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFObjectModelProperty.Bind_set_types, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsGLTFObjectModelProperty() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFObjectModelProperty() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("GLTFObjectModelProperty", func(ptr gd.Object) any {
		return [1]gdclass.GLTFObjectModelProperty{*(*gdclass.GLTFObjectModelProperty)(unsafe.Pointer(&ptr))}
	})
}

type GLTFObjectModelType = gdclass.GLTFObjectModelPropertyGLTFObjectModelType //gd:GLTFObjectModelProperty.GLTFObjectModelType

const (
	/*Unknown or not set object model type. If the object model type is set to this value, the real type still needs to be determined.*/
	GltfObjectModelTypeUnknown GLTFObjectModelType = 0
	/*Object model type "bool". Represented in the glTF JSON as a boolean, and encoded in a [GLTFAccessor] as "SCALAR". When encoded in an accessor, a value of [code]0[/code] is [code]false[/code], and any other value is [code]true[/code].*/
	GltfObjectModelTypeBool GLTFObjectModelType = 1
	/*Object model type "float". Represented in the glTF JSON as a number, and encoded in a [GLTFAccessor] as "SCALAR".*/
	GltfObjectModelTypeFloat GLTFObjectModelType = 2
	/*Object model type "float[lb][rb]". Represented in the glTF JSON as an array of numbers, and encoded in a [GLTFAccessor] as "SCALAR".*/
	GltfObjectModelTypeFloatArray GLTFObjectModelType = 3
	/*Object model type "float2". Represented in the glTF JSON as an array of two numbers, and encoded in a [GLTFAccessor] as "VEC2".*/
	GltfObjectModelTypeFloat2 GLTFObjectModelType = 4
	/*Object model type "float3". Represented in the glTF JSON as an array of three numbers, and encoded in a [GLTFAccessor] as "VEC3".*/
	GltfObjectModelTypeFloat3 GLTFObjectModelType = 5
	/*Object model type "float4". Represented in the glTF JSON as an array of four numbers, and encoded in a [GLTFAccessor] as "VEC4".*/
	GltfObjectModelTypeFloat4 GLTFObjectModelType = 6
	/*Object model type "float2x2". Represented in the glTF JSON as an array of four numbers, and encoded in a [GLTFAccessor] as "MAT2".*/
	GltfObjectModelTypeFloat2x2 GLTFObjectModelType = 7
	/*Object model type "float3x3". Represented in the glTF JSON as an array of nine numbers, and encoded in a [GLTFAccessor] as "MAT3".*/
	GltfObjectModelTypeFloat3x3 GLTFObjectModelType = 8
	/*Object model type "float4x4". Represented in the glTF JSON as an array of sixteen numbers, and encoded in a [GLTFAccessor] as "MAT4".*/
	GltfObjectModelTypeFloat4x4 GLTFObjectModelType = 9
	/*Object model type "int". Represented in the glTF JSON as a number, and encoded in a [GLTFAccessor] as "SCALAR". The range of values is limited to signed integers. For [code]KHR_interactivity[/code], only 32-bit integers are supported.*/
	GltfObjectModelTypeInt GLTFObjectModelType = 10
)

type VariantType int

const (
	/*Variable is [code]null[/code].*/
	TypeNil VariantType = 0
	/*Variable is of type [bool].*/
	TypeBool VariantType = 1
	/*Variable is of type [int].*/
	TypeInt VariantType = 2
	/*Variable is of type [float].*/
	TypeFloat VariantType = 3
	/*Variable is of type [String].*/
	TypeString VariantType = 4
	/*Variable is of type [Vector2].*/
	TypeVector2 VariantType = 5
	/*Variable is of type [Vector2i].*/
	TypeVector2i VariantType = 6
	/*Variable is of type [Rect2].*/
	TypeRect2 VariantType = 7
	/*Variable is of type [Rect2i].*/
	TypeRect2i VariantType = 8
	/*Variable is of type [Vector3].*/
	TypeVector3 VariantType = 9
	/*Variable is of type [Vector3i].*/
	TypeVector3i VariantType = 10
	/*Variable is of type [Transform2D].*/
	TypeTransform2d VariantType = 11
	/*Variable is of type [Vector4].*/
	TypeVector4 VariantType = 12
	/*Variable is of type [Vector4i].*/
	TypeVector4i VariantType = 13
	/*Variable is of type [Plane].*/
	TypePlane VariantType = 14
	/*Variable is of type [Quaternion].*/
	TypeQuaternion VariantType = 15
	/*Variable is of type [AABB].*/
	TypeAabb VariantType = 16
	/*Variable is of type [Basis].*/
	TypeBasis VariantType = 17
	/*Variable is of type [Transform3D].*/
	TypeTransform3d VariantType = 18
	/*Variable is of type [Projection].*/
	TypeProjection VariantType = 19
	/*Variable is of type [Color].*/
	TypeColor VariantType = 20
	/*Variable is of type [StringName].*/
	TypeStringName VariantType = 21
	/*Variable is of type [NodePath].*/
	TypeNodePath VariantType = 22
	/*Variable is of type [RID].*/
	TypeRid VariantType = 23
	/*Variable is of type [Object].*/
	TypeObject VariantType = 24
	/*Variable is of type [Callable].*/
	TypeCallable VariantType = 25
	/*Variable is of type [Signal].*/
	TypeSignal VariantType = 26
	/*Variable is of type [Dictionary].*/
	TypeDictionary VariantType = 27
	/*Variable is of type [Array].*/
	TypeArray VariantType = 28
	/*Variable is of type [PackedByteArray].*/
	TypePackedByteArray VariantType = 29
	/*Variable is of type [PackedInt32Array].*/
	TypePackedInt32Array VariantType = 30
	/*Variable is of type [PackedInt64Array].*/
	TypePackedInt64Array VariantType = 31
	/*Variable is of type [PackedFloat32Array].*/
	TypePackedFloat32Array VariantType = 32
	/*Variable is of type [PackedFloat64Array].*/
	TypePackedFloat64Array VariantType = 33
	/*Variable is of type [PackedStringArray].*/
	TypePackedStringArray VariantType = 34
	/*Variable is of type [PackedVector2Array].*/
	TypePackedVector2Array VariantType = 35
	/*Variable is of type [PackedVector3Array].*/
	TypePackedVector3Array VariantType = 36
	/*Variable is of type [PackedColorArray].*/
	TypePackedColorArray VariantType = 37
	/*Variable is of type [PackedVector4Array].*/
	TypePackedVector4Array VariantType = 38
	/*Represents the size of the [enum Variant.Type] enum.*/
	TypeMax VariantType = 39
)
