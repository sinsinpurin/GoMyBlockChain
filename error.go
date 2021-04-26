package gomyblockchain

import "fmt"

/*
BCError BlockChainのエラーメッセージの構成
*/
type BCError struct {
	Function string
	Msg      string
}

func (err *BCError) Error() string {
	return fmt.Sprintf("error %s [%s]", err.Msg, err.Function)
}
