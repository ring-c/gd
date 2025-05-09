package main

import (
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"reflect"
	"slices"
	"strings"

	"graphics.gd/internal/gdjson"
	"graphics.gd/internal/gdtype"
	"runtime.link/api/xray"
)

func main() {
	if err := generate(); err != nil {
		log.Fatal(err)
	}
}

func generate() error {
	spec, err := LoadSpecification()
	if err != nil {
		return xray.New(err)
	}
	if err := os.MkdirAll("./classdb", 0755); err != nil {
		return xray.New(err)
	}
	var classDB = make(ClassDB)
	for _, enum := range spec.GlobalEnums {
		classDB[strings.Replace(enum.Name, ".", "", -1)] = gdjson.Class{
			IsEnum:  true,
			Name:    strings.Replace(enum.Name, ".", "", -1),
			Package: "internal",
		}
	}
	for _, class := range spec.BuiltinClasses {
		for _, enum := range class.Enums {
			classDB[class.Name+strings.Replace(enum.Name, ".", "", -1)] = gdjson.Class{
				IsEnum:  true,
				Name:    class.Name + strings.Replace(enum.Name, ".", "", -1),
				Package: "internal",
			}
		}
	}
	var global_enums = make(map[string]gdjson.Enum)
	for _, enum := range spec.GlobalEnums {
		global_enums[enum.Name] = enum
	}
	for i, class := range spec.Classes {
		var pkg = "internal"
		if !gdtype.Name(class.Name).InCore() {
			pkg = "classdb"
		}
		class.Package = pkg
		spec.Classes[i] = class
		classDB[class.Name] = class
	}
	var singletons = make(map[string]bool)
	for _, class := range spec.Singletons {
		singletons[class.Name] = true
		mod := classDB[class.Name]
		mod.IsSingleton = true
		classDB[class.Name] = mod
	}
	gdtype.ClassDB = classDB
	file, err := os.Create("./classdb/objects.go")
	if err != nil {
		return xray.New(err)
	}
	defer file.Close()
	fmt.Fprintln(file, `package classdb`)
	fmt.Fprintln(file)
	fmt.Fprintln(file, `import "graphics.gd/internal/gdclass"`)
	fmt.Fprintln(file)
	for _, class := range spec.Classes {
		if gdtype.Name(class.Name).InCore() {
			continue
		}
		fmt.Fprintf(file, "type %s = [1]gdclass.%[1]s\n", class.Name)
		if err := classDB.generateObjectPackage(class, singletons[class.Name], global_enums); err != nil {
			return xray.New(err)
		}
	}
	return nil
}

func generateEnum(code io.Writer, prefix string, enum gdjson.Enum, classdb string) {
	rename := enum.Name
	if enum.Name == "MouseMode" {
		rename = "MouseModeValue"
	}
	original := enum.Name
	if prefix != "" {
		original = prefix + "." + original
	}
	rename = strings.Replace(rename, ".", "", -1)
	enum.Name = strings.Replace(enum.Name, ".", "", -1)

	if enum.Name == "Error" {
		return
	} else {
		if classdb != "" {
			fmt.Fprintf(code, "type %v = %s%s%s//gd:%s\n\n", rename, classdb, prefix, enum.Name, original)
		} else {
			fmt.Fprintf(code, "type %v int\n\n", rename)
		}
	}
	fmt.Fprintf(code, "const (\n")
	for _, value := range enum.Values {
		n := convertName(value.Name)
		if n == enum.Name {
			n += "Default"
		}
		if value.Description != "" {
			fmt.Fprint(code, "/*")
			fmt.Fprint(code, value.Description)
			fmt.Fprintln(code, "*/")
		}
		fmt.Fprintf(code, "\t%v %v = %v\n", n, rename, value.Value)
	}
	fmt.Fprintf(code, ")\n")
}

