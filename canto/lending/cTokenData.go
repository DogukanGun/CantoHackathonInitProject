package lending

const (
	cNOTE        = "0x04E52476d318CdF739C38BD41A922787D441900c"
	cCANTO       = "0x477eaF5DECf6299EE937954084f0d53EFc57346F"
	cETH         = "0x260fCD909ab9dfF97B03591F83BEd5bBfc89A571"
	cUSDC        = "0x9160c5760a540cAfA24F90102cAA14C50497d5b7"
	cUSDT        = "0x3BEe0A8209e6F8c5c743F21e0cA99F2cb780D0D8"
	cATOM        = "0x90FCcb79Ad6f013A4bf62Ad43577eed7a8eb961B"
	cCantoNoteLP = "0x2fd02CDB9Be9428d4eC2Ae969e52710601E219C6"
	cCantoAtomLP = "0x4777Dc2b41f1f2Bd878205A61c1eA2609749928C"
	cNoteUSDTLP  = "0xBeD263484AEDFD449eE1ed8f0b4799192026E190"
	cNoteUSDCLP  = "0xB2C5512a8A70835Cb9aBe830C9e61FBDdcd1dC81"
	cCantoETHLP  = "0xf301c9d5804Fab3dd207ef75f78509db6393f37F"
)

var LendingTokens = []string{cNOTE, cCANTO, cETH, cUSDC, cUSDT, cATOM, cCantoNoteLP, cCantoAtomLP, cNoteUSDTLP,
	cNoteUSDCLP, cCantoETHLP}

type LendingPools struct {
	TokenAddress string  `json:"token_address"`
	Name         string  `json:"name"`
	Apy          float64 `json:"apy"`
}
