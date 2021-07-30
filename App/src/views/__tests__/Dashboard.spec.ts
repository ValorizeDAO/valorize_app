import { shallowMount, DOMWrapper } from "@vue/test-utils"
import Register from "@/views/Register.vue"
import Vuex from 'vuex'

describe("Dashboard.vue", () => {
  it("mounts", () => {
    const wrapper = shallowMount(Register, {
      global: {
        plugins: [Vuex],
      },
    } as any);
    expect(wrapper).toBeTruthy()
  })
 it("renders as expected", () => {
  const wrapper = shallowMount(Register, {
    global: {
      plugins: [Vuex],
    },
  } as any);
    expect(wrapper.html()).toMatchSnapshot()
  })
})
