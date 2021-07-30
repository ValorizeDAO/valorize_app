import { shallowMount, DOMWrapper } from "@vue/test-utils"
import EditProfilePage from "@/views/EditProfilePage.vue"
import Vuex from 'vuex'

describe("Dashboard.vue", () => {
  it("mounts", () => {
    const wrapper = shallowMount(EditProfilePage, {
      global: {
        plugins: [Vuex],
      },
    } as any);
    expect(wrapper).toBeTruthy()
  })
 it("renders as expected", () => {
  const wrapper = shallowMount(EditProfilePage, {
    global: {
      plugins: [Vuex],
    },
  } as any);
    expect(wrapper.html()).toMatchSnapshot()
  })
})
