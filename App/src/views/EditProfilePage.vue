
<template>
  <div id="edit-profile-page" class="grid grid-cols-12 gap-8 min-h-screen">
    <div id="left-pane" class="md:col-span-7 pl-16 pt-4 h-full">
      <div class="">
        <div class="">
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
                class="btn w-48 my-4 bg-purple-100"
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
                class="btn w-48 my-4 bg-purple-100 mr-4"
              >
                Save
              </button>
              <button @click="resetPhoto" class="btn w-48 my-4 bg-purple-100">
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
          <!-- <ul class="flex justify-between">
            <li class="font-black mr-3 text-lg col-span-1">Followers</li>
            <li class="font-black text-lg col-span-1">Investors</li>
          </ul> -->
        </div>
        <div class="col-span-4 pr-8">
          <form @submit.prevent="updateProfile">
            <label>
              <p class="font-black mb-4">Your Name</p>
              <input
                name="name"
                type="text"
                v-model="fullName"
                placeholder="e.g. John Doe"
                class="bg-purple-50 border-black border-b-2 w-full"
              />
            </label>
            <label>
              <p class="font-black mt-8">Your Bio</p>
              <textarea
                name="about"
                type="textarea"
                cols="50"
                rows="10"
                v-model="about"
                placeholder="Something about yourself"
                class="block bg-purple-50 border-black border-b-2 w-full h-60"
              />
            </label>
            <button class="btn w-48 my-4 bg-purple-100">Update Info</button>
          </form>
        </div>
      </div>
    </div>
    <div
      id="right-pane"
      class="
        md:col-span-5
        p-4
        pr-16
        border-l-2 border-black
        h-full
        bg-paper-light
      "
    >
      <h3 class="text-3xl font-black">{{ user.username }}'s Token</h3>
      ( not yet deployed )
      <label>
        <p class="font-black mb-4">Token Name</p>
        <input
          type="text"
          name="tokenName"
          id="token-name"
          v-model="tokenName"
          class="bg-paper-light border-black border-b-2 w-full"
        />
      </label>
      <label>
        <p class="font-black my-4">Token Ticker</p>
        <input
          type="text"
          name="tokenTiker"
          id="token-tiker"
          v-model="tokenTiker"
          class="bg-paper-light border-black border-b-2 w-full"
        />
      </label>
      <div class="text-center">
        <p class="my-8">(Get 1000 of this token by deploying the contract)</p>
        <button @click="openModal" class="btn w-48 my-4 bg-paper-darker">
          Test Deploy
        </button>
      </div>
      <div
        v-if="modalIsOpen"
        class="h-screen w-screen absolute top-0 left-0 bg-paper-lighter"
      >
        <div class="flex w-full justify-between px-16 py-10">
          <button @click="openModal">X</button>
        </div>
        <div class="my-4 mx-auto">
          <h1 class="text-lg">Deploy <span class="font-black">{{ tokenName }}</span></h1>
          <transition name="fade">
            <button v-if="tokenDeployStatus === 'INIT'" @click="deployToTestNet" class="btn w-48 my-4 bg-paper-darker">
              Deploy Test Net
            </button>
            <SvgLoader v-else-if="tokenDeployStatus === 'DEPLOYING'" fill="purple"></SvgLoader>
            <a v-else-if="tokenDeployStatus === 'SUCCESS'" :href="'https://ropsten.etherscan.io/address/' + tokenTestnetAddress" target="_blank">See on Testnet</a>
          </transition>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import auth from "../services/authentication";
import { User } from "../models/user";
import ethApi, { TokenResponse } from "../services/ethApi";
import { useStore } from "vuex";
export default defineComponent({
  name: "EditProfilePage",
  props: {},
  setup() {
    return {
      ...composeProfileInfo(),
      ...composeUpdateImage(),
      ...composeDeployToken(),
    };
  },
});
function composeProfileInfo() {
  const store = useStore();
  const fullName = ref(store.state.authUser.user.name);
  const about = ref(store.state.authUser.user.about);
  
  const profileUpdateStatuses = [
    "INIT",
    "EDITED",
    "UPLOADING",
    "ERROR",
    "SUCCESS",
  ];
  const profileUpdateStatus = ref(profileUpdateStatuses[0]);
  async function updateProfile() {
    profileUpdateStatus.value = profileUpdateStatus[2];
    const response = await auth.updateProfile({
      name: fullName.value,
      about: about.value,
    });
    if (response.status !== 200) {
      profileUpdateStatus.value = profileUpdateStatus[3];
      return;
    }
    profileUpdateStatus.value = profileUpdateStatus[4];
    const userData = (await response.json()) as Promise<User>;
    store.commit("authUser/setUser", userData);
  }
  return {
    updateProfile,
    fullName,
    about,
    profileUpdateStatus,
  };
}
function composeUpdateImage() {
  const store = useStore();
  const pictureFormUpload = ref(null);
  const profileImage = ref(store.state.authUser.user.avatar);
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
      const responseJson = await (uploadRequest.json() as Promise<{
        image: string;
      }>);
      store.commit(
        "authUser/setUserPicture",
        responseJson.image + "?" + new Date().getTime()
      );
    } else {
      pictureStatus.value = pictureStatuses[3];
    }
  }
  function resetPhoto() {
    pictureStatus.value = pictureStatuses[0];
    profileImage.value = store.state.authUser.user.avatar;
    (pictureFormUpload.value as HTMLInputElement).outerHTML = (
      pictureFormUpload.value as HTMLInputElement
    ).outerHTML;
  }
  return {
    pictureFormUpload,
    pictureStatus,
    profileImage,
    changeProfile,
    changePic,
    sendPhoto,
    resetPhoto,
    user: store.state.authUser.user,
  };
}
function composeDeployToken() {
  const store = useStore();
  const tokenName = ref(store.state.authUser.user.username + ' token');
  const tokenTiker = ref("TKN");
  const modalIsOpen = ref(false);
  const tokenDeployStatuses = ["INIT", "DEPLOYING", "SUCCESS", "ERROR"];
  const tokenDeployStatus = ref(tokenDeployStatuses[0]);
  const tokenTestnetAddress = ref("")
  async function deployToTestNet() {
    tokenDeployStatus.value = tokenDeployStatuses[1];
    const apiResponse = await ethApi.deployTokenFromBackend({
      tokenName: tokenName.value,
      tokenTiker: tokenTiker.value,
    });
    if (apiResponse.status == 200) {
      tokenDeployStatus.value = tokenDeployStatuses[2];
      const responseJson = await (apiResponse.json() as Promise<TokenResponse>);
      store.commit("setToken", responseJson.token);
      tokenTestnetAddress.value = responseJson.token.address;
    } else {
      tokenDeployStatus.value = tokenDeployStatuses[3];
    }
  }
  function openModal() {
    modalIsOpen.value = !modalIsOpen.value;
  }
  return {
    tokenName,
    tokenTiker,
    modalIsOpen,
    deployToTestNet,
    openModal,
    tokenDeployStatus,
    tokenTestnetAddress,
  };
}
</script>

<style scoped lang="postcss">
.btn {
  @apply px-4 py-2 cursor-pointer border-2 border-black rounded-md font-black;
}
</style>