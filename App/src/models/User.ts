/* eslint-disable camelcase */
import { Token } from "./Token"
export interface User {
  id: number,
  email: string,
  name: string,
  username: string,
  avatar: string,
  about: string,
  has_deployed_token: boolean,
  has_verified_email: boolean,
  token?: Token
}

export const emptyUser: User = {
  id: 0,
  email: "",
  name: "",
  username: "",
  avatar: "",
  about: "",
  has_deployed_token: false,
  has_verified_email: false,
}
