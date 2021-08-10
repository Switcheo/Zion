package utils

import (
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// MethodID only used for register method handler and prepare native contract context ref.
func MethodID(ab abi.ABI, name string) string {
	m, ok := ab.Methods[name]
	if !ok {
		panic(fmt.Sprintf("method name %s not exist", name))
	}
	return hexutil.Encode(m.ID)
}

func PackMethod(ab abi.ABI, name string, args ...interface{}) ([]byte, error) {
	method, exist := ab.Methods[name]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", name)
	}
	arguments, err := method.Inputs.Pack(args...)
	if err != nil {
		return nil, err
	}
	return append(method.ID, arguments...), nil
}

func UnpackMethod(ab abi.ABI, name string, data interface{}, payload []byte) error {
	mth, ok := ab.Methods[name]
	if !ok {
		return fmt.Errorf("abi method %s not exist", name)
	}

	if len(payload) < 4 || len(payload[4:])%32 != 0 {
		return fmt.Errorf("invalid payload")
	}

	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		return fmt.Errorf("method input should be pointer")
	}

	args := mth.Inputs
	unpacked, err := args.Unpack(payload[4:])
	if err != nil {
		return err
	}
	return args.Copy(data, unpacked)
}

func PackOutputs(ab abi.ABI, method string, args ...interface{}) ([]byte, error) {
	mth, exist := ab.Methods[method]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", method)
	}
	return mth.Outputs.Pack(args...)
}

func UnpackOutputs(ab abi.ABI, name string, data interface{}, payload []byte) error {
	mth, ok := ab.Methods[name]
	if !ok {
		return fmt.Errorf("abi method %s not exist", name)
	}

	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		return fmt.Errorf("method output should be pointer")
	}

	args := mth.Outputs
	unpacked, err := args.Unpack(payload)
	if err != nil {
		return err
	}
	return args.Copy(data, unpacked)
}