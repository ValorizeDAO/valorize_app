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
      deployedNetworkName: "",
      circulatingSupply: 0,
      ownerAddress: "",
    }
  },
  getters: {
    tokenAddress: (state: TokenState) => state.token.address,
    token: (state: TokenState) => state.token,
    ownerAddress: (state: TokenState) => state.ownerAddress,
    deployed: (state: TokenState) => state.deployedNetworkName,
  },
  mutations: {
    setToken(state: TokenState, payload: Token) {
      state.token = payload
    },
  },
}

interface TokenState {
  token: Token,
  deployedNetworkName: string,
  circulatingSupply: number,
  ownerAddress: string
}
