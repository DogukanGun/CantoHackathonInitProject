package lending

const (
	cifBILL             = "0x897709FC83ba7a4271d22Ed4C01278cc1Da8d6F8"
	cfBILL              = "0xF1F89dF149bc5f2b6B29783915D1F9FE2d24459c"
	cUSYC               = "0x0355E393cF0cf5486D9CAefB64407b7B1033C2f1"
	cNOTETestnet        = "0x04E52476d318CdF739C38BD41A922787D441900c"
	cNOTE               = "0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C"
	cCANTOTestnet       = "0x477eaF5DECf6299EE937954084f0d53EFc57346F"
	cCANTO              = "0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488"
	cETHTestnet         = "0x260fCD909ab9dfF97B03591F83BEd5bBfc89A571"
	cETH                = "0x830b9849e7d79b92408a86a557e7baaacbec6030"
	cUSDCTestnet        = "0x9160c5760a540cAfA24F90102cAA14C50497d5b7"
	cUSDC               = "0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e"
	cUSDTTestnet        = "0x3BEe0A8209e6F8c5c743F21e0cA99F2cb780D0D8"
	cUSDT               = "0x6b46ba92d7e94FfA658698764f5b8dfD537315A9"
	cATOMTestnet        = "0x90FCcb79Ad6f013A4bf62Ad43577eed7a8eb961B"
	cATOM               = "0x617383F201076e7cE0f6E625D1a983b3D1bd277A"
	cCantoNoteLPTestnet = "0x2fd02CDB9Be9428d4eC2Ae969e52710601E219C6"
	cCantoNoteLP        = "0x3C96dCfd875253A37acB3D2B102b6f328349b16B"
	cCantoAtomLPTestnet = "0x4777Dc2b41f1f2Bd878205A61c1eA2609749928C"
	cCantoAtomLP        = "0xC0D6574b2fe71eED8Cd305df0DA2323237322557"
	cNoteUSDTLPTestnet  = "0xBeD263484AEDFD449eE1ed8f0b4799192026E190"
	cNoteUSDTLP         = "0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F"
	cNoteUSDCLPTestnet  = "0xB2C5512a8A70835Cb9aBe830C9e61FBDdcd1dC81"
	cNoteUSDCLP         = "0xD6a97e43FC885A83E97d599796458A331E580800"
	cCantoETHLPTestnet  = "0xf301c9d5804Fab3dd207ef75f78509db6393f37F"
	cCantoETHLP         = "0xb49A395B39A0b410675406bEE7bD06330CB503E3"
)

var LendingTokens = []string{cifBILL, cfBILL, cUSYC, cNOTE, cCANTO, cETH, cUSDC, cUSDT, cATOM, cCantoNoteLP, cCantoAtomLP,
	cNoteUSDTLP, cNoteUSDCLP, cCantoETHLP}

var LendingTokensTestnet = []string{cNOTETestnet, cCANTOTestnet, cETHTestnet, cUSDCTestnet, cUSDTTestnet, cATOMTestnet,
	cCantoNoteLPTestnet, cCantoAtomLPTestnet, cNoteUSDTLPTestnet, cNoteUSDCLPTestnet, cCantoETHLPTestnet}

type LendingPools struct {
	TokenAddress string  `json:"token_address"`
	Name         string  `json:"name"`
	Apy          float64 `json:"apy"`
}