func (classDB ClassDB) generateObjectPackage(class gdjson.Class, singleton bool, global_enums map[string]gdjson.Enum) error {
	if err := os.MkdirAll("./classdb/"+class.Name, 0755); err != nil {
		return xray.New(err)
	}
	file, err := os.Create("./classdb/" + class.Name + "/class.go")
	if err != nil {
		return xray.New(err)
	}

	defer file.Close()
	fmt.Fprintf(file, `// Package %s provides methods for working with %[1]s object instances.`, class.Name)
	fmt.Fprintf(file, "\npackage %s\n\n", class.Name)
	fmt.Fprintln(file, `import "unsafe"`)
	if singleton {
		fmt.Fprintln(file, `import "sync"`)
	}
	fmt.Fprintln(file, `import "reflect"`)
	fmt.Fprintln(file, `import "slices"`)
	fmt.Fprintln(file, `import "graphics.gd/internal/pointers"`)
	fmt.Fprintln(file, `import "graphics.gd/internal/callframe"`)
	fmt.Fprintln(file, `import gd "graphics.gd/internal"`)
	fmt.Fprintln(file, `import "graphics.gd/internal/gdclass"`)
	fmt.Fprintln(file, `import "graphics.gd/variant"`)
	if class.Inherits != "" {
		super := classDB[class.Inherits]
		for super.Name != "" && super.Name != "Object" && super.Name != "RefCounted" && !classDB[super.Name].IsSingleton {
			super = classDB[super.Inherits]
		}
	}
	for pkg := range gdtype.ImportsForClass(class) {
		fmt.Fprintf(file, "import %q\n", pkg)
	}
	fmt.Fprintln(file)
	fmt.Fprintln(file, "var _ Object.ID")
	fmt.Fprintln(file, "var _ RefCounted.Instance")
	fmt.Fprintln(file, "var _ unsafe.Pointer")
	fmt.Fprintln(file, "var _ reflect.Type")
	fmt.Fprintln(file, "var _ callframe.Frame")
	fmt.Fprintln(file, "var _ = pointers.Cycle")
	fmt.Fprintln(file, "var _ = Array.Nil")
	fmt.Fprintln(file, "var _ variant.Any")
	fmt.Fprintln(file, "var _ Callable.Function")
	fmt.Fprintln(file, "var _ Dictionary.Any")
	fmt.Fprintln(file, "var _ RID.Any")
	fmt.Fprintln(file, "var _ String.Readable")
	fmt.Fprintln(file, "var _ Path.ToNode")
	fmt.Fprintln(file, "var _ Packed.Bytes")
	fmt.Fprintln(file, "var _ Error.Code")
	fmt.Fprintln(file, "var _ Float.X")
	fmt.Fprintln(file, "var _ = slices.Delete[[]struct{}, struct{}]")
	fmt.Fprintln(file)
	var local_enums = make(map[string]bool)
	var hasVirtual bool
	for _, method := range class.Methods {
		if method.IsVirtual {
			hasVirtual = true
			break
		}
	}
	if class.Description != "" {
		fmt.Fprintln(file, "/*")
		fmt.Fprint(file, strings.Replace(class.Description, "*/", "", -1))
		fmt.Fprintln(file)
		if hasVirtual {
			fmt.Fprintf(file, "\t See [Interface] for methods that can be overridden by a [Class] that extends it.\n", class.Name)
		}
		fmt.Fprintln(file, "\n*/")
	}
	if singleton {
		fmt.Fprintf(file, "var self [1]gdclass.%s\n", class.Name)
		fmt.Fprintf(file, "var once sync.Once\n")
		fmt.Fprintf(file, "func singleton() {\n")
		fmt.Fprintf(file, "\tobj := gd.Global.Object.GetSingleton(gd.Global.Singletons.%s)\n", class.Name)
		fmt.Fprintf(file, "\tself = *(*[1]gdclass.%s)(unsafe.Pointer(&obj))\n", class.Name)
		fmt.Fprintf(file, "}\n")
	} else {
		fmt.Fprintf(file, "type Instance [1]gdclass.%s\n", class.Name)
		var hasDefaults bool
		for _, method := range class.Methods {
			for _, argument := range method.Arguments {
				if argument.DefaultValue != nil && !singleton && !method.IsStatic {
					hasDefaults = true
				}
			}
		}
		if hasDefaults {
			fmt.Fprintf(file, "type Expanded [1]gdclass.%s\n", class.Name)
		}
		fmt.Fprintf(file, "// Nil is a nil/null instance of the class. Equivalent to the zero value.\n")
		fmt.Fprintf(file, "var Nil Instance\n")
		fmt.Fprintf(file, "type Any interface {\n")
		fmt.Fprintf(file, "\tgd.IsClass\n")
		fmt.Fprintf(file, "\tAs%s() Instance\n", class.Name)
		fmt.Fprintf(file, "}\n")
		if hasVirtual {
			fmt.Fprintf(file, "type Interface interface {\n")
			for _, method := range class.Methods {
				if method.IsVirtual {
					if method.Description != "" {
						description := strings.Replace(method.Description, "*/", "", -1)
						description = strings.TrimSpace(description)
						description = strings.Replace(description, "\n", "\n\t\t//", -1)
						fmt.Fprintln(file, "\t\t//"+description)
					}
					fmt.Fprintf(file, "\t%s(", convertName(method.Name))
					for i, arg := range method.Arguments {
						if i > 0 {
							fmt.Fprint(file, ", ")
						}
						fmt.Fprint(file, fixReserved(arg.Name), " ", classDB.convertTypeSimple(class, class.Name+"."+method.Name+"."+arg.Name, arg.Meta, arg.Type))
					}
					fmt.Fprint(file, ") ", classDB.convertTypeSimple(class, "", method.ReturnValue.Meta, method.ReturnValue.Type))
					fmt.Fprintln(file)
				}
			}
			fmt.Fprintf(file, "}\n")

			fmt.Fprintf(file, "// Implementation implements [Interface] with empty methods.\n")
			fmt.Fprintf(file, "type Implementation = implementation\n\n")
			fmt.Fprintf(file, "type implementation struct{}\n")

			for _, method := range class.Methods {
				if method.IsVirtual {
					fmt.Fprintf(file, "func (self implementation) %[1]v(", convertName(method.Name))
					for i, arg := range method.Arguments {
						if i > 0 {
							fmt.Fprint(file, ", ")
						}
						fmt.Fprint(file, fixReserved(arg.Name), " ", classDB.convertTypeSimple(class, class.Name+"."+method.Name+"."+arg.Name, arg.Meta, arg.Type))
					}
					fmt.Fprint(file, ")")
					if method.ReturnValue.Type != "" {
						fmt.Fprintf(file, "(_ %s)", classDB.convertTypeSimple(class, "", method.ReturnValue.Meta, method.ReturnValue.Type))
					}
					fmt.Fprintln(file, " { return }")
				}
			}
		}
	}
	var getter_setters = make(map[string]bool)
	for _, property := range class.Properties {
		if property.Getter != "" {
			getter_setters[property.Getter] = true
		}
		if property.Setter != "" {
			getter_setters[property.Setter] = true
		}
	}
	for _, method := range class.Methods {
		if getter_setters[method.Name] {
			continue
		}
		classDB.simpleCall(file, class, method, singleton, true)
		var hasDefault bool
		for _, argument := range method.Arguments {
			if argument.DefaultValue != nil {
				hasDefault = true
			}
		}
		if hasDefault {
			classDB.simpleCall(file, class, method, singleton, false)
		}
	}
	fmt.Fprintf(file, `// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.`)
	if singleton {
		fmt.Fprintf(file, "\nfunc Advanced() class { once.Do(singleton); return self }\n")
	} else {
		fmt.Fprintf(file, "\ntype Advanced = class\n")
	}
	fmt.Fprintf(file, "type class [1]gdclass.%s\n", class.Name)
	fmt.Fprintln(file, "func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }")
	fmt.Fprintf(file, "\n\n//go:nosplit\nfunc (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }\n")
	if !singleton {
		fmt.Fprintln(file, "func (self Instance) AsObject() [1]gd.Object { return self[0].AsObject() }")
		fmt.Fprintf(file, "\n\n//go:nosplit\nfunc (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }\n")
	}
	if !singleton {
		classDB.new(file, class)
	}
	classDB.properties(file, class, singleton)
	for _, method := range class.Methods {
		classDB.methodCall(file, class.Package, class, method, callDefault)
	}
	for _, signal := range class.Signals {
		classDB.signalCall(file, class, signal, singleton)
	}
	super := classDB[class.Inherits]
	if class.Inherits != "" {
		var i = 1
		if !singleton {
			fmt.Fprintf(file, "\nfunc (self class) As%[1]v() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }\n", class.Name)
			fmt.Fprintf(file, "func (self Instance) As%[1]v() Instance { return *((*Instance)(unsafe.Pointer(&self))) }\n", class.Name)
		}
		super := classDB[class.Inherits]
		for super.Name != "" && super.Name != "Object" {
			if classDB[super.Name].IsSingleton {
				super = classDB[super.Inherits]
				continue
			}
			if super.Name == "RefCounted" {
				fmt.Fprintf(file, "func (self class) AsRefCounted() [1]gd.RefCounted { return *((*[1]gd.RefCounted)(unsafe.Pointer(&self))) }\n")
				if !singleton {
					fmt.Fprintf(file, "func (self Instance) AsRefCounted() [1]gd.RefCounted { return *((*[1]gd.RefCounted)(unsafe.Pointer(&self))) }\n")
				}
			} else {
				fmt.Fprintf(file, "func (self class) As%[2]v() %[2]v.Advanced { return *((*%[2]v.Advanced)(unsafe.Pointer(&self))) }\n", class.Name, super.Name)
				if !singleton {
					fmt.Fprintf(file, "func (self Instance) As%[2]v() %[2]v.Instance { return *((*%[2]v.Instance)(unsafe.Pointer(&self))) }\n", class.Name, super.Name)
				}
			}
			i++
			super = classDB[super.Inherits]
		}
	}
	for _, self := range []string{"class", "Instance"} {
		if self == "Instance" && singleton {
			continue
		}
		fmt.Fprintf(file, "\nfunc (self %s) Virtual(name string) reflect.Value {\n", self)
		fmt.Fprintf(file, "\tswitch name {\n")
		for _, method := range class.Methods {
			if method.IsVirtual {
				fmt.Fprintf(file, "\tcase \"%v\": return reflect.ValueOf(self.%v);\n", method.Name, method.Name)
			}
		}
		if class.Inherits != "" && !classDB[class.Inherits].IsSingleton {
			var name = "Instance"
			if self != "Instance" {
				name = "Advanced"
			}
			fmt.Fprintf(file, "\tdefault: return gd.VirtualByName(%s.%s(self.As%[1]s()), name)\n", super.Name, name)
		} else {
			fmt.Fprintf(file, "\tdefault: return reflect.Value{}\n")
		}
		fmt.Fprintf(file, "\t}\n")
		fmt.Fprintf(file, "}\n")
	}
	fmt.Fprintf(file, `func init() {`)
	fmt.Fprintf(file, `gdclass.Register("%s", func(ptr gd.Object) any { return [1]gdclass.%[1]s{*(*gdclass.%[1]v)(unsafe.Pointer(&ptr))} })`, class.Name)
	fmt.Fprintf(file, "}\n")
	if class.Name == "DisplayServer" {
		local_enums["MouseButton"] = true
	}
	for _, method := range class.Methods {
		for _, argument := range method.Arguments {
			name := strings.TrimPrefix(argument.Type, "enum::")
			name = strings.TrimPrefix(name, "bitfield::")
			if _, ok := global_enums[name]; ok {
				local_enums[name] = true
			}
		}
		name := strings.TrimPrefix(method.ReturnValue.Type, "enum::")
		name = strings.TrimPrefix(name, "bitfield::")
		if _, ok := global_enums[name]; ok {
			local_enums[name] = true
		}
	}
	for _, enum := range class.Enums {
		generateEnum(file, class.Name, enum, "gdclass.")
	}
	for _, key := range slices.Sorted(maps.Keys(local_enums)) {
		enum := global_enums[key]
		generateEnum(file, "", enum, "")
	}
	generateStructables(file)
	return nil
}

