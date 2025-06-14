package types

// FileInfo represents the parsed information about a Go file
type FileInfo struct {
	PackageName         string        // Name of the package
	Imports             []string      // List of imported packages
	Functions           []string      // List of function names
	Structs             []*StructInfo // List of struct names with their comments, fields, and methods
	UsedImportedStructs []*StructInfo // List of imported struct names used in the file, with fields and methods
}

// NewFileInfo creates a new FileInfo instance
func NewFileInfo() *FileInfo {
	return &FileInfo{
		Imports:             make([]string, 0),
		Functions:           make([]string, 0),
		Structs:             make([]*StructInfo, 0),
		UsedImportedStructs: make([]*StructInfo, 0),
	}
}

// StructField represents a field within a struct
type StructField struct {
	Name string // Field name
	Type string // Field type
}

// StructMethod represents a method associated with a struct
type StructMethod struct {
	Name        string   // Method name
	Comment     string   // Method comment
	Parameters  []string // List of parameter types
	ReturnTypes []string // List of return types
}

// StructInfo represents detailed information about a struct
type StructInfo struct {
	Name    string          // Struct name
	Comment string          // Struct comment
	Fields  []*StructField  // List of fields
	Methods []*StructMethod // List of methods
}

// NewStructInfo creates a new StructInfo instance
func NewStructInfo() *StructInfo {
	return &StructInfo{
		Fields:  make([]*StructField, 0),
		Methods: make([]*StructMethod, 0),
	}
}

// Node represents a package in the dependency graph
type Node struct {
	PkgPath   string   // Package path
	Functions []string // Exported functions
	DependsOn []string // Imported packages
	Files     []string // Source files in the package
}

// DependencyGraph represents the project's dependency structure
type DependencyGraph struct {
	Nodes map[string]*Node // Key: package path
}

// NewDependencyGraph creates a new DependencyGraph instance
func NewDependencyGraph() *DependencyGraph {
	return &DependencyGraph{
		Nodes: make(map[string]*Node),
	}
}
