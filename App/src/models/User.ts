export interface User {
  id: number,
  email: string,
  name: string,
  username: string,
  avatar: string,
  about: string
}

export const emptyUser: User = {
  id: 0,
  email: "",
  name: "",
  username: "",
  avatar: "",
  about: "",
}
