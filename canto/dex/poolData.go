package cantoDex

type Pool struct {
	PoolAddress string `json:"pool_address"`
	Token0      Token  `json:"token_0"`
	Token1      Token  `json:"token_1"`
}

type Token struct {
	Symbol            string `json:"symbol"`
	Address           string `json:"address"`
	PriceAsOtherToken int64  `json:"price_as_other_token"`
}

const (
	CantoNoteLPTestnet = "0x395E6ce7891f32278375Ff551B8ed61dF5579fE3"
	CantoNoteLP        = "0x1D20635535307208919f0b67c3B2065965A85aA9"
	CantoAtomLPTestnet = "0x2bDF6c1302efc3c03D9C95f6fb5a4826A6bD964b"
	CantoAtomLP        = "0x30838619C55B787BafC3A4cD9aEa851C1cfB7b19"
	NoteUSDTLPTestnet  = "0x252631e22e1ECc2fc0E811562605ed624B7E31d5"
	NoteUSDTLP         = "0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833"
	NoteUSDCLPTestnet  = "0x2db30A39Ec88247da8906506DB8E9dd933A5C775"
	NoteUSDCLP         = "0x9571997a66D63958e1B3De9647C22bD6b9e7228c"
	CantoETHLPTestnet  = "0x905D3d7F4C892d535160f1E2BA55f23Cd306718b"
	CantoETHLP         = "0x216400ba362d8FCE640085755e47075109718C8B"
	NOTE               = "0x03F734Bd9847575fDbE9bEaDDf9C166F880B5E5f"
	WCANTO             = "0x04a72466De69109889Db059Cb1A4460Ca0648d9D"
	ETH                = "0xCa03230E7FB13456326a234443aAd111AC96410A"
	USDC               = "0xc51534568489f47949A828C8e3BF68463bdF3566"
	USDT               = "0x4fC30060226c45D8948718C95a78dFB237e88b40"
	ATOM               = "0x40E41DC5845619E7Ba73957449b31DFbfB9678b2"
)

var DexPools = []string{CantoNoteLP, CantoAtomLP, NoteUSDCLP, NoteUSDTLP, CantoETHLP}
var DexPoolsTestnet = []string{CantoNoteLPTestnet, CantoAtomLPTestnet, NoteUSDCLPTestnet, NoteUSDTLPTestnet, CantoETHLPTestnet}
