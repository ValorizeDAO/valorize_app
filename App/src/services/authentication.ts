import { User, emptyUser } from "../models/user"

export default {
  async isLoggedIn(): Promise<{ isLoggedIn: boolean, user: User }> {
    const response = { isLoggedIn: false, user: emptyUser }
    const requestOptions = {
      method: "GET",
      credentials: "include",
    } as RequestInit
    const apiResponse = await fetch(import.meta.env.VITE_BACKEND_URL + "/api/v0/me", requestOptions)

    if (apiResponse.status === 200) {
      const user = await apiResponse.json()
      response.isLoggedIn = true
      response.user = user
    }

    return response
  },
  async logout(): Promise<void> {
    const requestOptions = {
      method: "GET",
      credentials: "include",
    } as RequestInit
    const apiResponse = await fetch(import.meta.env.VITE_BACKEND_URL + "/logout", requestOptions)
    if (apiResponse.status === 200) {
      return
    }
    throw new Error("Logout failed")
  },
  async uploadPicture(file: File): Promise<Response>{
    var formdata = new FormData();
    formdata.append("picture", file, "user_picture.png");
    var requestOptions = {
      method: 'PUT',
      body: formdata,
      redirect: 'follow',
      credentials: 'include',
    } as RequestInit;

    const apiResponse = await fetch(import.meta.env.VITE_BACKEND_URL + "/api/v0/me/picture", requestOptions) as Response;
    return apiResponse;
  },
  async updateProfile({ name, about }: { name: string, about: string }): Promise<Response>{
    var formdata = new FormData();
    console.log({ name, about });
    
    formdata.append("name", name);
    formdata.append("about", about);

    var requestOptions = {
      method: 'PUT',
      body: formdata,
      redirect: 'follow',
      credentials: 'include',
    } as RequestInit;

    const apiResponse = await fetch(import.meta.env.VITE_BACKEND_URL + "/api/v0/me/profile", requestOptions) as Response;
    return apiResponse;
  }
}
