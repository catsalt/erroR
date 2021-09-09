// erroR.go
package erroR

import (
	"fmt"
)

// ___ERR FYI ⬆erroR.Newerr.go⬆___

// 出现上述错误提示，请检查，参数 fnID 命名规则：
// 文件内 文件名小写+函数序号大写+错误序号："fB1" |-> fB_1:
// 文件间 包名小写 +小写+大写+错误序号："pfB1" |-> p_fB_1:
// 包间 包名大写 +小写+大写+错误序号："PfB1" |-> P_fB_1:
func nCfmt(fnID string) string {
	if len(fnID) < 3 {
		fnID = " ### ___ERR FYI ⬆erroR.Newerr.go⬆___ ###"
	}
	if fnID[2] > 47 && fnID[2] < 58 {
		return fnID[0:2] + "_" + fnID[2:] + ":"
	}
	return fnID[0:1] + "_" + fnID[1:3] + "_" + fnID[3:] + ":"
}

// 文件内"fB1" |fB_1: ; 包间"PfB1" |P_fB_1: ; 文件间"pfB1" |p_fB_1:
// nil <- errNew == ""
func RnKnew(fnID, errNew string) error {
	if errNew == "" {
		return nil
	}
	return fmt.Errorf("> %s %s <", nCfmt(fnID), errNew)
}

// 文件内"fB1" |fB_1: >>; 包间"PfB1" |P_fB_1: >>; 文件间"pfB1" |p_fB_1: >>
// nil <- wrapErr == nil
func RnIwrap(fnID string, wrapErr error) error {
	if wrapErr == nil {
		return nil
	}
	return fmt.Errorf("> %s %w<", nCfmt(fnID), wrapErr)
}

// nil <- errNew == "" && WrapErr == nil
func RnNwrapORnew(fnID string, wrapErr error, errNew string) error {
	if wrapErr == nil {
		return RnKnew(fnID, errNew)
	}
	return RnIwrap(fnID, wrapErr)

}

// 收敛累加Err; 文件内"fB1" |fB_1: >>; 包间"PfB1" |P_fB_1: >>; 文件间"pfB1" |p_fB_1: >>
// nil <- errlist 都是 nil
func RnGwrapErrList(fnID string, errlist ...error) error {
	if len(errlist) == 0 {
		return nil
	}
	var s string
	for _, e := range errlist {
		if e == nil {
			continue
		}
		s += "( " + e.Error() + " )"
	}
	if len(s) == 0 {
		return nil
	}
	return fmt.Errorf("> %s %v<", nCfmt(fnID), s)
}
