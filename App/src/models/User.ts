export interface User {
  id: string,
  email: string,
  name: string,
  username: string
}

export const emptyUser: User = {
  id: "",
  email: "",
  name: "",
  username: "",
}
