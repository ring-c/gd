//go:build !generate

package gd

import (
	"reflect"
	"unsafe"

	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
	PackedType "graphics.gd/variant/Packed"

	"runtime.link/api"
)

type VariantType int

func (t VariantType) String() string {
	return Global.Variants.GetTypeName(t).String()
}

type Address uintptr

// API specification for Godot's GDExtension.
type API struct {
	api.Specification

	GetGodotVersion     func() Version
	GetNativeStructSize func(StringName) uintptr

	Memory struct {
		Allocate   func(uintptr) Address
		Reallocate func(Address, uintptr) Address
		Free       func(Address)

		Index func(addr Address, n int, size uintptr) unsafe.Pointer
		Write func(dst Address, src unsafe.Pointer, size uintptr)
	}

	PrintError              func(code, function, file string, line int32, notifyEditor bool)
	PrintErrorMessage       func(code, message, function, file string, line int32, notifyEditor bool)
	PrintWarning            func(code, function, file string, line int32, notifyEditor bool)
	PrintWarningMessage     func(code, message, function, file string, line int32, notifyEditor bool)
	PrintScriptError        func(code, function, file string, line int32, notifyEditor bool)
	PrintScriptErrorMessage func(code, message, function, file string, line int32, notifyEditor bool)

	Variants struct {
		NewCopy                   func(src Variant) Variant
		NewNil                    func() Variant
		Destroy                   func(self Variant)
		Call                      func(self Variant, method StringName, args ...Variant) (Variant, error)
		CallStatic                func(vtype VariantType, method StringName, args ...Variant) (Variant, error)
		Evaluate                  func(operator Operator, a, b Variant) (ret Variant, ok bool)
		Set                       func(self, key, val Variant) bool
		SetNamed                  func(self Variant, key StringName, val Variant) bool
		SetKeyed                  func(self, key, val Variant) bool
		SetIndexed                func(self Variant, index Int, val Variant) (ok, oob bool)
		Get                       func(self, key Variant) (Variant, bool)
		GetNamed                  func(self Variant, key StringName) (Variant, bool)
		GetKeyed                  func(self, key Variant) (Variant, bool)
		GetIndexed                func(self Variant, index Int) (val Variant, ok, oob bool)
		IteratorInitialize        func(self Variant) (Variant, bool)
		IteratorNext              func(self Variant, iterator Variant) bool
		IteratorGet               func(self Variant, iterator Variant) (Variant, bool)
		Hash                      func(self Variant) Int
		RecursiveHash             func(self Variant, count Int) Int
		HashCompare               func(self, variant Variant) bool
		Booleanize                func(self Variant) bool
		Duplicate                 func(self Variant, deep bool) Variant
		Stringify                 func(self Variant) String
		GetType                   func(self Variant) VariantType
		HasMethod                 func(self Variant, method StringName) bool
		HasMember                 func(self Variant, member StringName) bool
		HasKey                    func(self Variant, key Variant) (hasKey, valid bool)
		GetTypeName               func(self VariantType) String
		CanConvert                func(self Variant, to VariantType) bool
		CanConvertStrict          func(self Variant, to VariantType) bool
		FromTypeConstructor       func(VariantType) func(ret callframe.Ptr[VariantPointers], arg callframe.Addr)
		ToTypeConstructor         func(VariantType) func(ret callframe.Addr, arg callframe.Ptr[VariantPointers])
		PointerOperatorEvaluator  func(op Operator, a, b VariantType) func(a, b, ret callframe.Addr)
		GetPointerBuiltinMethod   func(VariantType, StringName, Int) func(base callframe.Addr, args callframe.Args, ret callframe.Addr, c int32)
		GetPointerConstructor     func(vtype VariantType, index int32) func(base callframe.Addr, args callframe.Args)
		GetPointerDestructor      func(VariantType) func(base callframe.Addr)
		Construct                 func(t VariantType, args ...Variant) (Variant, error)
		GetPointerSetter          func(VariantType, StringName) func(base, arg callframe.Addr)
		GetPointerGetter          func(VariantType, StringName) func(base, ret callframe.Addr)
		GetPointerIndexedSetter   func(VariantType) func(base callframe.Addr, index Int, arg callframe.Addr)
		GetPointerIndexedGetter   func(VariantType) func(base callframe.Addr, index Int, ret callframe.Addr)
		GetPointerKeyedSetter     func(VariantType) func(base, key, arg callframe.Addr)
		GetPointerKeyedGetter     func(VariantType) func(base, key, ret callframe.Addr)
		GetPointerKeyedChecker    func(VariantType) func(base, key callframe.Addr) uint32
		GetConstantValue          func(t VariantType, name StringName) Variant
		GetPointerUtilityFunction func(name StringName, hash Int) func(ret callframe.Addr, args callframe.Args, c int32)
	}
	Strings struct {
		New        func(string) String
		Get        func(String) string
		SetIndex   func(String, Int, rune)
		Index      func(String, Int) rune
		Append     func(String, String)
		AppendRune func(String, rune)
		Resize     func(String, Int)
	}
	StringNames struct {
		New func(string) StringName
	}
	XMLParser struct {
		OpenBuffer func(Object, []byte) error
	}
	FileAccess struct {
		StoreBuffer func(Object, []byte)
		GetBuffer   func(Object, []byte) int
	}
	PackedByteArray    PackedFunctionsFor[PackedByteArray, byte]
	PackedColorArray   PackedFunctionsFor[PackedColorArray, Color]
	PackedFloat32Array PackedFunctionsFor[PackedFloat32Array, float32]
	PackedFloat64Array PackedFunctionsFor[PackedFloat64Array, float64]
	PackedInt32Array   PackedFunctionsFor[PackedInt32Array, int32]
	PackedInt64Array   PackedFunctionsFor[PackedInt64Array, int64]
	PackedStringArray  struct {
		Index         func(PackedStringArray, Int) String
		SetIndex      func(PackedStringArray, Int, String)
		CopyAsSlice   func(PackedStringArray) []String
		CopyFromSlice func(PackedStringArray, []String)
	}
	PackedVector2Array PackedFunctionsFor[PackedVector2Array, Vector2]
	PackedVector3Array PackedFunctionsFor[PackedVector3Array, Vector3]
	PackedVector4Array PackedFunctionsFor[PackedVector4Array, Vector4]
	Array              struct {
		Index    func(Array, Int) Variant
		Set      func(self, from Array)
		SetIndex func(Array, Int, Variant)
		SetTyped func(self Array, t VariantType, className StringName, script Object)
	}
	Dictionary struct {
		Index    func(dict Dictionary, key Variant) Variant
		SetIndex func(dict Dictionary, key, val Variant)
	}
	Object struct {
		MethodBindCall              func(method MethodBind, obj [1]Object, arg ...Variant) (Variant, error)
		MethodBindPointerCall       func(method MethodBind, obj [1]Object, arg callframe.Args, ret callframe.Addr)
		MethodBindPointerCallStatic func(method MethodBind, arg callframe.Args, ret callframe.Addr)
		Destroy                     func([1]Object)
		GetSingleton                func(name StringName) [1]Object
		GetInstanceBinding          func([1]Object, ExtensionToken, InstanceBindingType) any
		SetInstanceBinding          func([1]Object, ExtensionToken, any, InstanceBindingType)
		FreeInstanceBinding         func([1]Object, ExtensionToken)
		SetInstance                 func([1]Object, StringName, ObjectInterface)
		GetClassName                func([1]Object, ExtensionToken) String
		CastTo                      func([1]Object, ClassTag) [1]Object
		GetInstanceID               func([1]Object) ObjectID
		GetInstanceFromID           func(ObjectID) [1]Object
	}
	RefCounted struct {
		GetObject func([1]Object) [1]Object
		SetObject func([1]Object, [1]Object)
	}
	Callables struct {
		Create func(fn func(...Variant) (Variant, error)) Callable
		Get    func(Callable) (func(...Variant) (Variant, error), bool)
	}
	ClassDB struct {
		ConstructObject func(StringName) [1]Object
		GetClassTag     func(StringName) ClassTag
		GetMethodBind   func(class, method StringName, hash Int) MethodBind

		RegisterClass                 func(library ExtensionToken, name, extends StringName, info ClassInterface)
		RegisterClassMethod           func(library ExtensionToken, class StringName, info Method)
		RegisterClassIntegerConstant  func(library ExtensionToken, class, enum, name StringName, value int64, bitfield bool)
		RegisterClassProperty         func(library ExtensionToken, class StringName, info PropertyInfo, getter, setter StringName)
		RegisterClassPropertyIndexed  func(library ExtensionToken, class StringName, info PropertyInfo, getter, setter StringName, index int64)
		RegisterClassPropertyGroup    func(library ExtensionToken, class StringName, group, prefix String)
		RegisterClassPropertySubGroup func(library ExtensionToken, class StringName, subGroup, prefix String)
		RegisterClassSignal           func(library ExtensionToken, class, signal StringName, args []PropertyInfo)
		UnregisterClass               func(library ExtensionToken, class StringName)
	}
	EditorPlugins struct {
		Add    func(plugin StringName)
		Remove func(plugin StringName)
	}
	EditorHelp struct {
		Load func([]byte)
	}

	GetLibraryPath func(ExtensionToken) String

	// The following fields are primarily reserved for internal use within the gd module,
	// no backwards compatibility is guaranteed for these fields.

	ExtensionToken
	cache

	refCountedClassTag ClassTag

	// extensions instances are mapped here.
	Singletons singletons
}

