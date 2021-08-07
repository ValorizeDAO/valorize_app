
<template>
  <div id="edit-profile-page" class="grid grid-cols-12 gap-8 min-h-screen">
    <div id="left-pane" class="md:col-span-7 pl-16 pt-4 h-full">
      <div class="">
        <div class="">
          <h1 class="text-3xl font-black mb-6">Your Profile</h1>
          <h2 class="text-2xl font-black">{{ user.username }}</h2>
        </div>
        <div class="col-span-2">
          <div class="relative mt-6 -ml-2">
            <ImageContainer>
              <div v-if="pictureStatus === 'INIT'">
                <img
                  :src="user.avatar"
                  alt="avatar"
                  class="h-52 w-52 object-cover"
                />
              </div>
              <div v-else>
                <img
                  :src="
                    profileImage === ''
                      ? '/src/assets/img/default_avatar.png'
                      : profileImage
                  "
                  class="h-52 w-52 object-cover"
                />
              </div>
            </ImageContainer>
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
      <h2 class="text-3xl font-black mb-6">Your Token</h2>
      <TokenInfoComponent v-if="user.has_deployed_token" :username="user.username" />
      <div v-else>
      <h3 class="text-2xl font-black">{{ user.username }}'s Token</h3>
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
        <p class="font-black my-4">Token Symbol</p>
        <input
          type="text"
          name="tokenSymbol"
          id="token-tiker"
          v-model="tokenSymbol"
          placeholder="e.g TKN"
          class="bg-paper-light border-black border-b-2 w-full"
        />
      </label>
      <div class="text-center">
        <p class="my-8">(Get 1000 of this token by deploying the contract)</p>
        <button @click="openModal" class="btn w-48 my-4 bg-paper-darker">
          Test Deploy {{ tokenSymbol }}
        </button>
      </div>
      </div>
      <div
        @click.stop="openModal"
        id="modal-bg"
        v-if="modalIsOpen"
        class="
          mt-20
          h-screen
          w-screen
          absolute
          top-0
          left-0
          bg-black bg-opacity-70
          flex
          justify-center
          items-center
        "
      >
        <div
          id="modal-body"
          @click.stop
          class="bg-paper-light w-9/12 h-80 mx-auto px-10"
        >
          <div class="flex w-full justify-end">
            <button
              @click="openModal"
              class="my-2 -ml-8 p-1 border-black border rounded"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                  clip-rule="evenodd"
                />
              </svg>
            </button>
          </div>
          <div class="my-4 mx-auto">
            <transition name="fade" mode="out-in">
              <div
                class="text-center"
                v-if="
                  tokenDeployStatus === 'INIT' || tokenDeployStatus === 'ERROR'
                "
              >
                <h1 class="text-2xl">
                  Deploy <span class="font-black mb-12">{{ tokenName }}</span>
                </h1>
                <p class="my-4 max-w-sm mx-auto">Coin will be on the Ropsten Ethereum test network, you will have a chance to confirm details
                  there</p>
                <button
                  @click="deployToTestNet"
                  class="btn w-48 mt-4 bg-purple-50"
                >
                  Test Deploy {{ tokenSymbol }}
                </button>
              </div>
              <SvgLoader
                class="text-center mx-auto"
                v-else-if="tokenDeployStatus === 'DEPLOYING'"
                fill="#"
              ></SvgLoader>
              <div
                v-else-if="tokenDeployStatus === 'SUCCESS'"
                class="text-center my-6"
              >
                <h1 class="text-2xl font-black">
                  Woo! You can now review the test version of {{ tokenName }}!
                </h1>
                <p class="my-6">
                  <a
                      class="font-black underline text-center text-lg"
                      :href="
                    'https://ropsten.etherscan.io/tx/' +
                    tokenTestnetTx
                  "
                      target="_blank"
                  >
                    Confirm details
                  </a>
                </p>
                <p>(Might take a minute to Deploy)</p>
                <a :href="checkoutLink">
                  <div class="btn w-1/2 mx-auto bg-purple-100 mt-12">Deploy on Ethereum for $10</div>
                </a>
              </div>
            </transition>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent, computed } from "vue";
