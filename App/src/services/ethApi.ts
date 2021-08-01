export default {
  async deployTokenFromBackend({ tokenName, tokenTiker }: { tokenName: string, tokenTiker: string }): Promise<any> {
    var formdata = new FormData();
    formdata.append("tokenName", tokenName);
    formdata.append("tokenTicker", tokenTiker);

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
  status: string, 
  token: Token
}
export interface Token {
    address: string, 
    name: string, 
    tiker: string
  }