type Packed[T any, V PackedType.Type] interface {
	PackedByteArray | PackedInt32Array | PackedInt64Array | PackedFloat32Array |
		PackedFloat64Array | PackedStringArray |
		PackedVector2Array | PackedVector3Array | PackedVector4Array |
		PackedColorArray

	pointers.Generic[T, PackedPointers]

	New() T
	Len() int
	Resize(Int) Int
	Index(Int) V
	SetIndex(Int, V)
}

type PackedFunctionsFor[T Packed[T, V], V PackedType.Type] struct {
	Index         func(T, Int) V
	SetIndex      func(T, Int, V)
	CopyAsSlice   func(T) []V
	CopyFromSlice func(T, []V)
}

type (
	StringPtr     *uintptr
	StringNamePtr *uintptr
	VariantPtr    *[3]uintptr
)

type CallErrorType int32

const (
	OK CallErrorType = iota
	ErrInvalidMethod
	ErrInvalidArgument
	ErrTooManyArguments
	ErrTooFewArguments
	ErrInstanceIsNil
	ErrMethodNotConst
)

type CallError struct {
	ErrorType CallErrorType
	Argument  int32
	Expected  int32
}

type InstanceBindingType unsafe.Pointer

func (err CallError) Error() string {
	switch err.ErrorType {
	case ErrInvalidMethod:
		return "invalid method"
	case ErrInvalidArgument:
		return "invalid argument"
	case ErrTooManyArguments:
		return "too many arguments"
	case ErrTooFewArguments:
		return "too few arguments"
	case ErrInstanceIsNil:
		return "instance is nil"
	case ErrMethodNotConst:
		return "method not const"
	default:
		return "unknown error"
	}
}

