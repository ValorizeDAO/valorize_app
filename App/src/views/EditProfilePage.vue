
<template>
  <div id="edit-profile-page" class="flex h-screen">
    <div
      id="left-pane"
      class="md:w-9/12 pl-16 pt-4 border-r-2 border-black h-full"
    >
      <div class="grid grid-cols-6 gap-8">
        <div class="col-span-6">
          <h1 class="text-3xl font-black">Your Profile</h1>
          <h2 class="text-xl font-black">javier123454321</h2>
        </div>
        <div class="col-span-2">
          <div class="relative mt-6 -ml-2">
            <div class="h-52 w-52 object-cover absolute transform translate-x-2 -translate-y-2 overflow-hidden">
            <img
              :src="profileImage === ''? '/src/assets/img/default_avatar.png' : profileImage"
              alt="default avatar"
            />
            </div>
            <img class="h-52 w-52" src="../assets/img/dotted_shadow.png" alt="dotted shadow" />
          </div>
          <transform name="Fade">
            <button
              v-if="pictureStatusInit"
              @click="upload"
              class="
                px-4
                py-2
                cursor-pointer
                border-2 border-black
                rounded-md
                font-black
                w-48
                my-4
                bg-purple-100
              "
            >
              Change Picture
            </button>
            <div class="flex justify-around" v-else>
              <button
                class="
                  px-4
                  py-2
                  cursor-pointer
                  border-2 border-black
                  rounded-md
                  font-black
                  w-48
                  my-4
                  bg-purple-100
                "
              >
                Save
              </button>
              <button
                class="
                  px-4
                  py-2
                  cursor-pointer
                  border-2 border-black
                  rounded-md
                  font-black
                  w-48
                  my-4
                  bg-purple-100
                "
              >
                Cancel
              </button>
            </div>
            <input
              @change="changePic"
              type="file"
              name="picture"
              id="picture-upload"
              class="sr-only"
              ref="pictureFormUpload"
            />
          </transform>
          <ul class="flex justify-between">
            <li class="font-black mr-3 text-lg col-span-1">Followers</li>
            <li class="font-black text-lg col-span-1">Investors</li>
          </ul>
        </div>
        <div class="col-span-4 pr-8">
          <form @submit.prevent="(e) => console.log(e)">
            <label>
              <p class="font-black mb-4">Your Name</p>
              <input
                name="name"
                type="text"
                placeholder="E.G. John Doe"
                class="bg-purple-50 border-black border-b-2 w-full"
              />
            </label>
            <label>
              <p class="font-black mt-8">Your Bio</p>
              <textarea
                name="bio"
                type="textarea"
                cols="50"
                rows="10"
                placeholder="Something about yourself"
                class="block bg-purple-50 border-black border-b-2 w-full h-60"
              />
            </label>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
export default defineComponent({
  name: "EditProfilePage",
  props: {},
  setup: () => {
    const pictureFormUpload = ref(null);
    const pictureStatusInit = ref(true);
    const profileImage = ref("");
    function upload() {
      (pictureFormUpload.value as HTMLInputElement).click();
    }
    function changePic(e) {
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length)
        return;
      pictureStatusInit.value = false;
      console.log((pictureFormUpload.value as HTMLInputElement).value)
      var reader = new FileReader();

      reader.onload = (e) => {
        profileImage.value = e.target.result;
      };
      reader.readAsDataURL((pictureFormUpload.value as HTMLInputElement).files[0]);
    }
    return { pictureFormUpload, pictureStatusInit, profileImage, upload, changePic };
  },
});
</script>

<style scoped>
</style>
