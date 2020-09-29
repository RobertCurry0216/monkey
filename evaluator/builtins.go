package evaluator

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"monkey/object"
	"os"
	"strconv"
	"strings"
	"time"
)

var builtins = map[string]*object.Builtin{
	"len":     &object.Builtin{Fn: lenBuiltin},
	"first":   &object.Builtin{Fn: firstBuiltin},
	"last":    &object.Builtin{Fn: lastBuiltin},
	"rest":    &object.Builtin{Fn: restBuiltin},
	"push":    &object.Builtin{Fn: pushBuiltin},
	"pop":     &object.Builtin{Fn: popBuiltin},
	"replace": &object.Builtin{Fn: replaceBuiltin},
	"bool":    &object.Builtin{Fn: boolBuiltin},
	"puts":    &object.Builtin{Fn: putsBuiltin},
	"gets":    &object.Builtin{Fn: getsBuiltin},
	"geti":    &object.Builtin{Fn: getiBuiltin},
	"random":  &object.Builtin{Fn: randomBuiltin},
}

func lenBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got %s", arg.Type())
	}
}

func firstBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ArrayObj {
		return newError("argument to `first` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return nullObj
}

func lastBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ArrayObj {
		return newError("argument to `last` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return nullObj
}

func restBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ArrayObj {
		return newError("argument to `rest` must be ARRAY, got %s",
			args[0].Type())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	}

	return nullObj

}

func pushBuiltin(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2",
			len(args))
	}
	if args[0].Type() != object.ArrayObj {
		return newError("argument to `push` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	arr.Elements = append(arr.Elements, args[1])

	return arr
}

func popBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.ArrayObj {
		return newError("argument to `push` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	if length == 0 {
		return nullObj
	}

	output := arr.Elements[length-1]

	newElements := make([]object.Object, length-1, length-1)
	copy(newElements, arr.Elements)

	arr.Elements = newElements

	return output
}

func replaceBuiltin(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("wrong number of arguments. got=%d, want=3",
			len(args))
	}

	switch args[0].(type) {
	case *object.Array:
		return replaceArray(args...)
	default:
		return newError("argument to `push` must be ARRAY or HASH, got %s", args[0].Type())
	}
}

func replaceArray(args ...object.Object) object.Object {
	if args[1].Type() != object.IntegerObj {
		return newError("argument to `replace` must be INTEGER, got %s",
			args[1].Type())
	}
	array := args[0].(*object.Array)
	i := args[1].(*object.Integer).Value
	max := int64(len(array.Elements) - 1)

	if i < 0 || i > max {
		return newError("invalid index for given array, got=%d, array length=%d", i, len(array.Elements))
	}

	array.Elements[i] = args[2]

	return array

}

func boolBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if isTruthy(args[0]) {
		return trueObj
	}
	return falseObj
}

func putsBuiltin(args ...object.Object) object.Object {
	var out bytes.Buffer
	for _, arg := range args {
		out.WriteString(arg.Inspect())
		out.WriteString(" ")
	}
	fmt.Println(out.String())
	return nullObj
}

func getsBuiltin(args ...object.Object) object.Object {
	if len(args) == 1 {
		fmt.Printf(args[0].Inspect())
	}

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return &object.String{Value: text}
}

func getiBuiltin(args ...object.Object) object.Object {
	if len(args) == 1 {
		fmt.Printf(args[0].Inspect())
	}

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	value, err := strconv.Atoi(text)

	if err != nil {
		return nullObj
	}

	return &object.Integer{Value: int64(value)}
}

func randomBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.IntegerObj {
		return newError("argument to `random` must be INTEGER, got %s", args[0].Type())
	}

	cap := args[0].(*object.Integer).Value

	if cap < 1 {
		return newError("cap value must be at least 1, got %s", args[0].Type())
	}

	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(int(cap))

	return &object.Integer{Value: int64(value)}
}
