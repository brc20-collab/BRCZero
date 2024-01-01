package types

const (
	EventTypeBasicX         = ModuleName
	EventTypeManageContract = "manage_contract"
	EventTypeEntryPoint     = "entry_point"
	EventTypeCallEvm        = "call_evm"

	AttributeResult       = "result"
	AttributeProtocolName = "protocol"
	AttributeBTCTXID      = "btc_txid"
	AttributeBTCBlockHash = "btc_block_hash"

	AttributeManageContractOperation = "operation"
	AttributeManageContractAddress   = "contract_addrss"
	AttributeEvmOutput               = "evm_output"
	AttributeManageLog               = "log"
)
