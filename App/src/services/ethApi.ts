import { Token } from "../models/Token"

export default {
  async deployTokenToTestNet({ tokenName, tokenSymbol }: { tokenName: string, tokenSymbol: string }): Promise<any> {
    var formdata = new FormData();
    formdata.append("tokenName", tokenName);
    formdata.append("tokenTicker", tokenSymbol);

    var requestOptions = {
      method: 'POST',
      body: formdata,
      redirect: 'follow',
      credentials: 'include',
    } as RequestInit;

    const apiResponse = await fetch(import.meta.env.VITE_BACKEND_URL + "/api/v0/admin/deploy", requestOptions) as Response;

    return apiResponse;
  }
}

export interface TokenResponse {
  address: string;
  contract_version: string;
  name: string;
  network: string;
  symbol: string;
  tx: string;
  user_id: string;
}