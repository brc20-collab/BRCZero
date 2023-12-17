package types

const (
	EventTypeBasicx         = ModuleName
	EventTypeBasicxProtocol = "basicx_protocol"
	EventTypeManageContract = "manage_contract"
	EventTypeEntryPoint     = "entry_point"
	EventTypeCallEvm        = "call_evm"

	AttributeResult   = "result"
	AttributeProtocol = "protocol"
	AttributeBTCTXID  = "btc_txid"

	AttributeManageContractOperation = "operation"
	AttributeManageContractAddress   = "contract_addrss"
	AttributeEvmOutput               = "evm_output"
	AttributeManageLog               = "log"
)
