
<template>
  <div id="edit-profile-page" class="flex h-screen">
    <div
      id="left-pane"
      class="md:w-9/12 pl-16 pt-4 border-r-2 border-black h-full"
    >
      <div class="grid grid-cols-6 gap-8">
        <div class="col-span-6">
          <h1 class="text-3xl font-black">Your Profile</h1>
          <h2 class="text-xl font-black">{{ user.username }}</h2>
        </div>
        <div class="col-span-2">
          <div class="relative mt-6 -ml-2">
            <div
              class="
                h-52
                w-52
                object-cover
                absolute
                transform
                translate-x-2
                -translate-y-2
                overflow-hidden
              "
            >
              <div v-if="pictureStatus === 'INIT'">
                <img :src="user.avatar" alt="avatar" />
              </div>
              <div v-else>
                <img
                  :src="
                    profileImage === ''
                      ? '/src/assets/img/default_avatar.png'
                      : profileImage
                  "
                />
              </div>
            </div>
            <img
              class="h-52 w-52"
              src="../assets/img/dotted_shadow.png"
              alt="dotted shadow"
            />
          </div>
          <transform name="fade">
            <div v-if="pictureStatus === 'INIT'">
              <button
                @click="changeProfile"
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
            </div>
            <div
              class="flex justify-around"
              v-else-if="pictureStatus == 'PREVIEW' || pictureStatus == 'ERROR'"
            >
              <button
                @click="sendPhoto"
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
                  mr-4
                "
              >
                Save
              </button>
              <button
                @click="resetPhoto"
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
              <div v-if="pictureStatus == 'ERROR'">
                <span class="text-red-700 text-sm">
                  There was an error changing your photo. Try again or contact us.
                </span>
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
import auth from "../services/authentication";
import { useStore } from "vuex";
export default defineComponent({
  name: "EditProfilePage",
  props: {},
  setup: () => {
    const store = useStore();
    const pictureFormUpload = ref(null);
    const profileImage = ref(store.state.user.avatar);
    const pictureStatuses = ["INIT", "PREVIEW", "UPLOADING", "ERROR"];
    const pictureStatus = ref(pictureStatuses[0]);
    function changeProfile() {
      (pictureFormUpload.value as HTMLInputElement).click();
    }
    function changePic(e) {
      var files = e.target.files || e.dataTransfer.files;
      pictureStatus.value = pictureStatuses[1];

      if (!files.length) {
        return;
      }
      var reader = new FileReader();

      reader.onload = (e) => {
        profileImage.value = e.target.result;
      };
      reader.readAsDataURL(
        (pictureFormUpload.value as HTMLInputElement).files[0]
      );
    }
    async function sendPhoto() {
      pictureStatus.value = pictureStatuses[2];
      const uploadRequest = await auth.uploadPicture(
        (pictureFormUpload.value as HTMLInputElement).files[0]
      );
      if (uploadRequest.status == 200) {
        pictureStatus.value = pictureStatuses[0];
        const responseJson = await (uploadRequest.json() as Promise<{ image: string }>)
        store.commit("setUserPicture", responseJson.image+ '?' + new Date().getTime())
      } else {
        pictureStatus.value = pictureStatuses[3];
      }
    }
    function resetPhoto() {
      pictureStatus.value = pictureStatuses[0];
      profileImage.value = store.state.user.avatar
      (pictureFormUpload.value as HTMLInputElement).outerHTML = (pictureFormUpload.value as HTMLInputElement).outerHTML;
    }

    return {
      pictureFormUpload,
      pictureStatus,
      profileImage,
      changeProfile,
      changePic,
      sendPhoto,
      resetPhoto,
      user: store.state.user,
    };
  },
});
</script>

<style scoped>
</style>
