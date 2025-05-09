// Package OpenXRIPBindingModifier provides methods for working with OpenXRIPBindingModifier object instances.
package OpenXRIPBindingModifier

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/OpenXRBindingModifier"
import "graphics.gd/classdb/Resource"
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
Binding modifier that applies directly on an interaction profile.
*/
type Instance [1]gdclass.OpenXRIPBindingModifier

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRIPBindingModifier() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRIPBindingModifier

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRIPBindingModifier"))
	casted := Instance{*(*gdclass.OpenXRIPBindingModifier)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsOpenXRIPBindingModifier() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRIPBindingModifier() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsOpenXRBindingModifier() OpenXRBindingModifier.Advanced {
	return *((*OpenXRBindingModifier.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRBindingModifier() OpenXRBindingModifier.Instance {
	return *((*OpenXRBindingModifier.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(OpenXRBindingModifier.Advanced(self.AsOpenXRBindingModifier()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(OpenXRBindingModifier.Instance(self.AsOpenXRBindingModifier()), name)
	}
}
func init() {
	gdclass.Register("OpenXRIPBindingModifier", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRIPBindingModifier{*(*gdclass.OpenXRIPBindingModifier)(unsafe.Pointer(&ptr))}
	})
}