import auth from "../services/authentication";
import { User } from "../models/user";
import ethApi, { TokenResponse } from "../services/ethApi";
import { useStore } from "vuex";
import SvgLoader from "../components/SvgLoader.vue";
import ImageContainer from "../components/ImageContainer.vue";
import TokenInfoComponent from "../components/TokenInfoComponent.vue";
export default defineComponent({
  name: "EditProfilePage",
  props: {},
  components: { SvgLoader, ImageContainer, TokenInfoComponent },
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
  const hasToken = store.getters["authUser/hasToken"]

  const profileUpdateStatuses = [
    "INIT",
    "EDITED",
    "UPLOADING",
    "ERROR",
    "SUCCESS",
  ];
  const profileUpdateStatus = ref(profileUpdateStatuses[0]);
  async function updateProfile() {
    profileUpdateStatus.value = profileUpdateStatuses[2];
    const response = await auth.updateProfile({
      name: fullName.value,
      about: about.value,
    });
    if (response.status !== 200) {
      profileUpdateStatus.value = profileUpdateStatuses[3];
      return;
    }
    profileUpdateStatus.value = profileUpdateStatuses[4];
    const userData = (await response.json()) as Promise<User>;
    store.commit("authUser/setUser", userData);
  }
  return {
    updateProfile,
    fullName,
    about,
    profileUpdateStatus,
    hasToken
  };
}
function composeUpdateImage() {
  const store = useStore();
  const pictureFormUpload = ref(HTMLInputElement);
  const profileImage = ref(store.state.authUser.user.avatar);
  const pictureStatuses = ["INIT", "PREVIEW", "UPLOADING", "ERROR"];
  const pictureStatus = ref<string>(pictureStatuses[0]);
  const imageToUpload = ref<File>(new File([], ""))
  function changeProfile() {
    (pictureFormUpload.value as unknown as HTMLInputElement).click();
  }
  function changePic(e: Event) {
    if (e.target) {
      const inputElement = e.target as HTMLInputElement
      const files = inputElement.files
      pictureStatus.value = pictureStatuses[1];
      if (!files) {
        return
      }
      const fileReader = new FileReader()
      fileReader.addEventListener('load', (e) => {
        profileImage.value = URL.createObjectURL(files[0])
        imageToUpload.value = files[0]
      })
      fileReader.readAsArrayBuffer(files[0])
    }
  }

  async function sendPhoto() {
    pictureStatus.value = pictureStatuses[2];
    const uploadRequest = await auth.uploadPicture(imageToUpload.value)
    if (uploadRequest.status == 200) {
      pictureStatus.value = pictureStatuses[0];
      const responseJson = await ((await uploadRequest.json()) as Promise<{
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
  const tokenName = ref(store.state.authUser.user.username + "Coin");
  const tokenSymbol = ref("TKN");
  const modalIsOpen = ref(false);
  const tokenDeployStatuses = ["INIT", "DEPLOYING", "SUCCESS", "ERROR"];
  const tokenDeployStatus = ref(tokenDeployStatuses[0]);
  const tokenTestnetTx = ref("");
  const checkoutLink = computed(() => {
    const encodedName = encodeURIComponent(tokenName.value)
    const encodedSymbol = encodeURIComponent(tokenSymbol.value)
    return `${import.meta.env.VITE_BACKEND_URL}/create-checkout-session?tokenName=${encodedName}&tokenSymbol=${encodedSymbol}`
  });
  async function deployToTestNet() {
    tokenDeployStatus.value = tokenDeployStatuses[1];
    const apiResponse = await ethApi.deployTokenToTestNet({
      tokenName: tokenName.value,
      tokenSymbol: tokenSymbol.value,
    });
    if (apiResponse.status == 200) {
      tokenDeployStatus.value = tokenDeployStatuses[2];
      const responseJson = await (apiResponse.json() as Promise<TokenResponse>);
      tokenTestnetTx.value = responseJson.tx;
    } else {
      tokenDeployStatus.value = tokenDeployStatuses[3];
    }
  }
  function openModal(e: Event) {
    modalIsOpen.value = !modalIsOpen.value;
  }
  return {
    tokenName,
    tokenSymbol,
    modalIsOpen,
    deployToTestNet,
    openModal,
    tokenDeployStatus,
    tokenTestnetTx,
    checkoutLink,
  };
}
</script>

<style scoped lang="postcss">
</style>