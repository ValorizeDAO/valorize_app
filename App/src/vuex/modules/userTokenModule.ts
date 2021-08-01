import { Token } from "../../models/Token"

export default {
  namespaced: true,
  state() {
    return {
      token: {
        address: "",
        tokenName: "",
        tokenSymbol: "",
      },
      deployed: {
        mainnet: false,
        testnet: false,
      },
      circulatingSupply: 0,
      ownerAddress: "",
    }
  },
  getters: {
    tokenAddress: (state: TokenState) => state.token.address,
    token: (state: TokenState) => state.token,
    ownerAddress: (state: TokenState) => state.ownerAddress,
    deployed: (state: TokenState) => state.deployed,
  },
  mutations: {
    setToken(state: TokenState, payload: Token) {
      state.token = payload
    },
  },
}

interface TokenState {
  token: Token,
  deployed: {
    mainnet: boolean,
    testnet: boolean,
  },
  circulatingSupply: number,
  ownerAddress: string
}
