// Package ScriptEditorBase provides methods for working with ScriptEditorBase object instances.
package ScriptEditorBase

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/BoxContainer"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Container"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/VBoxContainer"
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
Base editor for editing scripts in the [ScriptEditor]. This does not include documentation items.
*/
type Instance [1]gdclass.ScriptEditorBase

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsScriptEditorBase() Instance
}

/*
Returns the underlying [Control] used for editing scripts. For text scripts, this is a [CodeEdit].
*/
func (self Instance) GetBaseEditor() [1]gdclass.Control { //gd:ScriptEditorBase.get_base_editor
	return [1]gdclass.Control(Advanced(self).GetBaseEditor())
}

/*
Adds a [EditorSyntaxHighlighter] to the open script.
*/
func (self Instance) AddSyntaxHighlighter(highlighter [1]gdclass.EditorSyntaxHighlighter) { //gd:ScriptEditorBase.add_syntax_highlighter
	Advanced(self).AddSyntaxHighlighter(highlighter)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ScriptEditorBase

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ScriptEditorBase"))
	casted := Instance{*(*gdclass.ScriptEditorBase)(unsafe.Pointer(&object))}
	return casted
}

/*
Returns the underlying [Control] used for editing scripts. For text scripts, this is a [CodeEdit].
*/
//go:nosplit
func (self class) GetBaseEditor() [1]gdclass.Control { //gd:ScriptEditorBase.get_base_editor
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditorBase.Bind_get_base_editor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Control{gd.PointerMustAssertInstanceID[gdclass.Control](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Adds a [EditorSyntaxHighlighter] to the open script.
*/
//go:nosplit
func (self class) AddSyntaxHighlighter(highlighter [1]gdclass.EditorSyntaxHighlighter) { //gd:ScriptEditorBase.add_syntax_highlighter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(highlighter[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditorBase.Bind_add_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnNameChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("name_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnEditedScriptChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("edited_script_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnRequestHelp(cb func(topic string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("request_help"), gd.NewCallable(cb), 0)
}

func (self Instance) OnRequestOpenScriptAtLine(cb func(script Object.Instance, line int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("request_open_script_at_line"), gd.NewCallable(cb), 0)
}

func (self Instance) OnRequestSaveHistory(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("request_save_history"), gd.NewCallable(cb), 0)
}

func (self Instance) OnRequestSavePreviousState(cb func(state map[any]any)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("request_save_previous_state"), gd.NewCallable(cb), 0)
}

func (self Instance) OnGoToHelp(cb func(what string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("go_to_help"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSearchInFilesRequested(cb func(text string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("search_in_files_requested"), gd.NewCallable(cb), 0)
}

func (self Instance) OnReplaceInFilesRequested(cb func(text string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("replace_in_files_requested"), gd.NewCallable(cb), 0)
}

func (self Instance) OnGoToMethod(cb func(script Object.Instance, method string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("go_to_method"), gd.NewCallable(cb), 0)
}

func (self class) AsScriptEditorBase() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScriptEditorBase() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsVBoxContainer() VBoxContainer.Advanced {
	return *((*VBoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVBoxContainer() VBoxContainer.Instance {
	return *((*VBoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsBoxContainer() BoxContainer.Advanced {
	return *((*BoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBoxContainer() BoxContainer.Instance {
	return *((*BoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VBoxContainer.Advanced(self.AsVBoxContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VBoxContainer.Instance(self.AsVBoxContainer()), name)
	}
}
func init() {
	gdclass.Register("ScriptEditorBase", func(ptr gd.Object) any {
		return [1]gdclass.ScriptEditorBase{*(*gdclass.ScriptEditorBase)(unsafe.Pointer(&ptr))}
	})
}
