
<template>
  <div id="edit-profile-page" class="flex h-screen">
    <div
      id="left-pane"
      class="md:w-9/12 pl-16 pt-4 border-r-2 border-black h-full"
    >
      <div class="grid grid-cols-6 gap-4">
        <div class="col-span-2">
          <h1 class="text-3xl font-black">Your Profile</h1>
          <h2 class="text-xl font-black">javier123454321</h2>
          <div class="relative mt-6 -ml-2">
            <img
              class="absolute transform translate-x-2 -translate-y-2"
              src="../assets/img/default_avatar.png"
              alt="default avatar"
            />
            <img src="../assets/img/dotted_shadow.png" alt="dotted shadow" />
          </div>
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
            "
          >
            Upload Picture
          </button>
          <ul class="flex">
            <li class="font-black mr-3 text-lg">Followers</li>
            <li class="font-black text-lg">Investors</li>
          </ul>
        </div>
        <div>col2</div>
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
    const count = ref(0);
    return { count };
  },
});
</script>

<style scoped>
</style>