func registerStructables(rtype reflect.Type) {
	if rtype.PkgPath() == "graphics.gd/internal/gdjson" {
		StructablesInThisPackageGlobalHack[rtype] = true
	}
	switch rtype.Kind() {
	case reflect.Array, reflect.Slice:
		registerStructables(rtype.Elem())
	case reflect.Map:
		registerStructables(rtype.Key())
		registerStructables(rtype.Elem())
	case reflect.Struct:
		for i := 0; i < rtype.NumField(); i++ {
			registerStructables(rtype.Field(i).Type)
		}
	}
}

func generateStructables(file io.Writer) {
	var sortable = make([]string, 0, len(StructablesInThisPackageGlobalHack))
	for rtype := range StructablesInThisPackageGlobalHack {
		switch rtype.Kind() {
		case reflect.Map:
			sortable = append(sortable,
				fmt.Sprintf("type %v map[%v]%v\n", rtype.Name(), rtype.Key().String(), rtype.Elem().String()))
		case reflect.Struct:
			var w strings.Builder
			fmt.Fprintf(&w, "type %v struct {\n", rtype.Name())
			for i := range rtype.NumField() {
				field := rtype.Field(i)
				if field.PkgPath != "" {
					continue
				}
				ftype := field.Type.String()
				if override, ok := field.Tag.Lookup("type"); ok {
					ftype = override
				}
				ftype = strings.Replace(ftype, "gdjson.", "", -1)
				gdTag := field.Tag.Get("gd")
				fmt.Fprintf(&w, "%v %v `gd:\"%s\"`\n", field.Name, ftype, gdTag)
			}
			fmt.Fprintf(&w, "}\n")
			sortable = append(sortable, w.String())
		}
	}
	slices.Sort(sortable)
	for _, s := range sortable {
		fmt.Fprint(file, s)
	}
	clear(StructablesInThisPackageGlobalHack)
}