type Version struct {
	Major uint32
	Minor uint32
	Patch uint32
	Value string
}

func (v Version) String() string {
	return v.Value
}

type MethodBind uintptr

type ClassTag uintptr
type ExtensionToken uintptr

type InstanceID uint64

type Operator int32

const (
	Equal Operator = iota
	NotEqual
	Less
	LessEqual
	Greater
	GreaterEqual
	Add
	Subtract
	Multiply
	Divide
	Negate
	Positive
	Module
	Power
	ShiftLeft
	ShiftRight
	BitAnd
	BitOr
	BitXor
	BitNegate
	LogicalAnd
	LogicalOr
	LogicalXor
	LogicalNegate
	In
)

type GDExtensionInitializationLevel int64

const (
	GDExtensionInitializationLevelCore    GDExtensionInitializationLevel = 0
	GDExtensionInitializationLevelServers GDExtensionInitializationLevel = 1
	GDExtensionInitializationLevelScene   GDExtensionInitializationLevel = 2
	GDExtensionInitializationLevelEditor  GDExtensionInitializationLevel = 3
)

type PropertyInfo struct {
	Type       VariantType
	Name       StringName
	ClassName  StringName
	Hint       int64
	HintString String
	Usage      int64
}

type MethodFlags int64

type MethodInfo struct {
	Name             StringName
	ReturnValue      PropertyInfo
	Flags            MethodFlags
	ID               int32
	Arguments        []PropertyInfo
	DefaultArguments []Variant
}

type ClassInterface interface {
	IsVirtual() bool
	IsAbstract() bool
	IsExposed() bool

	CreateInstance() [1]Object
	GetVirtual(StringName) any
}

type ObjectInterface interface {
	OnCreate(reflect.Value)
	Set(StringName, Variant) bool
	Get(StringName) (Variant, bool)
	GetPropertyList() []PropertyInfo
	PropertyCanRevert(StringName) bool
	PropertyGetRevert(StringName) (Variant, bool)
	ValidateProperty(*PropertyInfo) bool
	Notification(int32, bool)
	ToString() (String, bool)
	Reference()
	Unreference()
	CallVirtual(StringName, any, Address, Address)
	GetRID() RID
	Free()
}

type ExtensionClassCallVirtualFunc func(any, Address, Address)

type ClassMethodArgumentMetadata uint32

const (
	ArgumentMetadataNone ClassMethodArgumentMetadata = iota
	ArgumentMetadataIntIsInt8
	ArgumentMetadataIntIsInt16
	ArgumentMetadataIntIsInt32
	ArgumentMetadataIntIsInt64
	ArgumentMetadataIntIsUint8
	ArgumentMetadataIntIsUint16
	ArgumentMetadataIntIsUint32
	ArgumentMetadataIntIsUint64
	ArgumentMetadataRealIsFloat32
	ArgumentMetadataRealIsFloat64
)

type Method struct {
	Name                StringName
	Call                func(any, ...Variant) (Variant, error)
	PointerCall         func(any, Address, Address)
	MethodFlags         MethodFlags
	ReturnValueInfo     *PropertyInfo
	ReturnValueMetadata ClassMethodArgumentMetadata

	Arguments         []PropertyInfo
	ArgumentsMetadata []ClassMethodArgumentMetadata

	DefaultArguments []Variant
}
