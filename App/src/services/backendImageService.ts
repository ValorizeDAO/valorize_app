export default function setUserPicturePrefix(filename: string): string {
  return (
    import.meta.env.VITE_BACKEND_URL +
        "/static/images/" +
        (filename || "default_avatar.jpg")
  )
}
