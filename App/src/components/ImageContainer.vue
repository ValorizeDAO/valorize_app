<template>
  <div>
    <div
      class="absolute transform translate-x-2 -translate-y-2 overflow-hidden">
      <slot> </slot>
    </div>
    <img
      alt="dotted shadow"
      class="h-52 w-52"
      src="../assets/img/dotted_shadow.png"
    />
  </div>
</template>
<script lang="ts">
export default{
  name: "ImageContainer",
  props: {
    tailwindSize: {
      type: String,
      default: "h-52 w-52"
    }
  }
}
</script>
