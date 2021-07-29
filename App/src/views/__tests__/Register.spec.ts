import { shallowMount, DOMWrapper } from "@vue/test-utils"
import Register from "@/views/Register.vue"
import Vuex from 'vuex'

describe("Register.vue", () => {
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
  describe("Submit Button", () => {
    let goodData: any; 
    beforeEach(() => {
      goodData = {
        password: "test",
        passwordVerify: "test",
        email: "test@test.com",
        username: "fakename"
      }
    })
    it("is disabled when all fields are empty", () => {
      const wrapper = shallowMount(Register, {
        global: {
          plugins: [Vuex],
        },
      } as any);
      const submitButton = wrapper.find("input[type='submit']") as DOMWrapper<HTMLInputElement>;
      expect(submitButton.element.disabled).toBe(true)
    })
    // it("should need both passwords to equal to be enabled", async () => {
    //   const wrapper = shallowMount(Register, {
    //     data() { return {...goodData} },
    //     global: {
    //       plugins: [Vuex],
    //     },
    //   } as any)
    //   const submitButton = wrapper.find("input[type='submit']") as DOMWrapper<HTMLInputElement>
    //   const userNameInput = wrapper.find("input[name='password']") as DOMWrapper<HTMLInputElement>;
    //   expect(submitButton.element.disabled).toBe(false)
    // })
    it("should be disabled if one password is unlike the other", async () => {
      goodData.password = "a"
      const wrapper = shallowMount(Register, {
        data() { return { ...goodData } },
        global: {
          plugins: [Vuex],
        },
      } as any)
      const submitButton = wrapper.find("input[type='submit']") as DOMWrapper<HTMLInputElement>; 
      
      expect(submitButton.element.disabled).toBe(true)
    })
  })
})
